package lexer

import (
	"strconv"
	"unicode"
)

// トークンの定義
type TokenKind int

const (
	TK_RESERVED TokenKind = iota //演算子
	TK_NUM                       //数字
	TK_IDENT                     //識別子
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
	//予約文字トークン
	iftoken := []rune("if")
	whiletoken := []rune("while")
	elsetoken := []rune("else")
	functoken := []rune("fn")
	returntoken := []rune("return")

	i := 0
	for i < len(rs) {
		//予約文字のパースのためにとりあえずスペースはトークンとして残して、[]Tokenのリターン前にスペースのトークンをはじく
		if isspace(rs[i]) {
			i++
			continue
		}
		if rs[i] == '+' || rs[i] == '-' || rs[i] == '*' || rs[i] == '/' || rs[i] == ';' || rs[i] == '{' || rs[i] == '}' || rs[i] == ',' {
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
		if rs[i] == '<' {
			if rs[i+1] == '=' {
				tok := Token{Kind: TK_RESERVED, Str: string(rs[i : i+2])}
				result = append(result, tok)
				i++
			} else {
				tok := Token{Kind: TK_RESERVED, Str: string(rs[i])}
				result = append(result, tok)
			}
			i++
			continue
		}
		if rs[i] == '>' {
			if rs[i+1] == '=' {
				tok := Token{Kind: TK_RESERVED, Str: string(rs[i : i+2])}
				result = append(result, tok)
				i++
			} else {
				tok := Token{Kind: TK_RESERVED, Str: string(rs[i])}
				result = append(result, tok)
			}
			i++
			continue
		}
		if rs[i] == '=' {
			if rs[i+1] == '=' {
				tok := Token{Kind: TK_RESERVED, Str: string(rs[i : i+2])}
				result = append(result, tok)
				i++
			} else {
				//variable assignment
				tok := Token{Kind: TK_RESERVED, Str: string(rs[i])}
				result = append(result, tok)
			}
			i++
			continue
		}
		if rs[i] == '!' {
			if rs[i+1] == '=' {
				tok := Token{Kind: TK_RESERVED, Str: string(rs[i : i+2])}
				result = append(result, tok)
				i++
			} else {
				panic("Not Expected")
			}
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
			//i++
			continue
		}
		if is_char(rs[i]) {
			var charrune []rune
			for is_alnum(rs[i]) && i < len(rs) {
				charrune = append(charrune, rs[i])
				if i == (len(rs) - 1) {
					break
				}
				i++
			}
			if isEqaulSlice(charrune, iftoken) {
				tok := Token{Kind: TK_RESERVED, Str: "if"}
				result = append(result, tok)
				continue
			}
			if isEqaulSlice(charrune, elsetoken) {
				tok := Token{Kind: TK_RESERVED, Str: "else"}
				result = append(result, tok)
				continue
			}
			if isEqaulSlice(charrune, whiletoken) {
				tok := Token{Kind: TK_RESERVED, Str: "while"}
				result = append(result, tok)
				continue
			}
			if isEqaulSlice(charrune, functoken) {
				tok := Token{Kind: TK_RESERVED, Str: "fn"}
				result = append(result, tok)
				continue
			}
			if isEqaulSlice(charrune, returntoken) {
				tok := Token{Kind: TK_RESERVED, Str: "return"}
				result = append(result, tok)
				continue
			}
			tok := Token{Kind: TK_IDENT, Str: string(charrune)}
			result = append(result, tok)
			//i++
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

func isEqaulSlice(x []rune, y []rune) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
	return true
}

func is_char(x rune) bool {
	return ('a' <= x && x <= 'z') || ('A' <= x && x <= 'Z') || (x == '_')
}

func is_alnum(x rune) bool {
	return ('a' <= x && x <= 'z') || ('A' <= x && x <= 'Z') || (x == '_') || ('0' <= x && x <= '9')
}
