package ast

type Token struct {
	Tok int
	Lit string
	Pos Position
}

func (t *Token) GetPos() Position {
	return t.Pos
}

func (t *Token) SetPos(pos Position) {
	t.Pos = pos
}
