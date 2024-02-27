package evaluator

import "github.com/WilhelmWeber/go_ruby/src/parser"

func Evaluator(tree *parser.Node) int {
	switch tree.Kind {
	case parser.ND_NUM:
		return tree.Val
	case parser.ND_ADD:
		left := Evaluator(tree.Lhs)
		right := Evaluator(tree.Rhs)
		return left + right
	case parser.ND_SUB:
		left := Evaluator(tree.Lhs)
		right := Evaluator(tree.Rhs)
		return left - right
	case parser.ND_MUL:
		left := Evaluator(tree.Lhs)
		right := Evaluator(tree.Rhs)
		return left * right
	default:
		left := Evaluator(tree.Lhs)
		right := Evaluator(tree.Rhs)
		return int(left / right)
	}
}
