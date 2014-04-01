package ast

// Position provides interface to store code locations.
type Position struct {
	Line int
	Column int
}

// Pos interface provies two functions to get/set the position for expression or statement.
type Pos interface {
	GetPos() Position
	SetPos(Position)
}

// PosImpl provies commonly implementations for Pos.
type PosImpl struct {
	pos Position
}

// GetPos return the position of the expression or statement.
func (x *PosImpl) GetPos() Position {
	return x.pos
}

// SetPos is a function to specify position of the expression or statement.
func (x *PosImpl) SetPos(pos Position) {
	x.pos = pos
}
