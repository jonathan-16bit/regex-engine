package vari

import "testing"

func TestAcceptsConcatenatedLiterals(t *testing.T) {
	aStart, aEnd := literalFragment('a')
	bStart, bEnd := literalFragment('b')

	start, end := concatenate(aStart, aEnd, bStart, bEnd)
	end.terminal = true

	tests := []struct {
		input string
		want  bool
	}{
		{input: "ab", want: true},
		{input: "a", want: false},
		{input: "b", want: false},
		{input: "ac", want: false},
		{input: "abc", want: false},
		{input: "", want: false},
	}

	for _, test := range tests {
		got := accepts(start, test.input)

		if got != test.want {
			// Errorf instead of Fatal, so that one incorrect case doesnt prevent the others from running
			t.Errorf("accepts(%q) = %v; want %v", test.input, got, test.want)
		}
	}
}

func TestAcceptsAlternation(t *testing.T) {
	aStart, aEnd := literalFragment('a')
	bStart, bEnd := literalFragment('b')

	start, end := alternate(aStart, aEnd, bStart, bEnd)
	end.terminal = true

	tests := []struct {
		input string
		want  bool
	}{
		{input: "a", want: true},
		{input: "b", want: true},
		{input: "", want: false},
		{input: "ab", want: false},
		{input: "c", want: false},
	}

	for _, test := range tests {
		got := accepts(start, test.input)

		if got != test.want {
			t.Errorf("accepts(%q) = %v; want %v", test.input, got, test.want)
		}
	}
}

func TestAcceptsEpsilonCycle(t *testing.T) {
	// First test
	q0 := &State{
		transitions: make(map[uint8][]*State),
	}

	q1 := &State{
		transitions: make(map[uint8][]*State),
	}

	q0.transitions[epsilon] = []*State{q1}
	q1.transitions[epsilon] = []*State{q0}

	// Test
	if accepts(q0, "") {
		t.Fatal("non-accepting epsilon cycle should reject")
	}

	terminal := &State{
		terminal:    true,
		transitions: make(map[uint8][]*State),
	}

	q1.transitions[epsilon] = append(q1.transitions[epsilon], terminal)

	if !accepts(q0, "") {
		t.Fatal("accepting epsilon cycle should accept")
	}
}

func TestAcceptsKleeneStar(t *testing.T) {
	aStart, aEnd := literalFragment('a')
	start, end := kleeneStar(aStart, aEnd)
	end.terminal = true

	tests := []struct {
		input string
		want  bool
	}{
		{input: "", want: true},
		{input: "a", want: true},
		{input: "aa", want: true},
		{input: "b", want: false},
		{input: "ab", want: false},
		{input: "ba", want: false},
		{input: "aaaab", want: false},
	}

	for _, test := range tests {
		got := accepts(start, test.input)

		if got != test.want {
			t.Errorf("accepts(%q) = %v; want %v", test.input, got, test.want)
		}
	}
}
