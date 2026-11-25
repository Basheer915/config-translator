package tests

import (
	"testing"

	"config-translator/src/parser"
	"config-translator/src/translator"
)

func TestTranslate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected interface{}
		key      string
	}{
		{
			name:     "math",
			input:    "var result = #{10.0 + 5.0 - 3.0}",
			expected: 12.0,
			key:      "result",
		},
		{
			name:     "max function",
			input:    "var result = #{max(10.0, 20.0)}",
			expected: 20.0,
			key:      "result",
		},
		{
			name:     "concat function",
			input:    "var result = #{concat(q(foo), q(bar))}",
			expected: "foobar",
			key:      "result",
		},
		{
			name:     "variable reference",
			input:    "var a = 5.0\nvar b = #{a + 1.0}",
			expected: 6.0,
			key:      "b",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			prog, err := parser.Parse(test.input)
			if err != nil {
				t.Fatalf("parse error: %v", err)
			}
			env, err := translator.Translate(prog)
			if err != nil {
				t.Fatalf("translate error: %v", err)
			}
			if env[test.key] != test.expected {
				t.Errorf("expected %v, got %v", test.expected, env[test.key])
			}
		})
	}
}
