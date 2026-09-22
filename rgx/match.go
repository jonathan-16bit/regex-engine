package rgx

func accepts(start *State, input string) bool {
	return acceptsFrom(start, input, 0)
}

func acceptsFrom(curr *State, input string, pos int) bool {
	// If input exhausted and accepting state reached
	if pos == len(input) && curr.terminal {
		return true
	}

	// Epsilon transitions while keeping pos
	epsilonDestinations := curr.transitions[epsilon]
	for _, next := range epsilonDestinations {
		if acceptsFrom(next, input, pos) {
			return true
		}
	}

	// Character transitions
	if pos < len(input) {
		inputDestinations := curr.transitions[input[pos]]
		for _, next := range inputDestinations {
			if acceptsFrom(next, input, pos+1) {
				return true
			}
		}
	}

	return false
}
