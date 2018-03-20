package ast

type Token struct {
	PosImpl // PosImpl provides get/set the position function.
	Tok     int
	Lit     string
}
