package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// countLines counts the number of lines in the input
func countLines(input io.Reader) (int, error) {
	scanner := bufio.NewScanner(input)
	lines := 0
	for scanner.Scan() {
		lines++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return lines, nil
}

// countWords counts the number of words in the input
func countWords(input io.Reader) (int, error) {
	scanner := bufio.NewScanner(input)
	words := 0
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return words, nil
}

// countBytes counts the number of bytes in the input
func countBytes(input io.Reader) (int, error) {
	scanner := bufio.NewScanner(input)
	bytes := 0
	scanner.Split(bufio.ScanBytes)
	for scanner.Scan() {
		bytes++
	}
	if err := scanner.Err(); err != nil {
		return 0, err
	}
	return bytes, nil
}

// countAll counts the number of lines, words and bytes in the input
func countAll(input io.ReadSeeker)(int, int, int, error) {
	scanner := bufio.NewScanner(input)
	lines, _ := countLines(input)
	input.Seek(0, 0)

	words, _ := countWords(input)
	input.Seek(0, 0)

	bytes, _ := countBytes(input)
	input.Seek(0, 0)

	if err := scanner.Err(); err != nil {
		return 0, 0, 0, err
	}
	return lines, words, bytes, nil
}

func main() {
	var line bool
	flag.BoolVar(&line, "l", false, "Count lines")
	flag.BoolVar(&line, "lines", false, "Count lines")

	var word bool
	flag.BoolVar(&word, "w", false, "Count words")
	flag.BoolVar(&word, "words", false, "Count words")

	var byte bool
	flag.BoolVar(&byte, "c", false, "Count bytes")
	flag.BoolVar(&byte, "bytes", false, "Count bytes")

	flag.Parse()

	var input io.ReadSeeker
	filename := flag.Arg(0)

	if filename != "" {
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

	if line {
		lines, err := countLines(input)
		if err != nil {
			fmt.Println("Error reading input: ", err)
			os.Exit(1)
		}
		fmt.Println(lines, filename)
	} else if word {
		words, err := countWords(input)
		if err != nil {
			fmt.Println("Error reading input: ", err)
			os.Exit(1)
		}
		fmt.Println(words, filename)
	} else if byte {
		bytes, err := countBytes(input)
		if err != nil {
			fmt.Println("Error reading input: ", err)
			os.Exit(1)
		}
		fmt.Println(bytes, filename)
	} else {
		lines, words, bytes, err := countAll(input)
		if err != nil {
			fmt.Println("Error reading input: ", err)
			os.Exit(1)
		}
		fmt.Println(lines, words, bytes, filename)
	}
}
