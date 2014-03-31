package ast

type Position struct {
	Line   int
	Column int
}

type Stmt interface {
	stmt()
	SetPos(Position)
	GetPos() Position
}

type StmtImpl struct {
	pos Position
}

func (x *StmtImpl) stmt() {}

func (x *StmtImpl) SetPos(pos Position) {
	x.pos = pos
}

func (x *StmtImpl) GetPos() Position {
	return x.pos
}

type ExprStmt struct {
	StmtImpl
	Expr Expr
}

type VarStmt struct {
	StmtImpl
	Name string
	Expr Expr
}

type FuncStmt struct {
	StmtImpl
	Name   string
	Stmts  []Stmt
	Args   []string
	VarArg bool
}

type IfStmt struct {
	StmtImpl
	Expr      Expr
	ThenStmts []Stmt
	ElseStmts []Stmt
}

type ForStmt struct {
	StmtImpl
	Var   string
	Value Expr
	Stmts []Stmt
}

type ReturnStmt struct {
	StmtImpl
	Expr Expr
}

type ModuleStmt struct {
	StmtImpl
	Name  string
	Stmts []Stmt
}
