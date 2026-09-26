package vari

import "testing"

func TestParsePrimaryLiteral(t *testing.T) {
	p := parser{
		tokens: tokenize("a"),
	}

	expr, err := p.parsePrimary()

	if err != nil {
		t.Fatalf("parsePrimary returned an unexpected error: %v", err)
	}

	if p.current().kind != tokenEOF {
		t.Fatalf("expected tokenEOF after parsing literal, got %v", p.current())
	}

	if expr == nil {
		t.Fatal("valid expression must be returned on parsePrimary() for 'a'")
	}

	if expr.kind != expressionLiteral {
		t.Fatalf("expected kind expressionLiteral, got %v", expr.kind)
	}

	if expr.literal != 'a' {
		t.Fatalf("expected literal 'a', got %v", expr.literal)
	}
}

func TestParsePrimaryRejectStar(t *testing.T) {
	p := parser{
		tokens: tokenize("*"),
	}

	initialPos := p.pos
	expr, err := p.parsePrimary()

	if err == nil {
		t.Fatal("expected parsePrimary to return error for tokenStar")
	}

	if expr != nil {
		t.Fatalf("expected parsePrimary to not return expression, got %+v", expr)
	}

	if p.pos != initialPos {
		t.Fatalf("expected parser position %v, got %v", initialPos, p.pos)
	}

	if p.current().kind != tokenStar {
		t.Fatalf("expected current token to remain tokenStar, got %v", p.current())
	}
}
