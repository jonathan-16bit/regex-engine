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
