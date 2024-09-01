package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func walkDir(path string, info os.FileInfo, err error) error {
	// Check if there was an error accessing the file
	if err != nil {
		fmt.Printf("Error accessing file %s: %v\n", path, err)
		return nil // Ignore error and continue walking
	}

	// Check if it's a regular file
	if !info.Mode().IsRegular() {
		return nil // Continue walking
	}

	if strings.Contains(path, ":") {
		return nil // Continue walking
	}

	processFile(path)

	return nil
}

var dirPath = flag.String("input", "", "directory containing prom json files")

func main() {
	flag.Parse()

	os.MkdirAll("./dump", 0777)

	err := filepath.Walk(*dirPath, walkDir)
	if err != nil {
		fmt.Printf("Error walking directory: %v\n", err)
		return
	}
}

type Metric struct {
	Labels map[string]string `json:"metric"`
	Values [][]interface{}   `json:"values"`
}

type Sample struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string   `json:"resultType"`
		Result     []Metric `json:"result"`
	} `json:"data"`
}

var filesMaps = make(map[string]*os.File)

func processFile(fileName string) {

	fmt.Printf("Processing %q", fileName)

	contents, err := os.ReadFile(fileName)
	if err != nil {
		log.Printf("Failed to read sample file %q, err %v", fileName, err.Error())
		return
	}

	var sample Sample
	err = json.Unmarshal(contents, &sample)
	if err != nil {
		log.Printf("Failed to parse sample file %q, err %v", fileName, err.Error())
		return
	}

	file, err := os.Create("dump/" + filepath.Base(fileName))
	if err != nil {
		log.Fatalf("Failed to create %q file, err: %v", fileName, err.Error())
	}

	metricCount := 0

	for _, metric := range sample.Data.Result {
		var labels []string
		for l_, v_ := range metric.Labels {
			if l_ != "__name__" {
				labels = append(labels, fmt.Sprintf("%s=%q", l_, v_))
			}
		}

		labelStr := strings.Join(labels, ",")
		for _, mt := range metric.Values {
			ts_ := mt[0].(float64)
			data := fmt.Sprintf("%s{%s} %v %d\n", metric.Labels["__name__"], labelStr, mt[1], int64(ts_))
			file.Write([]byte(data))
			metricCount++
		}
	}
	file.Write([]byte("# EOF\n"))

	fmt.Printf("Processed %q, %d found\n", fileName, metricCount)
}
