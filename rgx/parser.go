package rgx

type parser struct {
	tokens []token
	pos int
}

// Method receiver
func (p *parser) current() token {
	return p.tokens[p.pos]
}

func (p *parser) advance() token {
	curr := p.current()
	if curr.kind != tokenEOF {
		p.pos += 1
	}
	return curr
}
