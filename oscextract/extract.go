// Package oscextract provides functionality to extract variable names from OSC (OpenSCAD) files.
// It identifies different types of variables based on their prefix:
//   - S.L: Written regular variables (saved locally)
//   - L.L: Read-only regular variables (loaded from memory)
//   - S.$: Written string variables
//   - L.$: Read-only string variables
package oscextract

import (
	"bufio"
	"io"
	"regexp"
	"sort"
	"strings"
)

// Variable represents a variable with a name and type extracted from an OSC file.
// Note: This struct is currently not used but kept for potential future extensions.
type Variable struct {
	Name string
	Type string
}

// ExtractedVars holds the extracted variables categorized by their type and access mode.
type ExtractedVars struct {
	// Variables contains regular variables that are written (S.L prefix)
	Variables []string
	// StringVars contains string variables that are written (S.$ prefix)
	StringVars []string
	// ReadVars contains regular variables that are read-only (L.L prefix)
	ReadVars []string
	// ReadStringVars contains string variables that are read-only (L.$ prefix)
	ReadStringVars []string
}

// ExtractVariables reads from an io.Reader and extracts variable names that match
// the OSC variable patterns. It categorizes variables based on their prefix:
//
//	(L.L.variable) - Read-only regular variable
//	(S.L.variable) - Written regular variable
//	(L.$.variable) - Read-only string variable
//	(S.$.variable) - Written string variable
//
// The function returns an ExtractedVars struct containing four slices of sorted,
// unique variable names, or an error if reading fails.
func ExtractVariables(r io.Reader) (*ExtractedVars, error) {
	// Regex pattern to match OSC variable patterns: (L.L., S.L., L.$, S.$.) followed by variable name
	re := regexp.MustCompile(`\((L\.L|S\.L|L\.\$|S\.\$)\.([a-zA-Z_][a-zA-Z0-9_]*)\)`)

	// Separate maps for each category
	vars := make(map[string]struct{})
	stringVars := make(map[string]struct{})
	readVars := make(map[string]struct{})
	readStringVars := make(map[string]struct{})

	// Scan the input line by line
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if len(match) == 3 {
				prefix := match[1]
				name := match[2]

				// Categorize based on prefix
				if prefix == "L.$" || prefix == "S.$" {
					// String variables
					if prefix == "L.$" {
						readStringVars[name] = struct{}{}
					} else {
						stringVars[name] = struct{}{}
					}
				} else {
					// Regular variables
					if prefix == "L.L" {
						readVars[name] = struct{}{}
					} else {
						vars[name] = struct{}{}
					}
				}
			}
		}
	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// Build result with sorted variable names
	result := &ExtractedVars{
		Variables:      sortMapKeys(vars),
		StringVars:     sortMapKeys(stringVars),
		ReadVars:       sortMapKeys(readVars),
		ReadStringVars: sortMapKeys(readStringVars),
	}

	return result, nil
}

// sortMapKeys converts a map's keys to a sorted slice of strings.
// This is used to ensure consistent, alphabetical ordering of extracted variables.
func sortMapKeys(m map[string]struct{}) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return strings.ToLower(keys[i]) < strings.ToLower(keys[j])
	})
	return keys
}
