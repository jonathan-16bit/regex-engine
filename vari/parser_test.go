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

func TestParseRepetitionStar(t *testing.T) {
	p := parser{
		tokens: tokenize("a*"),
	}

	expr, err := p.parseRepetition()
	if err != nil {
		t.Fatalf("expected no error, got %+v", err)
	}

	if expr == nil {
		t.Fatalf("expected valid expression, got %+v", expr)
	}

	if expr.kind != expressionStar {
		t.Fatalf("expected kind expressionStar, got %v", expr.kind)
	}

	if expr.child == nil {
		t.Fatal("expected valid child expression, got nil")
	}

	if expr.child.kind != expressionLiteral {
		t.Fatalf("expected expressionLiteral child kind, got %v", expr.child.kind)
	}

	if expr.child.literal != 'a' {
		t.Fatalf("expected child 'a', got %v", expr.child.literal)
	}

	curr := p.current()
	if curr.kind != tokenEOF {
		t.Fatalf("expected tokenEOF, got %v", curr)
	}
}
