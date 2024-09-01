# Json2Prom

Json2Prom is a Go application that converts Prometheus JSON dumps to OpenMetrics format. The output of this tool can be bulk loaded into a running Prometheus instance for visualization.

## Features

- Traverse directories recursively to find JSON files.
- Process and convert Prometheus JSON dumps to OpenMetrics format.
- Handle errors gracefully.
- Exclude files based on specific criteria.

## Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/samirdasiitr/json2prom.git
    ```
2. Navigate to the project directory:
    ```sh
    cd json2prom
    ```
3. Build the application:
    ```sh
    go build -o json2prom main.go
    ```

## Usage

1. Place the JSON dumps at a specified location.
2. Run the application with the following command:
    ```sh
    ./json2prom --input /path/to/json/files > tsdb.txt
    ```
3. Load the output into Prometheus:
    ```sh
    promtool tsdb create-blocks-from openmetrics tsdb.txt /prometheus
    ```

## Code Explanation

The main functionality is implemented in the `walkDir` function, which is called for each file and directory encountered during the traversal.

### `walkDir` Function

- **Parameters:**
  - `path`: The path of the current file or directory.
  - `info`: FileInfo object containing information about the file.
  - `err`: Error encountered while accessing the file.

- **Functionality:**
  - Checks for errors accessing the file and logs them.
  - Processes only regular files.
  - Excludes files containing a colon (`:`) in their path.
  - Calls `processFile` to process each valid file.

### `processFile` Function

- **Parameters:**
  - `fileName`: The name of the file to process.

- **Functionality:**
  - Reads the file contents.
  - Parses the JSON data into a `Sample` struct.
  - Converts the data to OpenMetrics format.
  - Writes the converted data to a new file in the `dump` directory.

## Dependencies

- Go standard library packages:
  - `encoding/json`
  - `flag`
  - `fmt`
  - `log`
  - `os`
  - `path/filepath`
  - `strings`

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -am 'Add new feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Create a new Pull Request.

## License

This project is licensed under the MIT License. See the LICENSE file for details.

## Contact

For any questions or suggestions, please open an issue or contact the repository owner.
