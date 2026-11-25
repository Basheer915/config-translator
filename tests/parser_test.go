package tests

import (
	"testing"

	"config-translator/src/parser"
)

func TestParseValidConfigurations(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "number and string",
			input: "var port = 8080.0\nvar host = q(localhost)",
		},
		{
			name:  "trailing dot number",
			input: "var count = 5.",
		},
		{
			name:  "nested arrays",
			input: "var data = [ q(first) [ 1.0 2.0 ] ]",
		},
		{
			name:  "chained arithmetic expression",
			input: "var result = #{10.0 + 5.0 - 2.0}",
		},
		{
			name:  "nested function call",
			input: "var result = #{max(10.0 + 5.0, 20.0)}",
		},
		{
			name:  "concat function",
			input: "var result = #{concat(q(hello), q(world))}",
		},
		{
			name:  "comments",
			input: "-- configuration comment\nvar port = 8080.0 -- inline comment",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := parser.Parse(test.input)
			if err != nil {
				t.Fatalf("expected valid input, got error: %v", err)
			}
		})
	}
}

func TestParseInvalidConfigurations(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{
			name:  "integer without decimal point",
			input: "var port = 8080",
		},
		{
			name:  "missing value",
			input: "var port =",
		},
		{
			name:  "invalid identifier",
			input: "var 123port = 8080.0",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			_, err := parser.Parse(test.input)
			if err == nil {
				t.Fatal("expected parsing error, got nil")
			}
		})
	}
}
