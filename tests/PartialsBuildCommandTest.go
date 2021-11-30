package main

import (
	"fmt"
	"os"
	"partey/src"
	"path"
)

func assertStringsMatch(expected string, actual string) {
	if actual != expected {
		fmt.Printf("Failing because expected \n%s\nDoes not match actual\n%s", expected, actual)
		panic("Ending tests!")
	}
}

func main() {
	// Given
	dir, _ := os.MkdirTemp("", "testing")
	partialsDir, _ := os.MkdirTemp(dir, "partials")
	aggregateFile := path.Join(dir, "agg")

	defer os.RemoveAll(dir)

	_ = os.WriteFile(path.Join(partialsDir, "partials1"), []byte("Partial 1"), 0644)
	_ = os.WriteFile(path.Join(partialsDir, "partials2"), []byte("Partial 2"), 0644)
	_ = os.WriteFile(aggregateFile, []byte(nil), 0644)

	// When
	command := src.NewPartialsBuildCommand(aggregateFile, partialsDir, "#")
	command.Run()

	// Then
	actual, _ := os.ReadFile(aggregateFile)
	expected := "# PARTIALS>>>>>\nPartial 1\nPartial 2\n# PARTIALS<<<<<\n"

	assertStringsMatch(expected, string(actual))

	println("PartialsBuildCommandTest.go passed!")
}
