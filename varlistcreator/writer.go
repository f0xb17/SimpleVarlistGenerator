// Package varlistcreator provides functionality to write extracted variables to varlist files.
// It creates two types of output files:
//   - Regular varlist files (*_varlist.txt) containing S.L and L.L variables
//   - String varlist files (*_stringvarlist.txt) containing S.$ and L.$ variables
package varlistcreator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// WriteVariables writes the given variables to a varlist file in the output/varlist directory.
// The output filename is based on the source OSC filename with "_varlist.txt" suffix.
//
// Variables are written one per line. If the variable slice is empty, no file is created.
func WriteVariables(filename string, variables []string) error {
	baseName := strings.TrimSuffix(filepath.Base(filename), filepath.Ext(filename))
	outputFilename := baseName + "_varlist.txt"

	// Create output directory if it doesn't exist
	if err := os.MkdirAll("output/varlist", 0755); err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	// Only write if there are variables to write
	if len(variables) > 0 {
		if err := writeFile(filepath.Join("output/varlist", outputFilename), variables); err != nil {
			return err
		}
	}

	return nil
}

// WriteStringVariables writes the given string variables to a stringvarlist file
// in the output/stringvarlisten directory.
// The output filename is based on the source OSC filename with "_stringvarlist.txt" suffix.
//
// String variables are written one per line. If the variable slice is empty, no file is created.
func WriteStringVariables(filename string, variables []string) error {
	baseName := strings.TrimSuffix(filepath.Base(filename), filepath.Ext(filename))
	outputFilename := baseName + "_stringvarlist.txt"

	// Create output directory if it doesn't exist
	if err := os.MkdirAll("output/stringvarlisten", 0755); err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	// Only write if there are variables to write
	if len(variables) > 0 {
		if err := writeFile(filepath.Join("output/stringvarlisten", outputFilename), variables); err != nil {
			return err
		}
	}

	return nil
}

// writeFile creates a new file at the given path and writes each variable
// from the slice on its own line.
func writeFile(path string, variables []string) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	for _, v := range variables {
		_, err := fmt.Fprintln(file, v)
		if err != nil {
			return fmt.Errorf("error writing to file: %w", err)
		}
	}

	return nil
}
