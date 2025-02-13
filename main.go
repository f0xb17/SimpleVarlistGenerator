package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strings"
)

const (
	typeVar       = "variable"
	typeStringVar = "stringvariable"
	Pattern       = `\([LS]\.[L$]\.([a-zA-Z0-9_]+)\)`
)

var storage = make(map[Variable]bool)

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
				for stored := range storage {
					if stored != variable {
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
func separateVariables(varMap map[Variable]bool) ([]Variable, []Variable, error) {
	if len(varMap) < 1 {
		return nil, nil, errors.New("map is empty")
	}

	variables := make([]Variable, 0, len(varMap))
	stringVariables := make([]Variable, 0, len(varMap))

	for val := range varMap {
		if val.getType() == "" || (val.getType() != typeVar && val.getType() != typeStringVar) {
			return nil, nil, fmt.Errorf("unknown or empty type: %s", val.getType())
		}

		if val.getType() == typeVar {
			variables = append(variables, val)
		} else {
			stringVariables = append(stringVariables, val)
		}
	}

	return variables, stringVariables, nil
}

func main() {
	varMap, err := collectVariables("./MAN_SG_Dash.osc")
	if err != nil {
		fmt.Println("error: " + err.Error())
	}

	for x := range varMap {
		fmt.Println(x.print())
	}

	varSlice, stringvarSlice, err := separateVariables(varMap)
	if err != nil {
		fmt.Println("error: " + err.Error())
	}

	fmt.Println(varSlice)
	fmt.Println(stringvarSlice)

	for stored := range storage {
		fmt.Println(stored)
	}
}
