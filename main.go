// SimpleVarlistGenerator is a tool to extract variable names from OSC files
// and write them to separate list files (varlist and stringvarlist).
// It can process a single file or all .osc files in the current directory.
//
// Usage:
//
//	./simplevarlistgenerator                - Process all .osc files in current directory
//	./simplevarlistgenerator -f <file.osc>  - Process a single file
//	./simplevarlistgenerator -c              - Check for duplicates in output directory
//	./simplevarlistgenerator -h               - Show help
//
// Output:
//
//	output/varlist/           - Contains variables (S.L and L.L)
//	output/stringvarlisten/  - Contains string variables (S.$ and L.$)
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"SimpleVarlistGenerator/checkvarlist"
	"SimpleVarlistGenerator/oscextract"
	"SimpleVarlistGenerator/varlistcreator"
)

// main is the entry point of the application.
// It handles command-line arguments and dispatches to the appropriate processing function.
func main() {
	// No arguments: process all .osc files in current directory
	if len(os.Args) < 2 {
		if err := processAllOscFiles(); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		return
	}

	arg := os.Args[1]

	// Check for duplicates in varlist files
	if arg == "--check" || arg == "-c" {
		dir := "output"
		if len(os.Args) > 2 {
			dir = os.Args[2]
		}
		report, err := checkvarlist.CheckVarlists(dir)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}
		report.Print()
		return
	}

	// Process a single file
	if arg == "-f" || arg == "--file" {
		if len(os.Args) < 3 {
			fmt.Println("Usage: go run main.go -f <file.osc>")
			os.Exit(1)
		}
		filename := os.Args[2]
		processSingleFile(filename)
		return
	}

	// Show help
	if arg == "-h" || arg == "--help" {
		fmt.Println("Usage: go run main.go [options]")
		fmt.Println("")
		fmt.Println("Options:")
		fmt.Println("  -f, --file <file.osc>  Create varlist from single file")
		fmt.Println("  -c, --check [dir]      Check varlists for duplicates (default: output)")
		fmt.Println("  -h, --help             Show this help")
		fmt.Println("")
		fmt.Println("Without options: Create varlists from all .osc files in current directory")
		os.Exit(0)
	}

	// Treat argument as filename (backward compatibility)
	filename := arg
	processSingleFile(filename)
}

