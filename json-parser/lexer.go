package jsonparser

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"io"
)

// Lexer reads the input file and returns a slice of strings
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