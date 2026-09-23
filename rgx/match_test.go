package rgx

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
