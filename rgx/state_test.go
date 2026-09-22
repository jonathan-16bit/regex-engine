package rgx

import "testing"

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
