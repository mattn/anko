package ast

type Stmt interface {
	Pos
	stmt()
}

type StmtImpl struct {
	PosImpl
}

func (x *StmtImpl) stmt() {}

type ExprStmt struct {
	StmtImpl
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
	If   Expr
	Then []Stmt
	Else []Stmt
}

type TryStmt struct {
	StmtImpl
	Try     []Stmt
	Var     string
	Catch   []Stmt
	Finally []Stmt
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

type ThrowStmt struct {
	StmtImpl
	Expr Expr
}

type ModuleStmt struct {
	StmtImpl
	Name  string
	Stmts []Stmt
}

type LetStmt struct {
	StmtImpl
	Names []string
	Exprs []Expr
}
