// Package checkvarlist provides functionality to check varlist files for duplicates.
// It can detect duplicates within a single file and across multiple files.
package checkvarlist

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

// DuplicateReport holds information about duplicate variables found during checking.
type DuplicateReport struct {
	// DuplicatesInFile contains variables that appear multiple times within a single file
	DuplicatesInFile map[string][]string
	// DuplicateInMultipleFiles contains variables that appear in more than one file
	DuplicateInMultipleFiles map[string][]string
	// TotalFilesChecked is the number of varlist files that were checked
	TotalFilesChecked int
	// TotalDuplicates is the total count of duplicate entries found
	TotalDuplicates int
}

// CheckVarlists scans the output directory for varlist and stringvarlist files
// and checks for duplicate variables within files and across files.
//
// The function looks for files matching:
//   - output/varlist/*_varlist.txt
//   - output/stringvarlisten/*_stringvarlist.txt
//
// It returns a DuplicateReport with any duplicates found, or an error if
// the directory cannot be read.
func CheckVarlists(dir string) (*DuplicateReport, error) {
	// Find all varlist files
	varFiles, err := filepath.Glob(filepath.Join(dir, "varlist", "*_varlist.txt"))
	if err != nil {
		return nil, fmt.Errorf("error finding varlist files: %w", err)
	}

	// Find all stringvarlist files
	stringFiles, err := filepath.Glob(filepath.Join(dir, "stringvarlisten", "*_stringvarlist.txt"))
	if err != nil {
		return nil, fmt.Errorf("error finding stringvarlist files: %w", err)
	}

	// Combine and sort file list
	files := append(varFiles, stringFiles...)
	sort.Strings(files)

	// Initialize report
	report := &DuplicateReport{
		DuplicatesInFile:         make(map[string][]string),
		DuplicateInMultipleFiles: make(map[string][]string),
		TotalFilesChecked:        len(files),
	}

	// Track which variables appear in which files
	fileVars := make(map[string][]string)
	varInFiles := make(map[string][]string)

	// Process each file
	for _, file := range files {
		vars, err := readVarlist(file)
		if err != nil {
			return nil, fmt.Errorf("error reading %s: %w", file, err)
		}

		filename := filepath.Base(file)
		fileVars[filename] = vars

		// Check for duplicates within this file
		seen := make(map[string]bool)
		var unique []string
		for _, v := range vars {
			if !seen[v] {
				seen[v] = true
				unique = append(unique, v)
			} else {
				report.DuplicatesInFile[filename] = append(report.DuplicatesInFile[filename], v)
			}
		}

		// Track which files contain each variable
		for _, v := range unique {
			varInFiles[v] = append(varInFiles[v], filename)
		}
	}

	// Find variables that appear in multiple files
	for varName, files := range varInFiles {
		if len(files) > 1 {
			report.DuplicateInMultipleFiles[varName] = files
		}
	}

	// Count total duplicates
	for _, dups := range report.DuplicatesInFile {
		report.TotalDuplicates += len(dups)
	}
	report.TotalDuplicates += len(report.DuplicateInMultipleFiles)

	return report, nil
}

// readVarlist reads a varlist file and returns its contents as a slice of strings.
// Each line in the file represents one variable name.
func readVarlist(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var vars []string
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			vars = append(vars, line)
		}
	}
	return vars, nil
}

// Print outputs the duplicate report in a human-readable format to stdout.
func (r *DuplicateReport) Print() {
	fmt.Printf("=== Varlist Check Report ===\n")
	fmt.Printf("Files checked: %d\n\n", r.TotalFilesChecked)

	// Print duplicates within files
	if len(r.DuplicatesInFile) > 0 {
		fmt.Println("--- Duplicates WITHIN files ---")
		for filename, dups := range r.DuplicatesInFile {
			fmt.Printf("\n%s:\n", filename)
			for _, d := range dups {
				fmt.Printf("  - %s\n", d)
			}
		}
		fmt.Println()
	}

	// Print duplicates across files
	if len(r.DuplicateInMultipleFiles) > 0 {
		fmt.Println("--- Duplicates ACROSS files ---")
		for varName, files := range r.DuplicateInMultipleFiles {
			fmt.Printf("\n  %s found in:\n", varName)
			for _, f := range files {
				fmt.Printf("    - %s\n", f)
			}
		}
		fmt.Println()
	}

	// Print summary
	if r.TotalDuplicates == 0 {
		fmt.Println("✓ No duplicates found!")
	} else {
		fmt.Printf("Total duplicates: %d\n", r.TotalDuplicates)
	}
}
