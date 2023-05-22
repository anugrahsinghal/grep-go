package main

import (
	"fmt"
	"io"
	"os"
	"regexp"
)

// Usage: echo <input_text> | your_grep.sh -E <pattern>
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
	// to ensure only exactly 1 character is provided in the pattern
	//if utf8.RuneCountInString(pattern) != 1 {
	//	return false, fmt.Errorf("unsupported pattern: %q", pattern)
	//}

	var ok bool

	println(pattern)
	compiledRegex := regexp.MustCompile(pattern)
	ok = compiledRegex.Match(line)

	// You can use print statements as follows for debugging, they'll be visible when running tests.
	fmt.Println("Logs from your program will appear here!")

	// Uncomment this to pass the first stage
	//ok = bytes.ContainsAny(line, pattern)

	return ok, nil
}
