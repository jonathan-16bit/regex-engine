package vari

type State struct {
	terminal bool
	// map input byte to destination states (as in, pointers to them)
	transitions map[uint8][]*State
}
