# SimpleVarlistGenerator

A command-line tool for extracting unique variables from `.osc` files (OMSI Script files) and generating variable list files.

## Overview

The SimpleVarlistGenerator reads OMSI Script files (.osc), extracts unique variable names, and outputs them into separate text files. It distinguishes between:

- **Written variables** (S.L) - Variables that are written/set by the script
- **Read variables** (L.L) - Variables that are only read by the script
- **Written string variables** (S.$) - String variables that are written/set
- **Read string variables** (L.$) - String variables that are only read

## Features

- Extracts unique variable names from `.osc` files
- Categorizes variables by type (regular/string) and access mode (read/write)
- Prevents duplicate variables across multiple files
- Checks for duplicate variables in existing varlists
- Clean, formatted console output

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/f0xb17/SimpleVarlistGenerator
   ```
2. Navigate to the project directory:
   ```sh
   cd SimpleVarlistGenerator
   ```
3. Build the application:
   ```sh
   go build -o simplevarlistgenerator .
   ```

## Usage

### Command Options

| Option | Description |
|--------|-------------|
| (none) | Process all `.osc` files in current directory |
| `-f, --file <file.osc>` | Process a single file |
| `-c, --check [dir]` | Check varlists for duplicates (default: output) |
| `-h, --help` | Show help |

### Examples

Process all `.osc` files in current directory:
```sh
./simplevarlistgenerator
```

Process a single file:
```sh
./simplevarlistgenerator -f antrieb.osc
```

Check for duplicates in output directory:
```sh
./simplevarlistgenerator -c
```

Check a specific directory:
```sh
./simplevarlistgenerator -c /path/to/output
```

## Output

The tool creates the following directory structure:

```
output/
├── varlist/
│   ├── antrieb_varlist.txt
│   ├── engine_varlist.txt
│   └── ...
└── stringvarlisten/
    ├── cockpit_stringvarlist.txt
    └── ...
```

- **output/varlist/**: Contains regular variables (both S.L and L.L)
- **output/stringvarlisten/**: Contains string variables (both S.$ and L.$)

## Console Output

When processing files, the tool shows:

- **[W]**: Written variables (S.L) - Variables that are written by the script
- **[S]**: Written string variables (S.$) - String variables that are written
- **[R]**: Read-only variables (L.L and L.$) - Variables that are only read (informational)

### Example Output

```
============================================
  Processing 26 OSC files
============================================

  [W] antrieb. variables (S.Losc: 28)
  [S] cockpit.osc: 1 string variables (S.$)
  [R] cockpit.osc: 167 variables (L.L) | 0 string variables (L.$)

============================================
  Total Written Variables (S.L):   245
  Total Written String Vars (S.$):  12
  Total Read Variables (L.L):       189
  Total Read String Vars (L.$):     5
============================================
```

## Check Function

The `-c` option scans all varlist files and reports:

- **Duplicates WITHIN files**: Variables that appear multiple times in a single file
- **Duplicates ACROSS files**: Variables that appear in more than one file

Example:
```
=== Varlist Check Report ===
Files checked: 28

--- Duplicates ACROSS files ---

  engine_on found in:
    - G81_main_varlist.txt
    - antrieb_varlist.txt

Total duplicates: 3
```

## Variable Prefixes

OMSI Script uses the following prefixes:

| Prefix | Meaning |
|--------|---------|
| `(S.L.variable)` | Local variable - written by the script |
| `(L.L.variable)` | Global variable - read from memory |
| `(S.$.variable)` | Local string variable - written by the script |
| `(L.$.variable)` | Global string variable - read from memory |

## Directory Structure

```
SimpleVarlistGenerator/
├── go.mod
├── main.go
├── oscextract/
│   └── extract.go
├── varlistcreator/
│   └── writer.go
├── checkvarlist/
│   └── check.go
└── output/
    ├── varlist/
    └── stringvarlisten/
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request or raise an issue.
