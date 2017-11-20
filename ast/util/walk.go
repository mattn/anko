package util

import (
	"fmt"
	"reflect"

	"github.com/mattn/anko/ast"
)

type WalkStatementFunc func(ast.Stmt) error
type WalkExpressionFunc func(ast.Expr) error

// WalkStmts walks the ASTs associated with stmts calling the statement and expression functions if provided
// if either function returns an error, the Walk function returns the error
// The function is useful for validating an AST or prepopulating values prior to running
func WalkStmts(stmts []ast.Stmt, sf WalkStatementFunc, ef WalkExpressionFunc) error {
	for _, stmt := range stmts {
		if err := walkStmt(stmt, sf, ef); err != nil {
			return err
		}
	}
	return nil
}

// WalkExprs recursively walkts a list of expressions calling the appropriate function at each
func WalkExprs(exprs []ast.Expr, sf WalkStatementFunc, ef WalkExpressionFunc) error {
	for _, expr := range exprs {
		if err := walkExpr(expr, sf, ef); err != nil {
			return err
		}
	}
	return nil
}

func walkStmt(stmt ast.Stmt, sf WalkStatementFunc, ef WalkExpressionFunc) error {
	//short circuit out if there are no functions
	if stmt == nil || (sf == nil && ef == nil) {
		return nil
	}
	if err := callStmtFunc(stmt, sf); err != nil {
		return err
	}
	switch stmt := stmt.(type) {
	case *ast.BreakStmt:
	case *ast.ContinueStmt:
	case *ast.ReturnStmt:
		return WalkExprs(stmt.Exprs, sf, ef)
	case *ast.ExprStmt:
		return walkExpr(stmt.Expr, sf, ef)
	case *ast.VarStmt:
		return WalkExprs(stmt.Exprs, sf, ef)
	case *ast.LetsStmt:
		if err := WalkExprs(stmt.Rhss, sf, ef); err != nil {
			return err
		}
		return WalkExprs(stmt.Lhss, sf, ef)
	case *ast.IfStmt:
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
	case *ast.TryStmt:
		if err := WalkStmts(stmt.Try, sf, ef); err != nil {
			return err
		}
		if err := WalkStmts(stmt.Catch, sf, ef); err != nil {
			return err
		}
		if err := WalkStmts(stmt.Finally, sf, ef); err != nil {
			return err
		}
	case *ast.LoopStmt:
		if err := walkExpr(stmt.Expr, sf, ef); err != nil {
			return err
		}
		if err := WalkStmts(stmt.Stmts, sf, ef); err != nil {
			return err
		}
	case *ast.ForStmt:
		if err := walkExpr(stmt.Value, sf, ef); err != nil {
			return err
		}
		if err := WalkStmts(stmt.Stmts, sf, ef); err != nil {
			return err
		}
	case *ast.CForStmt:
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
	case *ast.ThrowStmt:
		if err := walkExpr(stmt.Expr, sf, ef); err != nil {
			return err
		}
	case *ast.ModuleStmt:
		if err := WalkStmts(stmt.Stmts, sf, ef); err != nil {
			return err
		}
	case *ast.SwitchStmt:
		if err := walkExpr(stmt.Expr, sf, ef); err != nil {
			return err
		}
		for _, ss := range stmt.Cases {
			if ssd, ok := ss.(*ast.DefaultStmt); ok {
				if err := WalkStmts(ssd.Stmts, sf, ef); err != nil {
					return err
				}
				continue
			}
			if err := walkExpr(ss.(*ast.CaseStmt).Expr, sf, ef); err != nil {
				return err
			}
			if err := WalkStmts(ss.(*ast.CaseStmt).Stmts, sf, ef); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("Unknown statement %v", reflect.TypeOf(stmt))
	}
	return nil
}

func walkExpr(expr ast.Expr, sf WalkStatementFunc, ef WalkExpressionFunc) error {
	//short circuit out if there are no functions
	if expr == nil || (sf == nil && ef == nil) {
		return nil
	}
	if err := callExprFunc(expr, ef); err != nil {
		return err
	}
	switch expr := expr.(type) {
	case *ast.NumberExpr:
	case *ast.IdentExpr:
	case *ast.MemberExpr:
		return walkExpr(expr.Expr, sf, ef)
	case *ast.StringExpr:
	case *ast.ItemExpr:
		if err := walkExpr(expr.Value, sf, ef); err != nil {
			return err
		}
		return walkExpr(expr.Index, sf, ef)
	case *ast.SliceExpr:
		if err := walkExpr(expr.Value, sf, ef); err != nil {
			return err
		}
		if err := walkExpr(expr.Begin, sf, ef); err != nil {
			return err
		}
		return walkExpr(expr.End, sf, ef)
	case *ast.ArrayExpr:
		return WalkExprs(expr.Exprs, sf, ef)
	case *ast.MapExpr:
		for _, expr := range expr.MapExpr {
			if err := walkExpr(expr, sf, ef); err != nil {
				return err
			}
		}
	case *ast.DerefExpr:
		return walkExpr(expr.Expr, sf, ef)
	case *ast.AddrExpr:
		return walkExpr(expr.Expr, sf, ef)
	case *ast.UnaryExpr:
		return walkExpr(expr.Expr, sf, ef)
	case *ast.ParenExpr:
		return walkExpr(expr.SubExpr, sf, ef)
	case *ast.FuncExpr:
		return WalkStmts(expr.Stmts, sf, ef)
	case *ast.AssocExpr:
		return walkExpr(expr.Lhs, sf, ef)
	case *ast.LetExpr:
		if err := walkExpr(expr.Lhs, sf, ef); err != nil {
			return err
		}
		return walkExpr(expr.Rhs, sf, ef)
	case *ast.LetsExpr:
		if err := WalkExprs(expr.Lhss, sf, ef); err != nil {
			return err
		}
		return WalkExprs(expr.Rhss, sf, ef)
	case *ast.NewExpr:
	case *ast.BinOpExpr:
		if err := walkExpr(expr.Lhs, sf, ef); err != nil {
			return err
		}
		return walkExpr(expr.Rhs, sf, ef)
	case *ast.ConstExpr:
	case *ast.AnonCallExpr:
		if err := walkExpr(expr.Expr, sf, ef); err != nil {
			return err
		}
		return walkExpr(&ast.CallExpr{Func: nil, SubExprs: expr.SubExprs, VarArg: expr.VarArg, Go: expr.Go}, sf, ef)
	case *ast.CallExpr:
		return WalkExprs(expr.SubExprs, sf, ef)
	case *ast.TernaryOpExpr:
		if err := walkExpr(expr.Expr, sf, ef); err != nil {
			return err
		}
		if err := walkExpr(expr.Lhs, sf, ef); err != nil {
			return err
		}
		return walkExpr(expr.Rhs, sf, ef)
	case *ast.MakeExpr:
	case *ast.MakeChanExpr:
		return walkExpr(expr.SizeExpr, sf, ef)
	case *ast.MakeArrayExpr:
		if err := walkExpr(expr.LenExpr, sf, ef); err != nil {
			return err
		}
		return walkExpr(expr.CapExpr, sf, ef)
	case *ast.ChanExpr:
		if err := walkExpr(expr.Rhs, sf, ef); err != nil {
			return err
		}
		return walkExpr(expr.Lhs, sf, ef)
	default:
		return fmt.Errorf("Unknown expression %v", reflect.TypeOf(expr))
	}
	return nil
}

func callStmtFunc(stmt ast.Stmt, sf WalkStatementFunc) error {
	if stmt == nil || sf == nil {
		return nil
	}
	return sf(stmt)
}

func callExprFunc(expr ast.Expr, ef WalkExpressionFunc) error {
	if expr == nil || ef == nil {
		return nil
	}
	return ef(expr)
}
