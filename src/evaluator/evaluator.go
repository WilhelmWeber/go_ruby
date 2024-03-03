package evaluator

import (
	"fmt"

	"github.com/WilhelmWeber/go_ruby/src/parser"
)

type fnkind int

const (
	builtin fnkind = iota
	userdefined
)

type fn struct {
	kind fnkind
	args []string     //引数
	flow *parser.Node //実行する処理
}

func Evaluator(tree *parser.Node, env map[string]int, genv map[string]fn) int {
	switch tree.Kind {
	case parser.ND_ASN:
		//変数代入の場合は必ず左辺が識別子になっているはずなので
		ident := tree.Lhs.Str
		right := Evaluator(tree.Rhs, env, genv)
		env[ident] = right
		return right
	case parser.ND_REF:
		return env[tree.Str]
	case parser.ND_NUM:
		return tree.Val
	case parser.ND_ADD:
		left := Evaluator(tree.Lhs, env, genv)
		right := Evaluator(tree.Rhs, env, genv)
		return left + right
	case parser.ND_SUB:
		left := Evaluator(tree.Lhs, env, genv)
		right := Evaluator(tree.Rhs, env, genv)
		return left - right
	case parser.ND_MUL:
		left := Evaluator(tree.Lhs, env, genv)
		right := Evaluator(tree.Rhs, env, genv)
		return left * right
	case parser.ND_DIV:
		left := Evaluator(tree.Lhs, env, genv)
		right := Evaluator(tree.Rhs, env, genv)
		return int(left / right)
	case parser.ND_EQU:
		left := Evaluator(tree.Lhs, env, genv)
		right := Evaluator(tree.Rhs, env, genv)
		if left == right {
			return 1
		} else {
			return 0
		}
	case parser.ND_UEQ:
		left := Evaluator(tree.Lhs, env, genv)
		right := Evaluator(tree.Rhs, env, genv)
		if left != right {
			return 1
		} else {
			return 0
		}
	case parser.ND_ELS:
		left := Evaluator(tree.Lhs, env, genv)
		right := Evaluator(tree.Rhs, env, genv)
		if left <= right {
			return 1
		} else {
			return 0
		}
	case parser.ND_LES:
		left := Evaluator(tree.Lhs, env, genv)
		right := Evaluator(tree.Rhs, env, genv)
		if left < right {
			return 1
		} else {
			return 0
		}
	case parser.ND_EMR:
		left := Evaluator(tree.Lhs, env, genv)
		right := Evaluator(tree.Rhs, env, genv)
		if left >= right {
			return 1
		} else {
			return 0
		}
	case parser.ND_MOR:
		left := Evaluator(tree.Lhs, env, genv)
		right := Evaluator(tree.Rhs, env, genv)
		if left > right {
			return 1
		} else {
			return 0
		}
	case parser.ND_IF:
		if Evaluator(tree.If, env, genv) == 1 {
			return Evaluator(tree.Then, env, genv)
		} else {
			if tree.Then != nil {
				return Evaluator(tree.Else, env, genv)
			} else {
				return 0
			}
		}
	case parser.ND_WHL:
		for Evaluator(tree.If, env, genv) == 1 {
			Evaluator(tree.Then, env, genv)
		}
		return 0
	case parser.ND_BLC:
		//グローバル変数はブロック内でも参照されるが、ブロック内のローカル変数はブロックをぬけると参照できない
		new_env := env
		var a int
		for _, e := range tree.Stmts {
			a = Evaluator(e, new_env, genv)
		}
		return a
	case parser.ND_DEF:
		//関数定義
		genv[tree.Str] = fn{kind: userdefined, args: tree.Args, flow: tree.Flow}
		return 0
	case parser.ND_FNC:
		//まだprintlnだけを実装
		var args []int
		for _, e := range tree.CallArg {
			arg := Evaluator(e, env, genv)
			args = append(args, arg)
		}
		if genv[tree.Str].kind == builtin {
			builtinCall(tree.Str, args)
		} else {
			new_env := make(map[string]int)
			params := genv[tree.Str].args
			for i, e := range params {
				new_env[e] = args[i]
			}
			return Evaluator(genv[tree.Str].flow, new_env, genv)
		}
		return 0
	default:
		panic("Cannnot Evaluate")
	}
}

func SenteceEval(trees []*parser.Node) {
	env := make(map[string]int)
	genv := make(map[string]fn)
	//ビルトイン関数のセット
	genv["p"] = fn{kind: builtin}
	for _, e := range trees {
		Evaluator(e, env, genv)
	}
	fmt.Println("Evaluate ended with success")
}

func builtinCall(fnname string, args []int) {
	switch fnname {
	case "p":
		for _, e := range args {
			fmt.Println(e)
		}
	default:
		panic("No such function defined")
	}
}
