package main

import (
	"fmt"

	"github.com/WilhelmWeber/go_ruby/src/evaluator"
	"github.com/WilhelmWeber/go_ruby/src/lexer"
	"github.com/WilhelmWeber/go_ruby/src/parser"
)

func main() {
	Token := lexer.Tokenize("( 52 + 4 ) * 3")
	parser := parser.Parser{Token: Token, Index: 0}
	tree := parser.Expr()

	result := evaluator.Evaluator(tree)
	fmt.Println(result)
}
