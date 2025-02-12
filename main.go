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

type Variable struct {
	Name string
	Type string
}

func (v Variable) getName() string {
	return v.Name
}

func (v Variable) getType() string {
	return v.Type
}

func (v Variable) print() string {
	return v.getName() + ": " + v.getType()
}

func collectVariables(filePath string) (map[Variable]bool, error) {
	if filePath == "" {
		return nil, errors.New("file path is empty")
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
				variables[variable] = true
			}
		}
	}

	return variables, nil
}

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
}
