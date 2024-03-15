package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

const (
	curlyOpen  rune = '{'
	curlyClose rune = '}'
	// quoteOpen rune = '"'
	// quoteClose rune = '"'
	// arrayOpen rune = '['
	// arrayClose rune = ']'
)

func Lex(input io.Reader) ([]rune, error) {
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
		fmt.Println("Invalid JSON input, file is empty. Exiting")
		os.Exit(1)
	}

	tokens := []rune(fileInput[0])
	return tokens, nil
}

func Parse(tokens []rune) {
	if tokens[0] != curlyOpen && tokens[len(tokens)-1] != curlyClose {
		fmt.Println("Invalid JSON input. Exiting.")
		os.Exit(1)
	}
	fmt.Println("Valid JSON input. Parsing...")
}

func main() {
	flag.Parse()
	var input io.Reader

	if filename := flag.Arg(0); filename != "" {
		file, err := os.Open(filename)
		if err != nil {
			fmt.Println("Error opening file: ", err)
			os.Exit(1)
		}
		defer file.Close()
		input = file
	} else {
		// If no filename is provided, use stdin
		stat, _ := os.Stdin.Stat()
		if stat.Size() > 0 {
			input = os.Stdin
		} else {
			fmt.Println("No input source provided")
			os.Exit(1)
		}
	}
	lexedString, _ := Lex(input)
	Parse(lexedString)
	fmt.Println("Parsing JSON")
}
