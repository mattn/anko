package ast

type Token struct {
	Tok int
	Lit string
	pos Position
}

func (t *Token) GetPos() Position {
	return t.pos
}

func (t *Token) SetPos(pos Position) {
	t.pos = pos
}

