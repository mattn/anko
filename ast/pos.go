package ast

type Position struct {
	Line int
	Column int
}

type Pos interface {
	GetPos() Position
	SetPos(Position)
}

type PosImpl struct {
	pos Position
}

func (x *PosImpl) GetPos() Position {
	return x.pos
}

func (x *PosImpl) SetPos(pos Position) {
	x.pos = pos
}
