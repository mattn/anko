package ast

// Stmt provides all of interfaces for statement.
type Stmt interface {
	Pos
	stmt()
}

// StmtImpl provide commonly implementations for Stmt..
type StmtImpl struct {
	PosImpl // StmtImpl provide Pos() function.
}

// stmt provide restraint interface.
func (x *StmtImpl) stmt() {}

// ExprStmt provide expression statement.
type ExprStmt struct {
	StmtImpl
	Expr Expr
}

// IfStmt provide "if/else" statement.
type IfStmt struct {
	StmtImpl
	If     Expr
	Then   []Stmt
	ElseIf []Stmt // This is array of IfStmt
	Else   []Stmt
}

// TryStmt provide "try/catch/finally" statement.
type TryStmt struct {
	StmtImpl
	Try     []Stmt
	Var     string
	Catch   []Stmt
	Finally []Stmt
}

// ForStmt provide "for in" expression statement.
type ForStmt struct {
	StmtImpl
	Var   string
	Value Expr
	Stmts []Stmt
}

// LoopStmt provide "for expr" expression statement.
type LoopStmt struct {
	StmtImpl
	Expr  Expr
	Stmts []Stmt
}

// BreakStmt provide "break" expression statement.
type BreakStmt struct {
	StmtImpl
}

// ContinueStmt provide "continue" expression statement.
type ContinueStmt struct {
	StmtImpl
}

// ForStmt provide "return" expression statement.
type ReturnStmt struct {
	StmtImpl
	Expr Expr
}

// ThrowStmt provide "throw" expression statement.
type ThrowStmt struct {
	StmtImpl
	Expr Expr
}

// ModuleStmt provide "module" expression statement.
type ModuleStmt struct {
	StmtImpl
	Name  string
	Stmts []Stmt
}

// VarStmt provide statement to let variables in current scope.
type VarStmt struct {
	StmtImpl
	Names []string
	Exprs []Expr
}

// LetStmt provide statement to let variables.
// This can store multiple identity and values for "a,b=1,2".
type LetStmt struct {
	StmtImpl
	Names []string
	Exprs []Expr
}
