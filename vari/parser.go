package vari

import "fmt"

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

func (p *parser) parsePrimary() (*expression, error) {
	current := p.current()

	switch current.kind {
	case tokenLiteral:
		p.advance()
		return &expression {
			kind: expressionLiteral,
			literal: current.ch,
		}, nil

	default:
		return nil, fmt.Errorf (
			"expected expression at position %d", current.pos,
		)
	}
}
