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

func (v Variable) convert() string {
	match := regexp.MustCompile(Pattern).FindStringSubmatch(v.getName())
	return match[1]
}

func collectVariables(filePath string) (map[Variable]bool, error) {
	if filePath == "" {
		return nil, errors.New("error: file path is empty! Aborting")
	}

	file, err := os.Open(filePath)

	if err != nil {
		return nil, errors.New("error: %w" + err.Error())
	}

	variables := make(map[Variable]bool)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		matches := regexp.MustCompile(Pattern).FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if len(match) > 1 {
				var variable Variable
				if strings.Contains(match[0], "$") {
					variable = Variable{Name: match[0], Type: typeStringVar}
				} else {
					variable = Variable{Name: match[0], Type: typeVar}
				}
				variable = Variable{Name: variable.convert(), Type: variable.getType()}
				variables[variable] = true
			}
		}
	}

	return variables, nil
}

func main() {
	slice, err := collectVariables("./D1556.osc")
	if err != nil {
		fmt.Println("error: " + err.Error())
	}

	for x := range slice {
		fmt.Println(x.getName())
	}
}
