package jsonparser

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"io"
)

const (
	curlyOpen  rune = '{'
	curlyClose rune = '}'
	// quoteOpen rune = '"'
	// quoteClose rune = '"'
	// arrayOpen rune = '['
	// arrayClose rune = ']'
)

func Lexer(input io.Reader) ([]rune, error) {
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
		return []rune{}, errors.New(msg)
	}
	tokens := []rune(fileInput[0])
	return tokens, nil
}

func Parse(tokens []rune) (bool, string) {
	// Check if the input is a valid JSON
	if tokens[0] != curlyOpen && tokens[len(tokens)-1] != curlyClose {
		return false, fmt.Sprintln("invalid json input")
	}
	return true, fmt.Sprintln("valid json input")
}
