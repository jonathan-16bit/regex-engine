package rgx

func tokenize(pattern string) []token {
	tokens := make([]token, 0, len(pattern)+1)
	for pos := 0; pos < len(pattern); pos++ {
		ch := pattern[pos]

		// Assume literal token, take cases
		kind := tokenLiteral
		switch ch {
		case '*':
			kind = tokenStar
		case '|':
			kind = tokenPipe
		case '(':
			kind = tokenLeftParen
		case ')':
			kind = tokenRightParen
		}

		// Append to list of tokens
		tokens = append(tokens, token{
			kind: kind,
			ch:   ch,
			pos:  pos,
		})
	}

	tokens = append(tokens, token{
		kind: tokenEOF,
		ch:   0,
		pos:  len(pattern),
	})

	return tokens
}
