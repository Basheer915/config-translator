package parser

import (
	"github.com/alecthomas/participle/v2"
	"github.com/alecthomas/participle/v2/lexer"
)

var configLexer = lexer.MustSimple([]lexer.SimpleRule{
	{Name: "Comment", Pattern: `--[^\n]*`},
	{Name: "Number", Pattern: `[0-9]+\.[0-9]*`},
	{Name: "String", Pattern: `q\([^)]*\)`},
	{Name: "Ident", Pattern: `[a-zA-Z][_a-zA-Z0-9]*`},
	{Name: "Plus", Pattern: `\+`},
	{Name: "Minus", Pattern: `-`},
	{Name: "Comma", Pattern: `,`},
	{Name: "LParen", Pattern: `\(`},
	{Name: "RParen", Pattern: `\)`},
	{Name: "LBracket", Pattern: `\[`},
	{Name: "RBracket", Pattern: `\]`},
	{Name: "LBrace", Pattern: `\{`},
	{Name: "RBrace", Pattern: `\}`},
	{Name: "Assign", Pattern: `=`},
	{Name: "Hash", Pattern: `#`},
	{Name: "Whitespace", Pattern: `\s+`},
})

var configParser = participle.MustBuild[Program](
	participle.Lexer(configLexer),
	participle.Elide("Comment", "Whitespace"),
)

func Parse(input string) (*Program, error) {
	return configParser.ParseString("", input)
}
