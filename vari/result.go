package rgx

type Result struct {
	Matches bool // Acceptance
	// Capturing groups: which substring in input matches a subexpression
	Groups map[string]string
}
