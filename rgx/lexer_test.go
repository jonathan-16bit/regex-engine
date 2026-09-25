package rgx

import "testing"

func TestTokenize(t *testing.T) {
	pattern := "a(b|c)*d"

	want := []token{
		{
			kind: tokenLiteral,
			ch:   'a',
			pos:  0,
		},
		{
			kind: tokenLeftParen,
			ch:   '(',
			pos:  1,
		},
		{
			kind: tokenLiteral,
			ch:   'b',
			pos:  2,
		},
		{
			kind: tokenPipe,
			ch:   '|',
			pos:  3,
		},
		{
			kind: tokenLiteral,
			ch:   'c',
			pos:  4,
		},
		{
			kind: tokenRightParen,
			ch:   ')',
			pos:  5,
		},
		{
			kind: tokenStar,
			ch:   '*',
			pos:  6,
		},
		{
			kind: tokenLiteral,
			ch:   'd',
			pos:  7,
		},
		{
			kind: tokenEOF,
			ch:   0,
			pos:  8,
		},
	}

	got := tokenize(pattern)

	if len(got) != len(want) {
		t.Fatalf("expected length %v, got %v", len(want), len(got))
	}

	for i := range want {
		if want[i] != got[i] {
			t.Errorf("token %d: wanted %+v, got %+v", i, want[i], got[i])
		}
	}
}

func TestEmptyTokenize(t *testing.T) {
	want := token{
		kind: tokenEOF,
		ch:   0,
		pos:  0,
	}

	got := tokenize("")
	if len(got) != 1 {
		t.Fatalf("wanted 1 token, got %v", len(got))
	}

	if want != got[0] {
		t.Fatalf("wanted %+v, got %+v", want, got[0])
	}
}
