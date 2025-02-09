package main

import (
	"fmt"
	"regexp"
)

// convertVariable takes a variable name and returns the part of the name after the <L.L.> or <L.$.>.
// part. It is used to extract the "real" variable name.
func convertVariable(varName string) string {
	match := regexp.MustCompile(`\(L\.([L$])\.([a-zA-Z0-9_]+)\)`).FindStringSubmatch(varName)
	return match[2]
}

// cleanSlice removes duplicate variable names from the input slice by extracting
// the "real" variable name using convertVariable and returns a slice of unique
// variable names.
func cleanSlice(varSlice []string) []string {
	cleaned := make(map[string]bool)
	cleanedSlice := []string{}

	for _, v := range varSlice {
		varName := convertVariable(v)
		cleaned[varName] = true
	}

	for v := range cleaned {
		cleanedSlice = append(cleanedSlice, v)
	}

	return cleanedSlice
}

func main() {
	fmt.Println("Hello, World!")
}
