package parser

import (
	"github.com/WilhelmWeber/go_ruby/src/lexer"
)

// 疑似クラス化
type Parser struct {
	Token []lexer.Token //トークナイズした文字列
	Index int           //注目しているtokenのインデックス 初期値は0
}

type NodeKind int

const (
	ND_ADD NodeKind = iota
	ND_SUB
	ND_MUL
	ND_DIV
	ND_NUM
)

type Node struct {
	Kind NodeKind
	Lhs  *Node
	Rhs  *Node
	Val  int
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

// 期待している文字列であればトークンを読み進め、それ以外ではエラーを返す
func (parser *Parser) expect(op string) {
	if parser.Token[parser.Index].Kind != lexer.TK_RESERVED || parser.Token[parser.Index].Str != op {
		panic("Not Expected Token")
	}
	parser.Index++
}

func (parser *Parser) Expr() *Node {
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
	node := parser.primary()
	for {
		if parser.consume("*") {
			rhs := parser.primary()
			node = &Node{Kind: ND_MUL, Lhs: node, Rhs: rhs}
		} else if parser.consume("/") {
			rhs := parser.primary()
			node = &Node{Kind: ND_DIV, Lhs: node, Rhs: rhs}
		} else {
			return node
		}
	}
}

func (parser *Parser) primary() *Node {
	if parser.consume("(") {
		node := parser.Expr()
		parser.expect(")")
		return node
	}
	return &Node{Kind: ND_NUM, Val: parser.expect_number()}
}
