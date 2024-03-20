package jsonparser

import (
	"bufio"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"strings"
)

const (
	curlyOpen  string = "{"
	curlyClose string = "}"
	quote      string = "\""
	colon      string = ":"
	comma      string = ","
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
	// Check if the single line input is a valid JSON
	if len(tokens[0]) > 1 {
		return parseKeySingle(tokens)
	}

	// Check if multi-line input is a valid JSON
	return parseKeyMulti(tokens)
}

func parseKeySingle(tokens []string) (bool, string) {
	if !strings.HasPrefix(tokens[0], curlyOpen) && !strings.HasSuffix(tokens[0], curlyClose) {
		return false, fmt.Sprintln("invalid json input")
	}

	return true, fmt.Sprintln("valid json input")
}

func parseKeyMulti(tokens []string) (bool, string) {
	if tokens[0] != curlyOpen && tokens[len(tokens) - 1] != curlyClose {
		return false, fmt.Sprintln("invalid json input")
	}

	tokens = tokens[1 : len(tokens) - 1] // Delete curly brackets

	var (
		splitTokens [][]string
		commaCount  int
	)

	for _, token := range tokens {
		token, commaPresent := strings.CutSuffix(token, comma)
		if commaPresent {
			commaCount++
		}
		keyValue := strings.Split(token, ": ")
		splitTokens = append(splitTokens, keyValue)
	}

	if len(tokens) != commaCount+1 {
		return false, fmt.Sprintln("invalid json input")
	}

	// Check if key has double quotation marks
	for _, pair := range splitTokens {
		if strings.Count(pair[0], quote) != 2 {
			return false, fmt.Sprintln("invalid json input")
		}
	}

	return true, fmt.Sprintln("valid json input")
}
