// SimpleVarlistGenerator is a tool to extract variable names from OSC files
// and write them to a list file.
package main

import (
	"fmt"
	"os"

	"SimpleVarlistGenerator/oscextract"
	"SimpleVarlistGenerator/varlistcreator"
)

// main reads an OSC file, extracts variable names, and writes them to a varlist file.
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file.osc>")
		os.Exit(1)
	}

	filename := os.Args[1]
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	vars, err := oscextract.ExtractVariables(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error extracting variables: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Found %d unique variables:\n\n", len(vars))
	for _, v := range vars {
		fmt.Println(v)
	}

	if err := varlistcreator.WriteVariables(filename, vars); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing variables: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("\nVariables written to file.\n")
}
