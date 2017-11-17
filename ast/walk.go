package ast

import (
	"fmt"
	"reflect"
)

type WalkStatementFunc func(Stmt) error
type WalkExpressionFunc func(Expr) error

// WalkStmts walks the ASTs associated with stmts calling the statement and expression functions if provided
// if either function returns an error, the Walk function returns the error
// The function is useful for validating an AST or prepopulating values prior to running
func WalkStmts(stmts []Stmt, sf WalkStatementFunc, ef WalkExpressionFunc) error {
	for _, stmt := range stmts {
		if err := walkStmt(stmt, sf, ef); err != nil {
			return err
		}
	}
	return nil
}

// WalkExprs recursively walkts a list of expressions calling the appropriate function at each
func WalkExprs(exprs []Expr, sf WalkStatementFunc, ef WalkExpressionFunc) error {
	for _, expr := range exprs {
		if err := walkExpr(expr, sf, ef); err != nil {
			return err
		}
	}
	return nil
}

func walkStmt(stmt Stmt, sf WalkStatementFunc, ef WalkExpressionFunc) error {
	//short circuit out if there are no functions
	if stmt == nil || sf == nil && ef == nil {
		return nil
	}
	if err := callStmtFunc(stmt, sf); err != nil {
		return err
	}
	switch stmt := stmt.(type) {
	case *BreakStmt:
	case *ContinueStmt:
	case *ReturnStmt:
		return WalkExprs(stmt.Exprs, sf, ef)
	case *ExprStmt:
		return walkExpr(stmt.Expr, sf, ef)
	case *VarStmt:
		return WalkExprs(stmt.Exprs, sf, ef)
	case *LetsStmt:
		if err := WalkExprs(stmt.Rhss, sf, ef); err != nil {
			return err
		}
		return WalkExprs(stmt.Lhss, sf, ef)
	case *IfStmt:
		if err := walkExpr(stmt.If, sf, ef); err != nil {
			return err
		}
		if err := WalkStmts(stmt.Then, sf, ef); err != nil {
			return err
		}
		if err := WalkStmts(stmt.ElseIf, sf, ef); err != nil {
			return err
		}
		if err := WalkStmts(stmt.Else, sf, ef); err != nil {
			return err
		}
	case *TryStmt:
		if err := WalkStmts(stmt.Try, sf, ef); err != nil {
			return err
		}
		if err := WalkStmts(stmt.Catch, sf, ef); err != nil {
			return err
		}
		if err := WalkStmts(stmt.Finally, sf, ef); err != nil {
			return err
		}
	case *LoopStmt:
		if err := walkExpr(stmt.Expr, sf, ef); err != nil {
			return err
		}
		if err := WalkStmts(stmt.Stmts, sf, ef); err != nil {
			return err
		}
	case *ForStmt:
		if err := walkExpr(stmt.Value, sf, ef); err != nil {
			return err
		}
		if err := WalkStmts(stmt.Stmts, sf, ef); err != nil {
			return err
		}
	case *CForStmt:
		if err := walkExpr(stmt.Expr1, sf, ef); err != nil {
			return err
		}
		if err := walkExpr(stmt.Expr2, sf, ef); err != nil {
			return err
		}
		if err := walkExpr(stmt.Expr3, sf, ef); err != nil {
			return err
		}
		if err := WalkStmts(stmt.Stmts, sf, ef); err != nil {
			return err
		}
	case *ThrowStmt:
		if err := walkExpr(stmt.Expr, sf, ef); err != nil {
			return err
		}
	case *ModuleStmt:
		if err := WalkStmts(stmt.Stmts, sf, ef); err != nil {
			return err
		}
	case *SwitchStmt:
		if err := walkExpr(stmt.Expr, sf, ef); err != nil {
			return err
		}
		for _, ss := range stmt.Cases {
			if ssd, ok := ss.(*DefaultStmt); ok {
				if err := WalkStmts(ssd.Stmts, sf, ef); err != nil {
					return err
				}
				continue
			}
			if err := walkExpr(ss.(*CaseStmt).Expr, sf, ef); err != nil {
				return err
			}
			if err := WalkStmts(ss.(*CaseStmt).Stmts, sf, ef); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("Unknown statement %v", reflect.TypeOf(stmt))
	}
	return nil
}

func walkExpr(expr Expr, sf WalkStatementFunc, ef WalkExpressionFunc) error {
	//short circuit out if there are no functions
	if expr == nil || sf == nil && ef == nil {
		return nil
	}
	if err := callExprFunc(expr, ef); err != nil {
		return err
	}
	switch expr := expr.(type) {
	case *NumberExpr:
	case *IdentExpr:
	case *MemberExpr:
		return walkExpr(expr.Expr, sf, ef)
	case *StringExpr:
	case *ItemExpr:
		if err := walkExpr(expr.Value, sf, ef); err != nil {
			return err
		}
		return walkExpr(expr.Index, sf, ef)
	case *SliceExpr:
		if err := walkExpr(expr.Value, sf, ef); err != nil {
			return err
		}
		if err := walkExpr(expr.Begin, sf, ef); err != nil {
			return err
		}
		return walkExpr(expr.End, sf, ef)
	case *ArrayExpr:
		return WalkExprs(expr.Exprs, sf, ef)
	case *MapExpr:
		for _, expr := range expr.MapExpr {
			if err := walkExpr(expr, sf, ef); err != nil {
				return err
			}
		}
	case *DerefExpr:
		return walkExpr(expr.Expr, sf, ef)
	case *AddrExpr:
		return walkExpr(expr.Expr, sf, ef)
	case *UnaryExpr:
		return walkExpr(expr.Expr, sf, ef)
	case *ParenExpr:
		return walkExpr(expr.SubExpr, sf, ef)
	case *FuncExpr:
		return WalkStmts(expr.Stmts, sf, ef)
	case *AssocExpr:
		return walkExpr(expr.Lhs, sf, ef)
	case *LetExpr:
		if err := walkExpr(expr.Lhs, sf, ef); err != nil {
			return err
		}
		return walkExpr(expr.Rhs, sf, ef)
	case *LetsExpr:
		if err := WalkExprs(expr.Lhss, sf, ef); err != nil {
			return err
		}
		return WalkExprs(expr.Rhss, sf, ef)
	case *NewExpr:
	case *BinOpExpr:
		if err := walkExpr(expr.Lhs, sf, ef); err != nil {
			return err
		}
		return walkExpr(expr.Rhs, sf, ef)
	case *ConstExpr:
	case *AnonCallExpr:
		if err := walkExpr(expr.Expr, sf, ef); err != nil {
			return err
		}
		return walkExpr(&CallExpr{Func: nil, SubExprs: expr.SubExprs, VarArg: expr.VarArg, Go: expr.Go}, sf, ef)
	case *CallExpr:
		return WalkExprs(expr.SubExprs, sf, ef)
	case *TernaryOpExpr:
		if err := walkExpr(expr.Expr, sf, ef); err != nil {
			return err
		}
		if err := walkExpr(expr.Lhs, sf, ef); err != nil {
			return err
		}
		return walkExpr(expr.Rhs, sf, ef)
	case *MakeExpr:
	case *MakeChanExpr:
		return walkExpr(expr.SizeExpr, sf, ef)
	case *MakeArrayExpr:
		if err := walkExpr(expr.LenExpr, sf, ef); err != nil {
			return err
		}
		return walkExpr(expr.CapExpr, sf, ef)
	case *ChanExpr:
		if err := walkExpr(expr.Rhs, sf, ef); err != nil {
			return err
		}
		return walkExpr(expr.Lhs, sf, ef)
	default:
		return fmt.Errorf("Unknown expression %v", reflect.TypeOf(expr))
	}
	return nil
}

func callStmtFunc(stmt Stmt, sf WalkStatementFunc) error {
	if stmt == nil || sf == nil {
		return nil
	}
	return sf(stmt)
}

func callExprFunc(expr Expr, ef WalkExpressionFunc) error {
	if expr == nil || ef == nil {
		return nil
	}
	return ef(expr)
}
