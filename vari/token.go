package vari

type tokenKind uint8

const (
	tokenInvalid tokenKind = iota
	tokenLiteral
	tokenStar
	tokenPipe
	tokenLeftParen
	tokenRightParen
	tokenEOF
)

type token struct {
	kind tokenKind
	ch   uint8
	pos  int
}
