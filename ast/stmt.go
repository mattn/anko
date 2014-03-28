package ast

type Stmt interface {
	stmt()
}

type StmtImpl struct {
}

func (x *StmtImpl) stmt() {}

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
	Name string
	Stmts []Stmt
	Args []string
}

type ReturnStmt struct {
	StmtImpl
	Expr Expr
}
