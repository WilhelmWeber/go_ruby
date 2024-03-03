package parser

import (
	"github.com/WilhelmWeber/go_ruby/src/lexer"
)

// 疑似クラス化
type Parser struct {
	Token []lexer.Token //トークナイズした文字列
	Index int           //注目しているtokenのインデックス 初期値は0
	Env   []string      //使用されている変数の管理
}

type NodeKind int

const (
	ND_ADD NodeKind = iota
	ND_SUB
	ND_MUL
	ND_DIV
	ND_NUM
	ND_EQU // ==
	ND_UEQ // !=
	ND_LES // <
	ND_ELS // <=
	ND_MOR // >
	ND_EMR // >=
	ND_ASN // 変数代入
	ND_IDT // 識別子
	ND_REF // 変数参照
	ND_IF  // IF文
	ND_WHL // While文
	ND_BLC // ブロック文
	ND_RTN // return
	ND_DEF // 関数定義
	ND_FNC // 関数呼び出し 現時点ではprintlnのみ対応
)

type Node struct {
	Kind    NodeKind
	Lhs     *Node
	Rhs     *Node
	Val     int
	Str     string   //変数の管理用
	CallArg []*Node  //関数呼び出しの引数
	If      *Node    //条件式
	Then    *Node    //条件式が1だったときに実行する処理
	Else    *Node    //条件式が0だったときに実行する処理
	Stmts   []*Node  //ブロック内の複文
	Args    []string //関数定義の時の引数
	Flow    *Node    //関数であったときの処理
}

// methods
// 次のトークンが引数の記号のときはインデックスを一つ進めてtrueを返す
func (parser *Parser) consume(op string) bool {
	if parser.Token[parser.Index].Kind != lexer.TK_RESERVED || parser.Token[parser.Index].Str != op {
		return false
	}
	parser.Index++
	return true
}

// 次のトークンが数値の場合はインデックスを一つ進めて数値を返す、それ以外はerror
func (parser *Parser) expect_number() int {
	if parser.Token[parser.Index].Kind != lexer.TK_NUM {
		panic("Not Number")
	}
	val := parser.Token[parser.Index].Val
	parser.Index++
	return val
}

// 次のトークンが識別子である場合はインデックスを一つ進めて文字とtrueを返す
func (parser *Parser) consume_ident() (string, bool) {
	if parser.Token[parser.Index].Kind != lexer.TK_IDENT {
		return "", false
	}
	str := parser.Token[parser.Index].Str
	parser.Index++
	return str, true
}

// 期待している文字列であればトークンを読み進め、それ以外ではエラーを返す
func (parser *Parser) expect(op string) {
	if parser.Token[parser.Index].Kind != lexer.TK_RESERVED || parser.Token[parser.Index].Str != op {
		panic("Not Expected Token")
	}
	parser.Index++
}

func (parser *Parser) at_eof() bool {
	return parser.Token[parser.Index].Kind == lexer.TK_EOF
}

// Parser
func (parser *Parser) Parse() []*Node {
	var nodes []*Node

	for !parser.at_eof() {
		node := parser.stmt()
		nodes = append(nodes, node)
	}
	return nodes
}

func (parser *Parser) stmt() *Node {
	if parser.consume("if") {
		parser.expect("(")
		If := parser.expr()
		parser.expect(")")
		Then := parser.stmt()
		if parser.consume("else") {
			Else := parser.stmt()
			return &Node{Kind: ND_IF, If: If, Then: Then, Else: Else}
		}
		return &Node{Kind: ND_IF, If: If, Then: Then}
	}
	if parser.consume("while") {
		parser.expect("(")
		While := parser.expr()
		parser.expect(")")
		Do := parser.stmt()
		return &Node{Kind: ND_WHL, If: While, Then: Do}
	}
	if parser.consume("{") {
		var stmts []*Node
		for !parser.consume("}") {
			stmt := parser.stmt()
			stmts = append(stmts, stmt)
		}
		return &Node{Kind: ND_BLC, Stmts: stmts}
	}
	if parser.consume("fn") {
		var args []string
		funcname, isIdent := parser.consume_ident()
		if !isIdent {
			panic("Not named Function")
		}
		parser.expect("(")
		for {
			arg, _ := parser.consume_ident()
			args = append(args, arg)
			//関数ブロックの識別子を変数参照にしたいためにparser.Envにpush
			parser.Env = append(parser.Env, arg)
			if parser.consume(")") {
				break
			}
			parser.expect(",")
		}
		flow := parser.stmt()
		//parser.Envにプッシュした引数変数を削除してparser.Envをセットしなおす
		parser.Env = envDelate(parser.Env, args)
		return &Node{Kind: ND_DEF, Args: args, Flow: flow, Str: funcname}
	}
	if parser.consume("return") {
		rhs := parser.expr()
		node := &Node{Kind: ND_RTN, Rhs: rhs}
		parser.expect(";")
		return node
	}
	node := parser.expr()
	parser.expect(";")
	return node
}

