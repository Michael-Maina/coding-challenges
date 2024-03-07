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

func main() {
	linePtr := flag.Bool("l", false, "Count lines")
	wordPtr := flag.Bool("w", false, "Count words")

	flag.Parse()

	var input io.Reader
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
		if (stat.Mode() & os.ModeCharDevice) == 0 {
			input = os.Stdin
		} else {
			fmt.Println("No input source provided")
			os.Exit(1)
		}
	}

	if flag.NFlag() == 0 {
		fmt.Println("wc requires one flag to work, '-l' or '-w'")
		return
	}

	if *linePtr {
		lines, err := countLines(input)
		if err != nil {
			fmt.Println("Error reading input: ", err)
			os.Exit(1)
		}
		fmt.Println(lines, filename)
	} else if *wordPtr {
		words, err := countWords(input)
		if err != nil {
			fmt.Println("Error reading input: ", err)
			os.Exit(1)
		}
		fmt.Println(words, filename)
	}
}
