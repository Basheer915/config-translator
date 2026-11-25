package translator

import (
	"fmt"
	"strconv"

	"config-translator/src/parser"
)

const (
	expectedArgsCount = 2
	stringPrefixLen   = 2
	floatBitSize      = 64
)

type Environment map[string]interface{}

func Translate(program *parser.Program) (Environment, error) {
	env := make(Environment)
	for _, decl := range program.Declarations {
		val, err := evalValue(env, decl.Value)
		if err != nil {
			return nil, fmt.Errorf("declaration %s: %w", decl.Name, err)
		}
		env[decl.Name] = val
	}
	return env, nil
}

func evalValue(env Environment, val *parser.Value) (interface{}, error) {
	if val.Number != nil {
		return strconv.ParseFloat(*val.Number, floatBitSize)
	}
	if val.String != nil {
		s := *val.String
		return s[stringPrefixLen : len(s)-1], nil
	}
	if val.Array != nil {
		return evalArray(env, val.Array)
	}
	if val.Expression != nil {
		return evalExpression(env, val.Expression)
	}
	return nil, fmt.Errorf("unknown value type")
}

func evalArray(env Environment, array *parser.Array) ([]interface{}, error) {
	arr := make([]interface{}, 0, len(array.Values))
	for _, v := range array.Values {
		ev, err := evalValue(env, v)
		if err != nil {
			return nil, err
		}
		arr = append(arr, ev)
	}
	return arr, nil
}

func evalExpression(env Environment, expr *parser.Expression) (interface{}, error) {
	acc, err := evalOperand(env, expr.Left)
	if err != nil {
		return nil, err
	}
	for _, step := range expr.Tail {
		right, err := evalOperand(env, step.Right)
		if err != nil {
			return nil, err
		}
		acc, err = applyOperator(step.Op, acc, right)
		if err != nil {
			return nil, err
		}
	}
	return acc, nil
}

func applyOperator(op string, left, right interface{}) (interface{}, error) {
	l, ok1 := left.(float64)
	r, ok2 := right.(float64)
	if !ok1 || !ok2 {
		return nil, fmt.Errorf("math requires numbers")
	}
	switch op {
	case "+":
		return l + r, nil
	case "-":
		return l - r, nil
	default:
		return nil, fmt.Errorf("unknown operator: %s", op)
	}
}

func evalOperand(env Environment, op *parser.Operand) (interface{}, error) {
	if op.Number != nil {
		return strconv.ParseFloat(*op.Number, floatBitSize)
	}
	if op.String != nil {
		s := *op.String
		return s[stringPrefixLen : len(s)-1], nil
	}
	if op.Name != nil {
		val, exists := env[*op.Name]
		if !exists {
			return nil, fmt.Errorf("undefined variable: %s", *op.Name)
		}
		return val, nil
	}
	if op.Call != nil {
		return evalFunction(env, op.Call)
	}
	if op.Nested != nil {
		return evalExpression(env, op.Nested)
	}
	return nil, fmt.Errorf("unknown operand")
}

func evalFunction(env Environment, fn *parser.Function) (interface{}, error) {
	args := make([]interface{}, 0, len(fn.Args))
	for _, a := range fn.Args {
		v, err := evalExpression(env, a)
		if err != nil {
			return nil, err
		}
		args = append(args, v)
	}
	if len(args) != expectedArgsCount {
		return nil, fmt.Errorf("%s requires %d args", fn.Name, expectedArgsCount)
	}
	switch fn.Name {
	case "max":
		return evalMax(args)
	case "concat":
		return evalConcat(args)
	default:
		return nil, fmt.Errorf("unknown function: %s", fn.Name)
	}
}

func evalMax(args []interface{}) (interface{}, error) {
	a, ok1 := args[0].(float64)
	b, ok2 := args[1].(float64)
	if !ok1 || !ok2 {
		return nil, fmt.Errorf("max requires numbers")
	}
	if a > b {
		return a, nil
	}
	return b, nil
}

func evalConcat(args []interface{}) (interface{}, error) {
	a, ok1 := args[0].(string)
	b, ok2 := args[1].(string)
	if !ok1 || !ok2 {
		return nil, fmt.Errorf("concat requires strings")
	}
	return a + b, nil
}
