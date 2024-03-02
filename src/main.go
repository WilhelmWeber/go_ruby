package main

import (
	"github.com/WilhelmWeber/go_ruby/src/evaluator"
	"github.com/WilhelmWeber/go_ruby/src/filereader"
	"github.com/WilhelmWeber/go_ruby/src/lexer"
	"github.com/WilhelmWeber/go_ruby/src/parser"
)

func main() {
	result := filereader.Reader("test.text")
	tokens := lexer.Tokenize(result)
	var env_parser []string
	p := parser.Parser{Token: tokens, Index: 0, Env: env_parser}
	tree := p.Parse()
	evaluator.SenteceEval(tree)
}
