package evaluator

import (
	"fmt"

	"github.com/WilhelmWeber/go_ruby/src/parser"
)

func Evaluator(tree *parser.Node, env map[string]int) int {
	switch tree.Kind {
	case parser.ND_ASN:
		//変数代入の場合は必ず左辺が識別子になっているはずなので
		ident := tree.Lhs.Str
		right := Evaluator(tree.Rhs, env)
		env[ident] = right
		return right
	case parser.ND_REF:
		return env[tree.Str]
	case parser.ND_NUM:
		return tree.Val
	case parser.ND_ADD:
		left := Evaluator(tree.Lhs, env)
		right := Evaluator(tree.Rhs, env)
		return left + right
	case parser.ND_SUB:
		left := Evaluator(tree.Lhs, env)
		right := Evaluator(tree.Rhs, env)
		return left - right
	case parser.ND_MUL:
		left := Evaluator(tree.Lhs, env)
		right := Evaluator(tree.Rhs, env)
		return left * right
	case parser.ND_DIV:
		left := Evaluator(tree.Lhs, env)
		right := Evaluator(tree.Rhs, env)
		return int(left / right)
	case parser.ND_EQU:
		left := Evaluator(tree.Lhs, env)
		right := Evaluator(tree.Rhs, env)
		if left == right {
			return 1
		} else {
			return 0
		}
	case parser.ND_UEQ:
		left := Evaluator(tree.Lhs, env)
		right := Evaluator(tree.Rhs, env)
		if left != right {
			return 1
		} else {
			return 0
		}
	case parser.ND_ELS:
		left := Evaluator(tree.Lhs, env)
		right := Evaluator(tree.Rhs, env)
		if left <= right {
			return 1
		} else {
			return 0
		}
	case parser.ND_LES:
		left := Evaluator(tree.Lhs, env)
		right := Evaluator(tree.Rhs, env)
		if left < right {
			return 1
		} else {
			return 0
		}
	case parser.ND_EMR:
		left := Evaluator(tree.Lhs, env)
		right := Evaluator(tree.Rhs, env)
		if left >= right {
			return 1
		} else {
			return 0
		}
	case parser.ND_MOR:
		left := Evaluator(tree.Lhs, env)
		right := Evaluator(tree.Rhs, env)
		if left > right {
			return 1
		} else {
			return 0
		}
	case parser.ND_IF:
		if Evaluator(tree.If, env) == 1 {
			return Evaluator(tree.Then, env)
		} else {
			if tree.Then != nil {
				return Evaluator(tree.Else, env)
			} else {
				return 0
			}
		}
	case parser.ND_WHL:
		for Evaluator(tree.If, env) == 1 {
			Evaluator(tree.Then, env)
		}
		return 0
	case parser.ND_BLC:
		for _, e := range tree.Stmts {
			Evaluator(e, env)
		}
		return 0
	case parser.ND_FNC:
		//まだprintlnだけを実装
		value := Evaluator(tree.FncArg, env)
		fmt.Println(value)
		return 0
	default:
		panic("Cannnot Evaluate")
	}
}

func SenteceEval(trees []*parser.Node) {
	env := make(map[string]int)
	for _, e := range trees {
		Evaluator(e, env)
	}
	fmt.Println("Evaluate ended with success")
}
