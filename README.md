# SimpleVarlistGenerator

A command-line tool for extracting unique variables from `.osc` files and generating a variable list file.

## Overview

The SimpleVarlistGenerator is a Go application designed to read OMSI Script files (.osc Files), extract unique variable names, and output them into a separate text file.

## Features

- Extracts unique variable names from `.osc` files.
- Writes the extracted variables to a text file named `<original_filename>_varlist.txt`.
- Handles file and directory operations with error checking.

## Getting Started

### Prerequisites

- Go (version 1.16 or higher)

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/f0xb17/SimpleVarlistGenerator
   ```
2. Navigate to the project directory:
   ```sh
   cd SimpleVarlistGenerator
   ```

### Usage

To run the application, use the following command:

```sh
go run main.go <file.osc>
```

### Example

```sh
go run main.go example.osc
```

This command will extract unique variables from `example.osc` and create a file named `example_varlist.txt` in the `varlist directory`.

### Directory Structure

```sh
SimpleVarlistGenerator/
├── go.mod
├── main.go
├── oscextract/
│   └── extract.go
└── varlistcreator/
    └── writer.go
```

### Contributing

Contributions are welcome! Please feel free to submit a pull request or raise an issue.
