package ast

// Expr provides all of interfaces for expression.
type Expr interface {
	Pos
	expr()
}

// ExprImpl provide commonly implementations for Expr.
type ExprImpl struct {
	PosImpl // ExprImpl provide Pos() function.
}

// expr provide restraint interface.
func (x *ExprImpl) expr() {}

// NumberExpr provide Number expression.
type NumberExpr struct {
	ExprImpl
	Lit string
}

// StringExpr provide String expression.
type StringExpr struct {
	ExprImpl
	Lit string
}

// ArrayExpr provide Array expression.
type ArrayExpr struct {
	ExprImpl
	Exprs []Expr
}

// PairExpr provide one of Map key/value pair.
type PairExpr struct {
	ExprImpl
	Key   string
	Value Expr
}

// MapExpr provide Map expression.
type MapExpr struct {
	ExprImpl
	MapExpr map[string]Expr
}

// IdentExpr provide identity expression.
type IdentExpr struct {
	ExprImpl
	Lit string
}

// UnaryExpr provide unary minus expression. ex: -1, ^1, ~1.
type UnaryExpr struct {
	ExprImpl
	Operator string
	Expr     Expr
}

// ParenExpr provide parent block expression.
type ParenExpr struct {
	ExprImpl
	SubExpr Expr
}

// BinOpExpr provide binary operator expression.
type BinOpExpr struct {
	ExprImpl
	Lhs      Expr
	Operator string
	Rhs      Expr
}

type TernaryOpExpr struct {
	ExprImpl
	Expr Expr
	Lhs  Expr
	Rhs  Expr
}

// CallExpr provide calling expression.
type CallExpr struct {
	ExprImpl
	Func     interface{}
	Name     string
	SubExprs []Expr
}

// AnonCallExpr provide anonymous calling expression. ex: func(){}().
type AnonCallExpr struct {
	ExprImpl
	Expr     Expr
	SubExprs []Expr
}

// MemberExpr provide expression to refer menber.
type MemberExpr struct {
	ExprImpl
	Expr   Expr
	Method string
}

// ItemExpr provide expression to refer Map/Array item.
type ItemExpr struct {
	ExprImpl
	Value Expr
	Index Expr
}

// FuncExpr provide function expression.
type FuncExpr struct {
	ExprImpl
	Name   string
	Stmts  []Stmt
	Args   []string
	VarArg bool
}

// LetExpr provide expression to let variable.
type LetExpr struct {
	ExprImpl
	Lhs      Expr
	Rhs      Expr
}

// LetsExpr provide multiple expression of let.
type LetsExpr struct {
	ExprImpl
	Lhss     []Expr
	Operator string
	Rhss     []Expr
}

// AssocExpr provide expression to assoc operation.
type AssocExpr struct {
	ExprImpl
	Name     string
	Operator string
	Expr     Expr
}

// NewExpr provide expression to make new instance.
type NewExpr struct {
	ExprImpl
	Name     string
	SubExprs []Expr
}

// ConstExpr provide expression for constant variable.
type ConstExpr struct {
	ExprImpl
	Value interface{}
}
