package main

import (
	// Uncomment this to pass the first stage
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode/utf8"
)

// Usage: echo <input_text> | your_program.sh -E <pattern>
func main() {
	if len(os.Args) < 3 || os.Args[1] != "-E" {
		fmt.Fprintf(os.Stderr, "usage: mygrep -E <pattern>\n")
		os.Exit(2) // 1 means no lines were selected, >1 means error
	}

	pattern := os.Args[2]

	line, err := io.ReadAll(os.Stdin) // assume we're only dealing with a single line
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: read input text: %v\n", err)
		os.Exit(2)
	}

	ok, err := matchLine(line, pattern)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(2)
	}

	if !ok {
		os.Exit(1)
	}

	// default exit code is 0 which means success
}

func matchLine(line []byte, pattern string) (bool, error) {
	if utf8.RuneCountInString(pattern) == 0 {
		return false, fmt.Errorf("unsupported pattern: %q", pattern)
	}

	var ok bool

	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")
	numeric := "0123456789"
	alphabetMin := makeAlphabet(optionsAlphabet{minuscle: true})
	alphabetMai := makeAlphabet(optionsAlphabet{minuscle: false})

	if pattern == "\\d" {
		ok = bytes.ContainsAny(line, numeric)
	} else if pattern == "\\w" {
		ok = bytes.ContainsAny(line, numeric+alphabetMai+alphabetMin)
	} else if strings.HasPrefix(pattern, "[") {
		positiveChars := strings.TrimSuffix(strings.TrimPrefix(pattern, "["), "]")
		ok = bytes.ContainsAny(line, positiveChars)
	} else {
		ok = bytes.ContainsAny(line, pattern)
	}

	return ok, nil
}

type optionsAlphabet struct {
	minuscle bool
}

func makeAlphabet(options optionsAlphabet) string {
	alphabet := ""
	start := 65
	end := 90

	if options.minuscle {
		start = 97
		end = 122
	}

	for i := start; i <= end; i++ {
		alphabet = alphabet + string(i)
	}

	return alphabet
}
