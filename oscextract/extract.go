// Package oscextract provides functionality to extract variable names from OSC files.
package oscextract

import (
	"bufio"
	"io"
	"regexp"
	"sort"
	"strings"
)

// Variable represents a variable with a name and type extracted from an OSC file.
type Variable struct {
	Name string
	Type string
}

// ExtractVariables reads from an io.Reader and extracts unique variable names
// that match the OSC variable pattern (L.L., S.L., L.$, S.$).
// It returns a sorted slice of unique variable names or an error.
func ExtractVariables(r io.Reader) ([]string, error) {
	re := regexp.MustCompile(`\((L\.L|S\.L|L\.\$|S\.\$)\.([a-zA-Z_][a-zA-Z0-9_]*)\)`)

	vars := make(map[string]struct{})

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if len(match) == 3 {
				vars[match[2]] = struct{}{}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	sortedVars := make([]string, 0, len(vars))
	for v := range vars {
		sortedVars = append(sortedVars, v)
	}
	sort.Slice(sortedVars, func(i, j int) bool {
		return strings.ToLower(sortedVars[i]) < strings.ToLower(sortedVars[j])
	})

	return sortedVars, nil
}
