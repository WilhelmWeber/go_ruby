package lexer

import (
	"strconv"
	"unicode"
)

// トークンの定義
type TokenKind int

const (
	TK_RESERVED TokenKind = iota //演算子
	TK_NUM
	TK_EOF
)

type Token struct {
	Kind TokenKind
	Val  int
	Str  string
}

func Tokenize(s string) []Token {

	var result []Token
	rs := []rune(s)

	i := 0
	for i < len(rs) {
		if isspace(rs[i]) {
			i++
			continue
		}
		if rs[i] == '+' || rs[i] == '-' || rs[i] == '*' || rs[i] == '/' {
			tok := Token{Kind: TK_RESERVED, Str: string(rs[i])}
			result = append(result, tok)
			i++
			continue
		}
		if rs[i] == '(' {
			tok := Token{Kind: TK_RESERVED, Str: string(rs[i])}
			result = append(result, tok)
			i++
			continue
		}
		if rs[i] == ')' {
			tok := Token{Kind: TK_RESERVED, Str: string(rs[i])}
			result = append(result, tok)
			i++
			continue
		}
		if unicode.IsDigit(rs[i]) {
			var numrune []rune
			//numrune = append(numrune, rs[i])

			for unicode.IsDigit(rs[i]) && i < len(rs) {
				numrune = append(numrune, rs[i])
				if i == (len(rs) - 1) {
					break
				}
				i++
			}

			num, _ := strconv.Atoi(string(numrune))
			tok := Token{Kind: TK_NUM, Val: num}
			result = append(result, tok)
			//もうすでに読み進めているためインクリメントはいらない
			i++
			continue
		}
		panic("Cannnot Tokenize")
	}

	result = append(result, Token{Kind: TK_EOF})
	return result
}

//utility

func isspace(a rune) bool {
	if a == ' ' {
		return true
	} else {
		return false
	}
}
