package main

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
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

// createSlices takes a slice of strings and splits it into two slices. The first slice contains
// all the strings that contained <L.$.> (Stringvars) and the second slice contains all the strings that
// contained <L.L.> (Vars). The function returns the two slices as well as an error if the input slice
// is empty or the string contains no letters.
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

	return cleanSlice(stringvarlist), cleanSlice(varlist), nil
}

func main() {
	testSlice := []string{"(L.$.DasIstEineSehrLangeTestVariable)", "(L.L.DasIstEineSehrLangeTestVariable2)", "(L.$.DasIstEineSehrLangeTestVariable3)", "(L.L.DasIstEineSehrLangeTestVariable4)", "(L.L.DasIstEineSehrLangeTestVariable2)", "(L.L.DasIstEineSehrLangeTestVariable3)", "(L.L.DasIstEineSehrLangeTestVariable4)"}
	newSlice, newSlice2, err := createSlices(testSlice)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println("Varlist:")
	for _, x := range newSlice2 {
		fmt.Println(x)
	}

	fmt.Println("Stringvarlist:")
	for _, x := range newSlice {
		fmt.Println(x)
	}
}
