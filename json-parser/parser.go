package jsonparser

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"io"
)

const (
	curlyOpen  string = "{"
	curlyClose string = "}"
	// quote string = "\""
	// colon string = ":"
	// comma string = ","
	// arrayOpen string = "["
	// arrayClose string = ]"
)

func Lexer(input io.Reader) ([]string, error) {
	scanner := bufio.NewScanner(input)
	fileInput := []string{}
	for scanner.Scan() {
		fileInput = append(fileInput, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	// If the input is empty, return an error
	if len(fileInput) == 0 {
		msg := fmt.Sprintln("invalid json input, file is empty")
		return nil, errors.New(msg)
	}
	return fileInput, nil
}

func Parse(tokens []string) (bool, string) {
	// Check if the input is a valid JSON
	if tokens[0] != curlyOpen && tokens[len(tokens)-1] != curlyClose {
		return false, fmt.Sprintln("invalid json input")
	}
	return true, fmt.Sprintln("valid json input")
}
