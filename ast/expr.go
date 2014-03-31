package ast

type Expr interface {
	expr()
}

type ExprImpl struct {
}

func (x *ExprImpl) expr() {}

type NumberExpr struct {
	ExprImpl
	Lit string
}

type StringExpr struct {
	ExprImpl
	Lit string
}

type ArrayExpr struct {
	ExprImpl
	Exprs []Expr
}

type PairExpr struct {
	ExprImpl
	Key   string
	Value Expr
}

type MapExpr struct {
	ExprImpl
	MapExpr map[string]Expr
}

type IdentExpr struct {
	ExprImpl
	Lit string
}

type UnaryMinusExpr struct {
	ExprImpl
	SubExpr Expr
}

type ParenExpr struct {
	ExprImpl
	SubExpr Expr
}

type BinOpExpr struct {
	ExprImpl
	Lhs      Expr
	Operator string
	Rhs      Expr
}

type CallExpr struct {
	ExprImpl
	Name     string
	SubExprs []Expr
}

type ItemExpr struct {
	ExprImpl
	Value Expr
	Index Expr
}

type FuncExpr struct {
	ExprImpl
	Stmts  []Stmt
	Args   []string
	VarArg bool
}

type LetExpr struct {
	ExprImpl
	Name string
	Expr Expr
}

type NewExpr struct {
	ExprImpl
	Name     string
	SubExprs []Expr
}

type ConstExpr struct {
	ExprImpl
	Value interface{}
}
