package jsonparser

import (
	"fmt"
	"strings"
	"strconv"
)

const (
	curlyOpen  string = "{"
	curlyClose string = "}"
	quote      string = "\""
	colon      string = ":"
	colonSpace string = ": "
	comma      string = ","
	// arrayOpen string = "["
	// arrayClose string = ]"
)

// Parse checks if the input is a valid JSON
func Parse(tokens []string) (bool, string) {
	// Check if the single line input is a valid JSON
	if len(tokens[0]) > 1 {
		return parseSingle(tokens)
	}

	// Check if multi-line input is a valid JSON
	return parseMulti(tokens)
}

// Parse single line JSON
func parseSingle(tokens []string) (bool, string) {
	if !strings.HasPrefix(tokens[0], curlyOpen) && !strings.HasSuffix(tokens[0], curlyClose) {
		return false, fmt.Sprintln("invalid json input")
	}
	tokens[0], _ = strings.CutPrefix(tokens[0], curlyOpen)  // Delete curly brackets
	tokens[0], _ = strings.CutSuffix(tokens[0], curlyClose) // Delete curly brackets

	var splitTokens []string

	splitTokens = append(splitTokens, strings.SplitAfter(tokens[0], ",")...)
	if splitTokens[len(splitTokens)-1] == "" {
		return false, fmt.Sprintln("invalid json input")
	}

	for index, token := range splitTokens {
		if strings.HasSuffix(token, comma) {
			splitTokens[index] = strings.TrimSuffix(token, comma)
		}

		key, _ := strings.CutSuffix(token, colon)
		if isValid, msg := parseKey(key); !isValid{
			return false, fmt.Sprintln(msg)
		}

		value, _ := strings.CutPrefix(token, colonSpace)
		if isValid, msg := parseValue(value); !isValid {
			return false, fmt.Sprintln(msg)
		}
	}

	return true, fmt.Sprintln("valid json input")
}

// Parse multi-line JSON
func parseMulti(tokens []string) (bool, string) {
	if tokens[0] != curlyOpen && tokens[len(tokens)-1] != curlyClose {
		return false, fmt.Sprintln("invalid json input")
	}

	tokens = tokens[1 : len(tokens)-1] // Delete curly brackets

	var commaCount  int

	for _, token := range tokens {
		token, commaPresent := strings.CutSuffix(token, comma)
		if commaPresent {
			commaCount++
		}

		// Check if key has double quotation marks
		keyValue := strings.Split(token, colonSpace)
		if isValid, msg := parseKey(keyValue[0]); !isValid{
			return false, fmt.Sprintln(msg)
		}
		if isValid, msg := parseValue(keyValue[1]); !isValid {
			return false, fmt.Sprintln(msg)
		}
	}

	if len(tokens) != commaCount+1 {
		return false, fmt.Sprintln("invalid json input")
	}

	return true, fmt.Sprintln("valid json input")
}

// Check if the key is a string
func parseKey(key string) (bool, string) {
	if strings.Count(key, quote) != 2 {
		return false, fmt.Sprintln("invalid json input")
	}
	return true, fmt.Sprintln("valid json input")
}

// Check if the value is a string, int, or boolean
func parseValue(value string) (bool, string) {
	if strings.Count(value, quote) == 2 {
		return true, fmt.Sprintln("valid json input")
	}

	if _, err := strconv.Atoi(value); err == nil {
		return true, fmt.Sprintln("valid json input")
	}

	if value == "true" || value == "false" {
		return true, fmt.Sprintln("valid json input")
	}

	return false, fmt.Sprintln("invalid json input")
}
