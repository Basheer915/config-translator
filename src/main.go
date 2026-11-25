package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"config-translator/src/parser"
	"config-translator/src/translator"
)

const (
	defaultPermission = 0644
	exitCodeError     = 1
)

func main() {
	outputPath := flag.String("o", "result.json", "Path to output JSON file")
	flag.Parse()

	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error reading stdin:", err)
		os.Exit(exitCodeError)
	}

	program, err := parser.Parse(string(input))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Syntax error:", err)
		os.Exit(exitCodeError)
	}

	env, err := translator.Translate(program)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Translation error:", err)
		os.Exit(exitCodeError)
	}

	jsonData, err := json.MarshalIndent(env, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "JSON marshal error:", err)
		os.Exit(exitCodeError)
	}

	err = os.WriteFile(*outputPath, jsonData, defaultPermission)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error writing file:", err)
		os.Exit(exitCodeError)
	}
	fmt.Println("Successfully translated to", *outputPath)
}
