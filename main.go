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

// cleanSlice removes duplicate variable names from the input slice by extracting
// the "real" variable name using convertVariable and returns a slice of unique
// variable names.
//
//go:deprecated
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

// createSlices takes a slice of strings and splits it into two slices. The first slice contains
// all the strings that contained <L.$.> and <S.$.> (Stringvars) and the second slice contains all the strings that
// contained <L.L.> and <S.L.> (Vars). The function returns the two slices as well as an error if the input slice
// is empty or the string contains no letters.
//
//go:deprecated
func createSlices(varSlice []string) ([]string, []string, error) {
	stringvarlist := []string{}
	varlist := []string{}

	if len(varSlice) == 0 {
		return nil, nil, errors.New("error while creating slices. input slice was empty")
	}

	for _, v := range varSlice {
		if v == "" {
			return nil, nil, errors.New("string contains no letters. Is it empty")
		}

		if strings.Contains(v, "$") {
			stringvarlist = append(stringvarlist, v)
		} else {
			varlist = append(varlist, v)
		}

	}

	return cleanSlice(varlist), cleanSlice(stringvarlist), nil
}

// readFile takes a file path and reads the file line by line. It uses a regular expression to extract
// all the strings that contain <L.L.> or <L.$.> and <S.L.> or <S.$.> and returns them in a slice. If the function encounters
// an error while reading the file, it will return an empty slice and an error.
//
//go:deprecated
func readFile(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, errors.New("something went wrong")
	}
	variables := []string{}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		matches := regexp.MustCompile(getPattern()).FindAllStringSubmatch(line, -1)
		for _, match := range matches {
			if len(match) > 1 {
				variables = append(variables, match[0])
			}
		}
	}

	return variables, nil
}

func main() {
	slice, err := readFile("./MAN_SG_Dash.osc")

	if err != nil {
		fmt.Println(err)
	}

	variables, stringvariables, err := createSlices(slice)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Variables:")
	for _, varVal := range variables {
		fmt.Println(varVal)
	}

	fmt.Println("Stringvariables:")
	for _, stringvarVal := range stringvariables {
		fmt.Println(stringvarVal)
	}
}
