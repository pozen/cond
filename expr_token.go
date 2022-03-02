package cond

type Token struct {
	Val string
	OP  *ExprOperator
}

func idLegalAssert(c rune) {
	if c >= rune('0') && c <= rune('9') ||
		c >= rune('a') && c <= rune('z') ||
		c >= rune('A') && c <= rune('Z') ||
		c == rune('_') {
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
		idLegalAssert(v)
		id = id + string(v)
	}
	return tokens
}
