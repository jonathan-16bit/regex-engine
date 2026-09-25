package rgx

import "testing"

func TestParsePrimaryLiteral(t *testing.T) {
	// First test
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
