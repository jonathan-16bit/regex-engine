package rgx

type matchConfiguration struct {
	state *State
	pos   int
}

func accepts(start *State, input string) bool {
	visited := make(map[matchConfiguration]bool)
	return acceptsFrom(start, input, 0, visited)
}

func acceptsFrom(curr *State, input string, pos int, visited map[matchConfiguration]bool) bool {
	currConfig := matchConfiguration{
		state: curr,
		pos:   pos,
	}

	if visited[currConfig] {
		return false
	}

	visited[currConfig] = true

	// If input exhausted and accepting state reached
	if pos == len(input) && curr.terminal {
		return true
	}

	// Epsilon transitions while keeping pos
	epsilonDestinations := curr.transitions[epsilon]
	for _, next := range epsilonDestinations {
		if acceptsFrom(next, input, pos, visited) {
			return true
		}
	}

	// Character transitions
	if pos < len(input) {
		inputDestinations := curr.transitions[input[pos]]
		for _, next := range inputDestinations {
			if acceptsFrom(next, input, pos+1, visited) {
				return true
			}
		}
	}

	return false
}
