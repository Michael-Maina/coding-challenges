package jsonparser

import (
	"bufio"
	"fmt"
	"io"
	"strings"

	"github.com/pkg/errors"
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
	if len(tokens[0]) > 1 {
		if !strings.HasPrefix(tokens[0], curlyOpen) && !strings.HasSuffix(tokens[0], curlyClose) {
			return false, fmt.Sprintln("invalid json input")
			} else {
			return true, fmt.Sprintln("valid json input")
		}
	}

	if tokens[0] != curlyOpen && tokens[len(tokens)-1] != curlyClose {
		return false, fmt.Sprintln("invalid json input")
	}
	return true, fmt.Sprintln("valid json input")
}
