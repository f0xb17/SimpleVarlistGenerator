package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const (
	typeVar       = "variable"
	typeStringVar = "stringvariable"
	Pattern       = `\([LS]\.[L$]\.([a-zA-Z0-9_]+)\)`
)

var storage = make(map[Variable]bool)
var files = make([]string, 0)

type Variable struct {
	Name string
	Type string
}

// getName returns the name of the variable.
func (v Variable) getName() string {
	return v.Name
}

// getType returns the type of the variable, which is either "variable" or "stringvariable".
func (v Variable) getType() string {
	return v.Type
}

// print returns a string representation of the Variable in the format "Name: Type"
func (v Variable) print() string {
	return v.getName() + ": " + v.getType()
}

// store adds the Variable to the storage map, marking it as true.
func (v Variable) store() {
	storage[v] = true
}

// collectVariables reads a file and collects all variables in the format (L|S)\.(L|\$)([a-zA-Z0-9_]+)
// into a map. The key of the map is an object from type Variable and the value is always true.
// The function returns an error if the file path is empty or if there was an error opening the file.
func collectVariables(filePath string) (map[Variable]bool, error) {
	if filePath == "" {
		return nil, fmt.Errorf("file path is empty")
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	variables := make(map[Variable]bool)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		matches := regexp.MustCompile(Pattern).FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if len(match) > 1 {
				variable := Variable{
					Name: match[1],
					Type: typeVar,
				}
				if strings.Contains(match[0], "$") {
					variable.Type = typeStringVar
				}
				variable.store()
				duplicate := false
				for stored := range storage {
					if variable.getName() == stored.getName() {
						duplicate = true
						break
					}
					if !duplicate {
						variables[variable] = true
					}
				}
			}
		}
	}

	return variables, nil
}

// separateVariables takes a map of variables and separates them into two lists.
// The first list contains all "variable"s and the second list contains all "stringvariable"s.
// The function returns an error if the map is empty or if there is an unknown or empty type in the map.
func separateVariables(varMap map[Variable]bool) ([]string, []string, error) {
	if len(varMap) < 1 {
		return nil, nil, errors.New("map is empty")
	}

	variables := make([]string, 0, len(varMap))
	stringVariables := make([]string, 0, len(varMap))

	for val := range varMap {
		if val.getType() == "" || (val.getType() != typeVar && val.getType() != typeStringVar) {
			return nil, nil, fmt.Errorf("unknown or empty type: %s", val.getType())
		}

		if val.getType() == typeVar {
			variables = append(variables, val.getName())
		} else {
			stringVariables = append(stringVariables, val.getName())
		}
	}

	return variables, stringVariables, nil
}

// findFiles walks through the directory tree rooted at filePath and appends all paths ending
// with ".osc" to the global files slice.
func findFiles(filePath string) {
	err := filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".osc" {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("error: " + err.Error())
	}
}

// write takes two slices of strings and a filename and writes the first slice to
// fileName+_varlist.txt" and the second slice to fileName+_stringvarlist.txt".
func write(variables []string, stringVariables []string, fileName string) error {
	fullPath := filepath.Join("./output/", fileName)
	if len(variables) > 0 {
		err := os.WriteFile(fullPath+"_varlist.txt", []byte(strings.Join(variables, "\n")), 0644)
		if err != nil {
			return errors.New("error: " + err.Error())
		}
	}
	if len(stringVariables) > 0 {
		err := os.WriteFile(fullPath+"_stringvarlist.txt", []byte(strings.Join(stringVariables, "\n")), 0644)
		if err != nil {
			return errors.New("error: " + err.Error())
		}
	}
	return nil
}

func main() {
	findFiles("./test/")
	for _, file := range files {
		varMap, err := collectVariables(file)
		if err != nil {
			fmt.Println("error: " + err.Error())
		}
		varSlice, stringvarSlice, err := separateVariables(varMap)
		if err != nil {
			fmt.Println("error: " + err.Error())
		}
		write(varSlice, stringvarSlice, filepath.Base(strings.TrimSuffix(file, ".osc")))
	}
}
