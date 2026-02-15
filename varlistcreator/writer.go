// Package varlistcreator provides functionality to write variables to a file.
package varlistcreator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// WriteVariables takes a source filename and a slice of variable names,
// then writes them to a varlist file in a "varlist" directory.
// The output filename is based on the source filename with "_varlist.txt" suffix.
func WriteVariables(filename string, variables []string) error {
	baseName := strings.TrimSuffix(filepath.Base(filename), filepath.Ext(filename))
	outputFilename := baseName + "_varlist.txt"

	if err := os.MkdirAll("varlist", 0755); err != nil {
		return fmt.Errorf("error creating directory: %w", err)
	}

	file, err := os.Create(filepath.Join("varlist", outputFilename))
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