func (parser *Parser) expr() *Node {
	node := parser.equality()
	if parser.consume("=") {
		rhs := parser.expr()
		node = &Node{Kind: ND_ASN, Lhs: node, Rhs: rhs}
	}
	return node
}

func (parser *Parser) equality() *Node {
	node := parser.relat()
	for {
		if parser.consume("==") {
			rhs := parser.relat()
			node = &Node{Kind: ND_EQU, Lhs: node, Rhs: rhs}
		} else if parser.consume("!=") {
			rhs := parser.relat()
			node = &Node{Kind: ND_UEQ, Lhs: node, Rhs: rhs}
		} else {
			return node
		}
	}
}

func (parser *Parser) relat() *Node {
	node := parser.add()
	for {
		if parser.consume("<=") {
			rhs := parser.add()
			node = &Node{Kind: ND_ELS, Lhs: node, Rhs: rhs}
		} else if parser.consume("<") {
			rhs := parser.add()
			node = &Node{Kind: ND_LES, Lhs: node, Rhs: rhs}
		} else if parser.consume(">=") {
			rhs := parser.add()
			node = &Node{Kind: ND_EMR, Lhs: node, Rhs: rhs}
		} else if parser.consume(">") {
			rhs := parser.add()
			node = &Node{Kind: ND_MOR, Lhs: node, Rhs: rhs}
		} else {
			return node
		}
	}
}

func (parser *Parser) add() *Node {
	node := parser.mul()
	for {
		if parser.consume("+") {
			rhs := parser.mul()
			node = &Node{Kind: ND_ADD, Lhs: node, Rhs: rhs}
		} else if parser.consume("-") {
			rhs := parser.mul()
			node = &Node{Kind: ND_SUB, Lhs: node, Rhs: rhs}
		} else {
			return node
		}
	}
}

func (parser *Parser) mul() *Node {
	node := parser.unary()
	for {
		if parser.consume("*") {
			rhs := parser.unary()
			node = &Node{Kind: ND_MUL, Lhs: node, Rhs: rhs}
		} else if parser.consume("/") {
			rhs := parser.unary()
			node = &Node{Kind: ND_DIV, Lhs: node, Rhs: rhs}
		} else {
			return node
		}
	}
}

func (parser *Parser) unary() *Node {
	if parser.consume("+") {
		return parser.primary()
	}
	if parser.consume("-") {
		rhs := parser.primary()
		zero_node := &Node{Kind: ND_NUM, Val: 0}
		return &Node{Kind: ND_SUB, Lhs: zero_node, Rhs: rhs}
	}
	return parser.primary()
}

func (parser *Parser) primary() *Node {
	if parser.consume("(") {
		node := parser.expr()
		parser.expect(")")
		return node
	}
	if str, isIdent := parser.consume_ident(); isIdent {
		//括弧が次に来れば関数呼び出しのはず
		if parser.consume("(") {
			var args []*Node
			for {
				arg := parser.expr()
				args = append(args, arg)
				if parser.consume(")") {
					break
				}
				parser.expect(",")
			}
			return &Node{Kind: ND_FNC, Str: str, CallArg: args}
		} else {
			if isIn(parser.Env, str) {
				//既出の変数が含まれていれば参照として扱う
				return &Node{Kind: ND_REF, Str: str}
			} else {
				//初出の変数はEnvにpushする
				parser.Env = append(parser.Env, str)
				return &Node{Kind: ND_IDT, Str: str}
			}
		}
	}
	return &Node{Kind: ND_NUM, Val: parser.expect_number()}
}

// utility

func isIn(slice []string, str string) bool {
	for _, e := range slice {
		if e == str {
			return true
		}
	}
	return false
}

func envDelate(envs []string, args []string) []string {
	var slice []string
	for _, env := range envs {
		if isIn(args, env) {
			continue
		} else {
			slice = append(slice, env)
		}
	}
	return slice
}
