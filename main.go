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

func main() {
	fmt.Println("Hello, World!")
}
