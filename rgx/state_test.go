package rgx

import "testing"

func TestLiteralStateGraph(t *testing.T) {
	end := &State{
		terminal: true,
		transitions: make(map[uint8][]*State),
	}

	start := &State{
		transitions: make(map[uint8][]*State),
	}

	start.transitions['a'] = []*State{end}
	destinations := start.transitions['a']

	if len(destinations) != 1 {
		t.Fatalf("Expected 1 destination, got %d\n", len(destinations))
	}

	if destinations[0] != end {
		t.Fatal("Transition does not lead to expected state")
	}

	if !destinations[0].terminal {
		t.Fatal("Destination should be terminal")
	}
}
