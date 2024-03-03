# Go_Ruby (ruby interpreter made with golang)  
[Rubyで学ぶRuby](https://ascii.jp/serialarticles/1230449/)に触発されて作ったGolang製インタプリンタ  
上の記事ではmini_rubyと呼ばれる字句解析・構文解析機がライブラリとして用意されており、読者は評価機だけを実装するようになっていますが、本レポジトリは字句解析機（lexer.go）と構文解析機（parser.go）も[この記事](https://www.sigbus.info/compilerbook)を参考に自前で実装しています。  
```go
package main

import (
	"github.com/WilhelmWeber/go_ruby/src/evaluator"
	"github.com/WilhelmWeber/go_ruby/src/filereader"
	"github.com/WilhelmWeber/go_ruby/src/lexer"
	"github.com/WilhelmWeber/go_ruby/src/parser"
)

func main() {

	result := filereader.Reader("<コードを書いたファイル>")
	tokens := lexer.Tokenize(result)
	var env_parser []string
	p := parser.Parser{Token: tokens, Index: 0, Env: env_parser}
	tree := p.Parse()
	evaluator.SenteceEval(tree)

}
```
