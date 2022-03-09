package cond

type TokenType int64

const (
	TokenTypeNum    TokenType = 1
	TokenTypeString TokenType = 2
	TokenTypeOP     TokenType = 3
)

type Token struct {
	Val  string
	OP   *ExprOperator
	Type TokenType
}

func tokenLegalAssert(c rune) {
	if c >= rune('0') && c <= rune('9') ||
		c >= rune('a') && c <= rune('z') ||
		c >= rune('A') && c <= rune('Z') ||
		c == rune('_') ||
		c == rune('$') ||
		c == rune('&') {
		return
	}
	panic("expr id is invalid, contains iilegal rune : " + string(c))
}

func genToken(src string) []Token {
	tokens := []Token{}
	id := ""

	addToken := func() {
		if id != "" {
			tokens = append(tokens, Token{Val: id})
			id = ""
		}
	}

	for _, v := range src {
		if v == rune(' ') {
			addToken()
			continue
		}
		if op := matchExprOperator(string(v)); op != nil {
			addToken()
			tokens = append(tokens, Token{Val: string(v), OP: op})
			continue
		}
		tokenLegalAssert(v)
		id = id + string(v)
	}
	return tokens
}
