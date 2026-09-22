package rgx

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
