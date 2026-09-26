package vari

type expressionKind uint8
const (
	expressionInvalid expressionKind = iota
	expressionLiteral
	expressionConcatenation
	expressionAlternation
	expressionStar
)

type expression struct {
	kind expressionKind
	literal uint8
	child *expression
	left *expression
	right *expression
}
