package rgx

// Special transition key
const epsilon uint8 = 0

// No terminal states set yet
func literalFragment(ch uint8) (*State, *State) {
	start := &State{
		transitions: make(map[uint8][]*State),
	}

	end := &State{
		transitions: make(map[uint8][]*State),
	}

	start.transitions[ch] = []*State{end}

	return start, end
}

func concatenate(leftStart, leftEnd, rightStart, rightEnd *State) (*State, *State) {
	leftEnd.transitions[epsilon] = append(leftEnd.transitions[epsilon], rightStart)
	return leftStart, rightEnd
}
