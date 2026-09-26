package rgx

import "testing"
import "slices"

func TestLiteralFragment(t *testing.T) {
	start, end := literalFragment('a')
	destinations := start.transitions['a']

	if len(destinations) != 1 {
		t.Fatalf("expected 1 destination, got %d", len(destinations))
	}

	if destinations[0] != end {
		t.Fatal("transition does not lead to expected state")
	}

	if end.terminal {
		t.Fatal("fragment end should not be terminal before completion")
	}

	if start.transitions == nil || end.transitions == nil {
		t.Fatal("fragment states should have initialized transition maps")
	}
}

func TestConcatenateFragments(t *testing.T) {
	aStart, aEnd := literalFragment('a')
	bStart, bEnd := literalFragment('b')
	start, end := concatenate(aStart, aEnd, bStart, bEnd)

	if start != aStart {
		t.Fatal("incorrect start state returned")
	}

	if end != bEnd {
		t.Fatal("incorrect end state returned")
	}

	aDestinations := aStart.transitions['a']
	if len(aDestinations) != 1 {
		t.Fatalf("expected 1 'a' destination, got %d", len(aDestinations))
	}
	if aDestinations[0] != aEnd {
		t.Fatal("transition on 'a' from aStart fails")
	}

	epsilonDestinations := aEnd.transitions[epsilon]
	if len(epsilonDestinations) != 1 {
		t.Fatalf("expected 1 'epsilon' destination, got %d", len(epsilonDestinations))
	}
	if epsilonDestinations[0] != bStart {
		t.Fatal("transition on 'epsilon' from aEnd fails")
	}

	bDestinations := bStart.transitions['b']
	if len(bDestinations) != 1 {
		t.Fatalf("expected 1 'b' destination, got %d", len(bDestinations))
	}
	if bDestinations[0] != bEnd {
		t.Fatal("transition on 'b' from bStart fails")
	}
}

func TestAlternateFragments(t *testing.T) {
	aStart, aEnd := literalFragment('a')
	bStart, bEnd := literalFragment('b')
	start, end := alternate(aStart, aEnd, bStart, bEnd)

	foundLeft := false
	foundRight := false

	if len(start.transitions[epsilon]) != 2 {
		t.Fatalf("2 epsilon transitions expected, %v found", len(start.transitions[epsilon]))
	}

	for _, destination := range start.transitions[epsilon] {
		if destination == aStart {
			foundLeft = true
		}
		if destination == bStart {
			foundRight = true
		}
	}

	if foundLeft == false {
		t.Fatalf("start of a is not connected to alternate")
	}

	if foundRight == false {
		t.Fatalf("start of b is not connected to alternate")
	}

	if !slices.Contains(aEnd.transitions[epsilon], end) {
		t.Fatalf("no transition from aEnd to end")
	}

	if !slices.Contains(bEnd.transitions[epsilon], end) {
		t.Fatalf("no transition from bEnd to end")
	}

	if end.transitions == nil {
		t.Fatal("alternation end should have an initialized transition map")
	}

	if end.terminal {
		t.Fatal("alternation end should not be terminal before completion")
	}
}

func TestKleeneStarFragment(t *testing.T) {
	aStart, aEnd := literalFragment('a')
	newStart, newEnd := kleeneStar(aStart, aEnd);

	// Epsilon transitions from newStart
	newStartEpsilonTransitions := newStart.transitions[epsilon]
	if len(newStartEpsilonTransitions) != 2 {
		t.Fatalf("expected 2 epsilon transitions from newStart, got %v", len(newStartEpsilonTransitions))
	}

	if !slices.Contains(newStartEpsilonTransitions, aStart) {
		t.Fatal("aStart is unreachable from newStart")
	}

	if !slices.Contains(newStartEpsilonTransitions, newEnd) {
		t.Fatal("newEnd is unreachable from newStart")
	}

	// Epsilon transitions from aEnd
	aEndEpsilonTransitions := aEnd.transitions[epsilon]
	if len(aEndEpsilonTransitions) != 2 {
		t.Fatalf("expected 2 epsilon transitions from aEnd, got %v", len(aEndEpsilonTransitions))
	}

	if !slices.Contains(aEndEpsilonTransitions, aStart) {
		t.Fatal("aStart unreachable from aEnd")
	}

	if !slices.Contains(aEndEpsilonTransitions, newEnd) {
		t.Fatal("newEnd unreachable from aEnd")
	}

	// No epsilon transition from aStart to aEnd
	if len(aStart.transitions[epsilon]) != 0 {
		t.Fatal("aStart has epsilon transition to aEnd")
	}

	// newEnd has an initialized map
	if newEnd.transitions == nil {
		t.Fatal("newEnd has uninitialized map")
	}

	// newEnd must not be terminal
	if newEnd.terminal {
		t.Fatal("newEnd is terminal")
	}
}
