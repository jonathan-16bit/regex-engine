package vari

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

func alternate(leftStart, leftEnd, rightStart, rightEnd *State) (*State, *State) {
	start := &State{
		transitions: make(map[uint8][]*State),
	}

	end := &State{
		transitions: make(map[uint8][]*State),
	}

	start.transitions[epsilon] = []*State{leftStart, rightStart}
	leftEnd.transitions[epsilon] = append(leftEnd.transitions[epsilon], end)
	rightEnd.transitions[epsilon] = append(rightEnd.transitions[epsilon], end)

	return start, end
}

func kleeneStar(innerStart, innerEnd *State) (*State, *State) {
	newStart := &State{
		transitions: make(map[uint8][]*State),
	}

	newEnd := &State{
		transitions: make(map[uint8][]*State),
	}

	newStart.transitions[epsilon] = append(newStart.transitions[epsilon], innerStart, newEnd)
	innerEnd.transitions[epsilon] = append(innerEnd.transitions[epsilon], innerStart, newEnd)

	return newStart, newEnd
}
