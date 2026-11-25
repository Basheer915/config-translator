package parser

type Program struct {
	Declarations []*Declaration `parser:"@@*"`
}

type Declaration struct {
	Name  string `parser:"'var' @Ident Assign"`
	Value *Value `parser:"@@"`
}

type Value struct {
	Number     *string     `parser:"  @Number"`
	String     *string     `parser:"| @String"`
	Array      *Array      `parser:"| @@"`
	Expression *Expression `parser:"| Hash LBrace @@ RBrace"`
}

type Array struct {
	Values []*Value `parser:"LBracket @@* RBracket"`
}

type Expression struct {
	Left *Operand     `parser:"@@"`
	Tail []*OpOperand `parser:"@@*"`
}

type OpOperand struct {
	Op    string   `parser:"@(Plus | Minus)"`
	Right *Operand `parser:"@@"`
}

type Operand struct {
	Number *string     `parser:"  @Number"`
	String *string     `parser:"| @String"`
	Call   *Function   `parser:"| @@"`
	Name   *string     `parser:"| @Ident"`
	Nested *Expression `parser:"| LParen @@ RParen"`
}

type Function struct {
	Name string        `parser:"@Ident LParen"`
	Args []*Expression `parser:"@@ (Comma @@)* RParen"`
}