// processAllOscFiles finds all .osc files in the current directory,
// extracts variables from each, and writes them to varlist files.
// It tracks unique variables across all files to ensure no duplicates.
func processAllOscFiles() error {
	files, err := filepath.Glob("*.osc")
	if err != nil {
		return fmt.Errorf("error finding osc files: %w", err)
	}

	if len(files) == 0 {
		fmt.Println("No .osc files found in current directory.")
		return nil
	}

	// Sort files alphabetically for consistent output
	sort.Strings(files)

	fmt.Printf("\n============================================\n")
	fmt.Printf("  Processing %d OSC files\n", len(files))
	fmt.Printf("============================================\n\n")

	// Track used variables to prevent duplicates across files
	usedVars := make(map[string]bool)
	usedStringVars := make(map[string]bool)
	totalVars := 0
	totalStringVars := 0
	totalReadVars := 0
	totalReadStringVars := 0

	// Process each OSC file
	for _, filename := range files {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file %s: %v\n", filename, err)
			continue
		}

		extracted, err := oscextract.ExtractVariables(file)
		file.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error extracting variables from %s: %v\n", filename, err)
			continue
		}

		// Filter unique regular variables (S.L)
		var uniqueVars []string
		for _, v := range extracted.Variables {
			if !usedVars[v] {
				uniqueVars = append(uniqueVars, v)
				usedVars[v] = true
			}
		}

		// Write unique variables to varlist
		if len(uniqueVars) > 0 {
			if err := varlistcreator.WriteVariables(filename, uniqueVars); err != nil {
				fmt.Fprintf(os.Stderr, "Error writing variables for %s: %v\n", filename, err)
			} else {
				fmt.Printf("  [W] %s: %d variables (S.L)\n", filename, len(uniqueVars))
				totalVars += len(uniqueVars)
			}
		}

		// Filter unique string variables (S.$)
		var uniqueStringVars []string
		for _, v := range extracted.StringVars {
			if !usedStringVars[v] {
				uniqueStringVars = append(uniqueStringVars, v)
				usedStringVars[v] = true
			}
		}

		// Write unique string variables to stringvarlist
		if len(uniqueStringVars) > 0 {
			if err := varlistcreator.WriteStringVariables(filename, uniqueStringVars); err != nil {
				fmt.Fprintf(os.Stderr, "Error writing string variables for %s: %v\n", filename, err)
			} else {
				fmt.Printf("  [S] %s: %d string variables (S.$)\n", filename, len(uniqueStringVars))
				totalStringVars += len(uniqueStringVars)
			}
		}

		// Display read-only variables (only informational, not written)
		readCount := len(extracted.ReadVars) + len(extracted.ReadStringVars)
		if readCount > 0 {
			fmt.Printf("  [R] %s: %d variables (L.L) | %d string variables (L.$)\n", filename, len(extracted.ReadVars), len(extracted.ReadStringVars))
			totalReadVars += len(extracted.ReadVars)
			totalReadStringVars += len(extracted.ReadStringVars)
		}

		// Handle files with no variables
		if len(uniqueVars) == 0 && len(uniqueStringVars) == 0 && readCount == 0 {
			fmt.Printf("  [-] %s: no variables found\n", filename)
		}
	}

	// Print summary
	fmt.Printf("\n============================================\n")
	fmt.Printf("  Total Written Variables (S.L):   %d\n", totalVars)
	fmt.Printf("  Total Written String Vars (S.$):  %d\n", totalStringVars)
	fmt.Printf("  Total Read Variables (L.L):       %d\n", totalReadVars)
	fmt.Printf("  Total Read String Vars (L.$):     %d\n", totalReadStringVars)
	fmt.Printf("============================================\n\n")
	return nil
}

// processSingleFile processes a single OSC file, extracts all variables,
// and writes them to varlist files. It also displays detailed information
// about which variables are written vs. read-only.
func processSingleFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	extracted, err := oscextract.ExtractVariables(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error extracting variables: %v\n", err)
		os.Exit(1)
	}

	// Display file header
	fmt.Printf("\n=== %s ===\n\n", filename)

	// Display written variables (S.L)
	fmt.Printf("  [Written Variables (S.L)]: %d\n", len(extracted.Variables))
	if len(extracted.Variables) > 0 {
		for _, v := range extracted.Variables {
			fmt.Printf("    %s\n", v)
		}
	}

	// Display written string variables (S.$)
	fmt.Printf("\n  [Written String Variables (S.$)]: %d\n", len(extracted.StringVars))
	if len(extracted.StringVars) > 0 {
		for _, v := range extracted.StringVars {
			fmt.Printf("    %s\n", v)
		}
	}

	// Display read-only variables (L.L)
	fmt.Printf("\n  [Read Variables (L.L)]: %d\n", len(extracted.ReadVars))
	if len(extracted.ReadVars) > 0 {
		for _, v := range extracted.ReadVars {
			fmt.Printf("    %s\n", v)
		}
	}

	// Display read-only string variables (L.$)
	fmt.Printf("\n  [Read String Variables (L.$)]: %d\n", len(extracted.ReadStringVars))
	if len(extracted.ReadStringVars) > 0 {
		for _, v := range extracted.ReadStringVars {
			fmt.Printf("    %s\n", v)
		}
	}

	// Write variables to output directory
	if err := varlistcreator.WriteVariables(filename, extracted.Variables); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing variables: %v\n", err)
		os.Exit(1)
	}

	if err := varlistcreator.WriteStringVariables(filename, extracted.StringVars); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing string variables: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\n  => Variables written to output/varlist/\n")
	fmt.Printf("  => String variables written to output/stringvarlisten/\n\n")
}
