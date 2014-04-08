package vm

import (
	"errors"
	"fmt"
	"github.com/mattn/anko/ast"
	"math"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

var Nil unsafe.Pointer
var NilValue = reflect.ValueOf(Nil)
var TrueValue = reflect.ValueOf(true)
var FalseValue = reflect.ValueOf(false)

// Error provides a convenient interface for handling runtime error.
// It can be Error interface with type cast which can call Pos().
type Error struct {
	message string
	pos     ast.Position
}

var BreakError = errors.New("Unexpected break statement")
var ContinueError = errors.New("Unexpected continue statement")
var ReturnError = errors.New("Unexpected error")

// NewStringError makes error interface with message.
func NewStringError(pos ast.Pos, err string) error {
	return &Error{message: err, pos: pos.GetPos()}
}

// NewStringError makes error interface with message.
func NewErrorf(pos ast.Pos, format string, args ...interface{}) error {
	return &Error{message: fmt.Sprintf(format, args...), pos: pos.GetPos()}
}

// NewError makes error interface with message. This doesn't overwrite last error.
func NewError(pos ast.Pos, err error) error {
	if err == nil {
		return nil
	}
	if err == BreakError || err == ContinueError || err == ReturnError {
		return err
	}
	if ee, ok := err.(*Error); ok {
		return ee
	}
	return &Error{message: err.Error(), pos: pos.GetPos()}
}

// Error return the error message.
func (e *Error) Error() string {
	return e.message
}

// Pos return the position of error.
func (e *Error) Pos() ast.Position {
	return e.pos
}

// Func is function interface to reflect functions internaly.
type Func func(args ...reflect.Value) (reflect.Value, error)

func ToFunc(f Func) reflect.Value {
	return reflect.ValueOf(f)
}

// Run execute statements in the environment which specified.
func Run(stmts []ast.Stmt, env *Env) (reflect.Value, error) {
	rv := NilValue
	var err error
	for _, stmt := range stmts {
		if _, ok := stmt.(*ast.BreakStmt); ok {
			return NilValue, BreakError
		}
		if _, ok := stmt.(*ast.ContinueStmt); ok {
			return NilValue, ContinueError
		}
		rv, err = RunSingleStmt(stmt, env)
		if _, ok := stmt.(*ast.ReturnStmt); ok {
			return reflect.ValueOf(rv), ReturnError
		}
		if err != nil {
			return rv, NewError(stmt, err)
		}
	}
	return rv, nil
}

// RunSingleStmt execute one statement in the environment which specified.
func RunSingleStmt(stmt ast.Stmt, env *Env) (reflect.Value, error) {
	switch stmt := stmt.(type) {
	case *ast.ExprStmt:
		rv, err := invokeExpr(stmt.Expr, env)
		if err != nil {
			return rv, NewError(stmt, err)
		}
		return rv, nil
	case *ast.VarStmt:
		rv := NilValue
		var err error
		rvs := []reflect.Value{}
		for _, expr := range stmt.Exprs {
			rv, err = invokeExpr(expr, env)
			if err != nil {
				return rv, NewError(expr, err)
			}
			rvs = append(rvs, rv)
		}
		result := []interface{}{}
		for i, name := range stmt.Names {
			if i < len(rvs) {
				env.Define(name, rvs[i])
				result = append(result, rvs[i].Interface())
			}
		}
		return reflect.ValueOf(result), nil
	case *ast.IfStmt:
		// If
		rv, err := invokeExpr(stmt.If, env)
		if err != nil {
			return rv, NewError(stmt, err)
		}
		if toBool(rv) {
			// Then
			rv, err = Run(stmt.Then, env.NewEnv())
			if err != nil {
				return rv, NewError(stmt, err)
			}
			return rv, nil
		}
		done := false
		if len(stmt.ElseIf) > 0 {
			for _, stmt := range stmt.ElseIf {
				stmt_if := stmt.(*ast.IfStmt)
				// ElseIf
				rv, err = invokeExpr(stmt_if.If, env)
				if err != nil {
					return rv, NewError(stmt, err)
				}
				if !toBool(rv) {
					continue
				}
				// ElseIf Then
				done = true
				rv, err = Run(stmt_if.Then, env)
				if err != nil {
					return rv, NewError(stmt, err)
				}
				break
			}
		}
		if !done && len(stmt.Else) > 0 {
			// Else
			rv, err = Run(stmt.Else, env.NewEnv())
			if err != nil {
				return rv, NewError(stmt, err)
			}
		}
		return rv, nil
	case *ast.TryStmt:
		_, err := Run(stmt.Try, env.NewEnv())
		if err != nil {
			cenv := env.NewEnv()
			if stmt.Var != "" {
				cenv.Define(stmt.Var, reflect.ValueOf(err))
			}
			_, e1 := Run(stmt.Catch, cenv)
			if e1 != nil {
				err = NewError(stmt.Catch[0], e1)
			} else {
				err = nil
			}
		}
		if len(stmt.Finally) > 0 {
			_, e2 := Run(stmt.Finally, env.NewEnv())
			if e2 != nil {
				err = NewError(stmt.Finally[0], e2)
			}
		}
		return NilValue, NewError(stmt, err)
	case *ast.LoopStmt:
		newenv := env.NewEnv()
		for {
			if stmt.Expr != nil {
				ev, ee := invokeExpr(stmt.Expr, newenv)
				if ee != nil {
					return ev, ee
				}
				if !toBool(ev) {
					break
				}
			}

			rv, err := Run(stmt.Stmts, newenv)
			if err != nil {
				if err == BreakError {
					err = nil
					break
				}
				if err == ContinueError {
					err = nil
					continue
				}
				if err == ReturnError {
					err = nil
					break
				}
				return rv, NewError(stmt, err)
			}
		}
		return NilValue, nil
	case *ast.ForStmt:
		val, ee := invokeExpr(stmt.Value, env)
		if ee != nil {
			return val, ee
		}
		if val.Kind() != reflect.Array && val.Kind() != reflect.Slice {
			return NilValue, NewStringError(stmt, "Invalid operation for non-array value")
		}
		newenv := env.NewEnv()
		for i := 0; i < val.Len(); i++ {
			newenv.Define(stmt.Var, val.Index(i))
			rv, err := Run(stmt.Stmts, newenv)
			if err != nil {
				if err == BreakError {
					err = nil
					break
				}
				if err == ContinueError {
					err = nil
					continue
				}
				if err == ReturnError {
					err = nil
					break
				}
				return rv, NewError(stmt, err)
			}
		}
		return NilValue, nil
	case *ast.ReturnStmt:
		rvs := []interface{}{}
		switch len(stmt.Exprs) {
		case 0:
			return NilValue, nil
		case 1:
			rv, err := invokeExpr(stmt.Exprs[0], env)
			if err != nil {
				return rv, NewError(stmt, err)
			}
			return rv, nil
		}
		for _, expr := range stmt.Exprs {
			rv, err := invokeExpr(expr, env)
			if err != nil {
				return rv, NewError(stmt, err)
			}
			if isNil(rv) {
				rvs = append(rvs, Nil)
			} else if rv.IsValid() {
				rvs = append(rvs, rv.Interface())
			} else {
				rvs = append(rvs, Nil)
			}
		}
		return reflect.ValueOf(rvs), nil
	case *ast.ThrowStmt:
		rv, err := invokeExpr(stmt.Expr, env)
		if err != nil {
			return rv, NewError(stmt, err)
		}
		if !rv.IsValid() {
			return NilValue, NewError(stmt, err)
		}
		return rv, NewStringError(stmt, fmt.Sprint(rv.Interface()))
	case *ast.ModuleStmt:
		newenv := env.NewEnv()
		newenv.SetName(stmt.Name)
		rv, err := Run(stmt.Stmts, newenv)
		if err != nil {
			return rv, NewError(stmt, err)
		}
		env.DefineGlobal(stmt.Name, reflect.ValueOf(newenv))
		return rv, nil
	case *ast.SwitchStmt:
		rv, err := invokeExpr(stmt.Expr, env)
		if err != nil {
			return rv, NewError(stmt, err)
		}
		done := false
		var default_stmt *ast.DefaultStmt
		for _, ss := range stmt.Cases {
			if ssd, ok := ss.(*ast.DefaultStmt); ok {
				default_stmt = ssd
				continue
			}
			case_stmt := ss.(*ast.CaseStmt)
			cv, err := invokeExpr(case_stmt.Expr, env)
			if err != nil {
				return rv, NewError(stmt, err)
			}
			if !equal(rv, cv) {
				continue
			}
			rv, err = Run(case_stmt.Stmts, env)
			if err != nil {
				return rv, NewError(stmt, err)
			}
			done = true
			break
		}
		if !done && default_stmt != nil {
			rv, err = Run(default_stmt.Stmts, env)
			if err != nil {
				return rv, NewError(stmt, err)
			}
		}
		return rv, nil
	default:
		return NilValue, NewStringError(stmt, "Unknown statement")
	}
}

// toString convert all reflect.Value-s into string.
func toString(v reflect.Value) string {
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	if v.Kind() == reflect.String {
		return v.String()
	}
	return fmt.Sprint(v.Interface())
}

// toBool convert all reflect.Value-s into bool.
func toBool(v reflect.Value) bool {
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		return v.Float() != 0.0
	case reflect.Int, reflect.Int32, reflect.Int64:
		return v.Int() != 0
	case reflect.Bool:
		return v.Bool()
	case reflect.String:
		if v.String() == "true" {
			return true
		}
		if toInt64(v) != 0 {
			return true
		}
	}
	return false
}

// toFloat64 convert all reflect.Value-s into float64.
func toFloat64(v reflect.Value) float64 {
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		return v.Float()
	case reflect.Int, reflect.Int32, reflect.Int64:
		return float64(v.Int())
	}
	return 0.0
}

func isNil(v reflect.Value) bool {
	if v.Kind().String() == "unsafe.Pointer" {
		return true
	}
	if (v.Kind() == reflect.Interface || v.Kind() == reflect.Ptr) && v.IsNil() {
		return true
	}
	return false
}

// equal return true when lhsV and rhsV is same value.
func equal(lhsV, rhsV reflect.Value) bool {
	if isNil(lhsV) && isNil(rhsV) {
		return true
	}
	if lhsV.Kind() == reflect.Interface || lhsV.Kind() == reflect.Ptr {
		lhsV = lhsV.Elem()
	}
	if rhsV.Kind() == reflect.Interface || rhsV.Kind() == reflect.Ptr {
		rhsV = rhsV.Elem()
	}
	if !lhsV.IsValid() || !rhsV.IsValid() {
		return true
	}
	if lhsV.CanInterface() && rhsV.CanInterface() {
		return reflect.DeepEqual(lhsV.Interface(), rhsV.Interface())
	}
	return reflect.DeepEqual(lhsV, rhsV)
}

// toInt64 convert all reflect.Value-s into int64.
func toInt64(v reflect.Value) int64 {
	if v.Kind() == reflect.Interface {
		v = v.Elem()
	}
	switch v.Kind() {
	case reflect.Float32, reflect.Float64:
		return int64(v.Float())
	case reflect.Int, reflect.Int32, reflect.Int64:
		return v.Int()
	case reflect.String:
		i, err := strconv.Atoi(v.String())
		if err == nil {
			return int64(i)
		}
	}
	return 0
}

func invokeLetExpr(expr ast.Expr, rv reflect.Value, env *Env) (reflect.Value, error) {
	switch lhs := expr.(type) {
	case *ast.IdentExpr:
		if env.Set(lhs.Lit, rv) != nil {
			if strings.Contains(lhs.Lit, ".") {
				return NilValue, NewErrorf(expr, "Undefined symbol '%s'", lhs.Lit)
			}
			env.Define(lhs.Lit, rv)
		}
		return rv, nil
	case *ast.MemberExpr:
		v, err := invokeExpr(lhs.Expr, env)
		if err != nil {
			return v, NewError(expr, err)
		}

		if v.Kind() == reflect.Interface {
			v = v.Elem()
		}
		if v.Kind() == reflect.Slice {
			v = v.Index(0)
		}

		if !v.IsValid() {
			return NilValue, NewStringError(expr, "Cannot assignable")
		}

		if v.Kind() == reflect.Ptr {
			v = v.Elem()
		}
		if v.Kind() == reflect.Struct {
			v = v.FieldByName(lhs.Name)
		} else if v.Kind() == reflect.Map {
			v = v.MapIndex(reflect.ValueOf(lhs.Name))
		}

		if !v.CanSet() {
			return NilValue, NewStringError(expr, "Cannot assignable")
		}
		v.Set(rv)
		return v, nil
	case *ast.ItemExpr:
		v, err := invokeExpr(lhs.Value, env)
		if err != nil {
			return v, NewError(expr, err)
		}
		i, err := invokeExpr(lhs.Index, env)
		if err != nil {
			return i, NewError(expr, err)
		}
		if v.Kind() == reflect.Interface {
			v = v.Elem()
		}
		if v.Kind() == reflect.Array || v.Kind() == reflect.Slice {
			if i.Kind() != reflect.Int && i.Kind() != reflect.Int64 {
				return NilValue, NewStringError(expr, "Array index should be int")
			}
			ii := int(i.Int())
			if ii < 0 || ii >= v.Len() {
				return NilValue, NewStringError(expr, "Cannot assignable")
			}
			vv := v.Index(ii)
			if !vv.CanSet() {
				return NilValue, NewStringError(expr, "Cannot assignable")
			}
			vv.Set(rv)
			return rv, nil
		}
		if v.Kind() == reflect.Map {
			if i.Kind() != reflect.String {
				return NilValue, NewStringError(expr, "Map key should be string")
			}
			v.SetMapIndex(i, rv)
			return rv, nil
		}
		return v, NewStringError(expr, "Invalid operation")
	}
	return NilValue, NewStringError(expr, "Invalid operation")
}

// invokeExpr evaluate one expression.
func invokeExpr(expr ast.Expr, env *Env) (reflect.Value, error) {
	switch e := expr.(type) {
	case *ast.NumberExpr:
		if strings.Contains(e.Lit, ".") {
			v, err := strconv.ParseFloat(e.Lit, 64)
			if err != nil {
				return NilValue, NewError(expr, err)
			}
			return reflect.ValueOf(float64(v)), nil
		}
		v, err := strconv.Atoi(e.Lit)
		if err != nil {
			return NilValue, NewError(expr, err)
		}
		return reflect.ValueOf(int64(v)), nil
	case *ast.IdentExpr:
		return env.Get(e.Lit)
	case *ast.StringExpr:
		return reflect.ValueOf(e.Lit), nil
	case *ast.ArrayExpr:
		a := make([]interface{}, len(e.Exprs))
		for i, expr := range e.Exprs {
			arg, err := invokeExpr(expr, env)
			if err != nil {
				return arg, NewError(expr, err)
			}
			a[i] = arg.Interface()
		}
		return reflect.ValueOf(a), nil
	case *ast.MapExpr:
		m := make(map[string]interface{})
		for k, expr := range e.MapExpr {
			v, err := invokeExpr(expr, env)
			if err != nil {
				return v, NewError(expr, err)
			}
			m[k] = v.Interface()
		}
		return reflect.ValueOf(m), nil
	case *ast.DerefExpr:
		v := NilValue
		var err error
		switch ee := e.Expr.(type) {
		case *ast.IdentExpr:
			v, err = env.Get(ee.Lit)
			if err != nil {
				return v, err
			}
		case *ast.MemberExpr:
			v, err := invokeExpr(ee.Expr, env)
			if err != nil {
				return v, NewError(expr, err)
			}
			if v.Kind() == reflect.Interface {
				v = v.Elem()
			}
			if v.Kind() == reflect.Slice {
				v = v.Index(0)
			}
			if v.IsValid() && v.CanInterface() {
				if vme, ok := v.Interface().(*Env); ok {
					m, err := vme.Get(ee.Name)
					if !m.IsValid() || err != nil {
						return NilValue, NewStringError(expr, fmt.Sprintf("Invalid operation '%s'", ee.Name))
					}
					return m, nil
				}
			}

			m := v.MethodByName(ee.Name)
			if !m.IsValid() {
				if v.Kind() == reflect.Ptr {
					v = v.Elem()
				}
				if v.Kind() == reflect.Struct {
					m = v.FieldByName(ee.Name)
					if !m.IsValid() {
						return NilValue, NewStringError(expr, fmt.Sprintf("Invalid operation '%s'", ee.Name))
					}
				} else if v.Kind() == reflect.Map {
					m = v.MapIndex(reflect.ValueOf(ee.Name))
					if !m.IsValid() {
						return NilValue, NewStringError(expr, fmt.Sprintf("Invalid operation '%s'", ee.Name))
					}
				} else {
					return NilValue, NewStringError(expr, fmt.Sprintf("Invalid operation '%s'", ee.Name))
				}
				v = m
			} else {
				v = m
			}
		default:
			return NilValue, NewStringError(expr, "Invalid operation for the value")
		}
		if v.Kind() != reflect.Ptr {
			return NilValue, NewStringError(expr, "Cannot deference for the value")
		}
		return v.Addr(), nil
	case *ast.AddrExpr:
		v := NilValue
		var err error
		switch ee := e.Expr.(type) {
		case *ast.IdentExpr:
			v, err = env.Get(ee.Lit)
			if err != nil {
				return v, err
			}
		case *ast.MemberExpr:
			v, err := invokeExpr(ee.Expr, env)
			if err != nil {
				return v, NewError(expr, err)
			}
			if v.Kind() == reflect.Interface {
				v = v.Elem()
			}
			if v.Kind() == reflect.Slice {
				v = v.Index(0)
			}
			if v.IsValid() && v.CanInterface() {
				if vme, ok := v.Interface().(*Env); ok {
					m, err := vme.Get(ee.Name)
					if !m.IsValid() || err != nil {
						return NilValue, NewStringError(expr, fmt.Sprintf("Invalid operation '%s'", ee.Name))
					}
					return m, nil
				}
			}

			m := v.MethodByName(ee.Name)
			if !m.IsValid() {
				if v.Kind() == reflect.Ptr {
					v = v.Elem()
				}
				if v.Kind() == reflect.Struct {
					m = v.FieldByName(ee.Name)
					if !m.IsValid() {
						return NilValue, NewStringError(expr, fmt.Sprintf("Invalid operation '%s'", ee.Name))
					}
				} else if v.Kind() == reflect.Map {
					m = v.MapIndex(reflect.ValueOf(ee.Name))
					if !m.IsValid() {
						return NilValue, NewStringError(expr, fmt.Sprintf("Invalid operation '%s'", ee.Name))
					}
				} else {
					return NilValue, NewStringError(expr, fmt.Sprintf("Invalid operation '%s'", ee.Name))
				}
				v = m
			} else {
				v = m
			}
		default:
			return NilValue, NewStringError(expr, "Invalid operation for the value")
		}
		if !v.CanAddr() {
			i := v.Interface()
			return reflect.ValueOf(&i), nil
		}
		return v.Addr(), nil
	case *ast.UnaryExpr:
		v, err := invokeExpr(e.Expr, env)
		if err != nil {
			return v, NewError(expr, err)
		}
		switch e.Operator {
		case "-":
			if v.Kind() == reflect.Float64 {
				return reflect.ValueOf(-v.Float()), nil
			}
			return reflect.ValueOf(-v.Int()), nil
		case "^":
			return reflect.ValueOf(^toInt64(v)), nil
		case "!":
			return reflect.ValueOf(!toBool(v)), nil
		default:
			return NilValue, NewStringError(e, "Unknown operator ''")
		}
	case *ast.ParenExpr:
		v, err := invokeExpr(e.SubExpr, env)
		if err != nil {
			return v, NewError(expr, err)
		}
		return v, nil
	case *ast.FuncExpr:
		f := reflect.ValueOf(func(expr *ast.FuncExpr, env *Env) Func {
			return func(args ...reflect.Value) (reflect.Value, error) {
				if !expr.VarArg {
					if len(args) != len(expr.Args) {
						return NilValue, NewStringError(expr, "Arguments Number of mismatch")
					}
				}
				newenv := env.NewEnv()
				if expr.VarArg {
					newenv.Define(expr.Args[0], reflect.ValueOf(args))
				} else {
					for i, arg := range expr.Args {
						newenv.Define(arg, args[i])
					}
				}
				rr, err := Run(expr.Stmts, newenv)
				if err == ReturnError {
					err = nil
					rr = rr.Interface().(reflect.Value)
				}
				return rr, err
			}
		}(e, env))
		env.Define(e.Name, f)
		return f, nil
	case *ast.MemberExpr:
		v, err := invokeExpr(e.Expr, env)
		if err != nil {
			return v, NewError(expr, err)
		}
		if v.Kind() == reflect.Interface {
			v = v.Elem()
		}
		if v.Kind() == reflect.Slice {
			v = v.Index(0)
		}
		if v.IsValid() && v.CanInterface() {
			if vme, ok := v.Interface().(*Env); ok {
				m, err := vme.Get(e.Name)
				if !m.IsValid() || err != nil {
					return NilValue, NewStringError(expr, fmt.Sprintf("Invalid operation '%s'", e.Name))
				}
				return m, nil
			}
		}

		m := v.MethodByName(e.Name)
		if !m.IsValid() {
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			if v.Kind() == reflect.Struct {
				m = v.FieldByName(e.Name)
				if !m.IsValid() {
					return NilValue, NewStringError(expr, fmt.Sprintf("Invalid operation '%s'", e.Name))
				}
			} else if v.Kind() == reflect.Map {
				m = v.MapIndex(reflect.ValueOf(e.Name))
				if !m.IsValid() {
					return NilValue, NewStringError(expr, fmt.Sprintf("Invalid operation '%s'", e.Name))
				}
			} else {
				return NilValue, NewStringError(expr, fmt.Sprintf("Invalid operation '%s'", e.Name))
			}
		}
		return m, nil
	case *ast.ItemExpr:
		v, err := invokeExpr(e.Value, env)
		if err != nil {
			return v, NewError(expr, err)
		}
		i, err := invokeExpr(e.Index, env)
		if err != nil {
			return i, NewError(expr, err)
		}
		if v.Kind() == reflect.Interface {
			v = v.Elem()
		}
		if v.Kind() == reflect.Array || v.Kind() == reflect.Slice {
			if i.Kind() != reflect.Int && i.Kind() != reflect.Int64 {
				return NilValue, NewStringError(expr, "Array index should be int")
			}
			ii := int(i.Int())
			if ii < 0 || ii >= v.Len() {
				return NilValue, nil
			}
			return v.Index(ii), nil
		}
		if v.Kind() == reflect.Map {
			if i.Kind() != reflect.String {
				return NilValue, NewStringError(expr, "Map key should be string")
			}
			return v.MapIndex(i), nil
		}
		if v.Kind() == reflect.String {
			rs := []rune(v.Interface().(string))
			ii := int(i.Int())
			if ii < 0 || ii >= len(rs) {
				return NilValue, nil
			}
			return reflect.ValueOf(rs[ii]), nil
		}
		return v, NewStringError(expr, "Invalid operation")
	case *ast.AssocExpr:
		v, err := env.Get(e.Name)
		if err != nil {
			return v, err
		}
		switch e.Operator {
		case "++":
			if v.Kind() == reflect.Float64 {
				v = reflect.ValueOf(toFloat64(v) + 1.0)
			} else {
				v = reflect.ValueOf(toInt64(v) + 1)
			}
		case "--":
			if v.Kind() == reflect.Float64 {
				v = reflect.ValueOf(toFloat64(v) - 1.0)
			} else {
				v = reflect.ValueOf(toInt64(v) - 1)
			}
		default:
			v, err = invokeExpr(&ast.BinOpExpr{Lhs: &ast.IdentExpr{Lit: e.Name}, Operator: e.Operator[0:1], Rhs: e.Expr}, env)
			if err != nil {
				return v, err
			}

		}
		if env.Set(e.Name, v) != nil {
			env.Define(e.Name, v)
		}
		return v, nil
	case *ast.LetExpr:
		rv, err := invokeExpr(e.Rhs, env)
		if err != nil {
			return rv, NewError(e, err)
		}
		if rv.Kind() == reflect.Interface {
			rv = rv.Elem()
		}
		return invokeLetExpr(e.Lhs, rv, env)
	case *ast.LetsExpr:
		rv := NilValue
		var err error
		vs := []interface{}{}
		for _, rhs := range e.Rhss {
			rv, err = invokeExpr(rhs, env)
			if err != nil {
				return rv, NewError(rhs, err)
			}
			if rv == NilValue {
				vs = append(vs, nil)
			} else if rv.IsValid() && rv.CanInterface() {
				vs = append(vs, rv.Interface())
			} else {
				vs = append(vs, nil)
			}
		}
		rvs := reflect.ValueOf(vs)
		if len(e.Lhss) > 1 && rvs.Len() == 1 {
			item := rvs.Index(0)
			if item.Kind() == reflect.Interface {
				item = item.Elem()
			}
			if item.Kind() == reflect.Slice {
				rvs = item
			}
		}
		for i, lhs := range e.Lhss {
			if i >= rvs.Len() {
				break
			}
			v := rvs.Index(i)
			if v.Kind() == reflect.Interface {
				v = v.Elem()
			}
			_, err = invokeLetExpr(lhs, v, env)
			if err != nil {
				return rvs, NewError(lhs, err)
			}
		}
		if rvs.Len() == 1 {
			return rvs.Index(0), nil
		}
		return rvs, nil
	//case *ast.NewExpr:
	//	println("NEW")
	//	return NilValue, nil
	case *ast.BinOpExpr:
		lhsV := NilValue
		rhsV := NilValue
		var err error

		lhsV, err = invokeExpr(e.Lhs, env)
		if err != nil {
			return lhsV, NewError(expr, err)
		}
		if lhsV.Kind() == reflect.Interface {
			lhsV = lhsV.Elem()
		}
		if e.Rhs != nil {
			rhsV, err = invokeExpr(e.Rhs, env)
			if err != nil {
				return rhsV, NewError(expr, err)
			}
			if rhsV.Kind() == reflect.Interface {
				rhsV = rhsV.Elem()
			}
		}
		switch e.Operator {
		case "+":
			if lhsV.Kind() == reflect.String || rhsV.Kind() == reflect.String {
				return reflect.ValueOf(toString(lhsV) + toString(rhsV)), nil
			}
			if (lhsV.Kind() == reflect.Array || lhsV.Kind() == reflect.Slice) && (rhsV.Kind() != reflect.Array && rhsV.Kind() != reflect.Slice) {
				return reflect.Append(lhsV, rhsV), nil
			}
			if (lhsV.Kind() == reflect.Array || lhsV.Kind() == reflect.Slice) && (rhsV.Kind() == reflect.Array || rhsV.Kind() == reflect.Slice) {
				return reflect.AppendSlice(lhsV, rhsV), nil
			}
			if lhsV.Kind() == reflect.Float64 || rhsV.Kind() == reflect.Float64 {
				return reflect.ValueOf(toFloat64(lhsV) + toFloat64(rhsV)), nil
			}
			return reflect.ValueOf(toInt64(lhsV) + toInt64(rhsV)), nil
		case "-":
			if lhsV.Kind() == reflect.Float64 || rhsV.Kind() == reflect.Float64 {
				return reflect.ValueOf(toFloat64(lhsV) - toFloat64(rhsV)), nil
			}
			return reflect.ValueOf(toInt64(lhsV) - toInt64(rhsV)), nil
		case "*":
			if lhsV.Kind() == reflect.Float64 || rhsV.Kind() == reflect.Float64 {
				return reflect.ValueOf(toFloat64(lhsV) * toFloat64(rhsV)), nil
			}
			return reflect.ValueOf(toInt64(lhsV) * toInt64(rhsV)), nil
		case "/":
			return reflect.ValueOf(toFloat64(lhsV) / toFloat64(rhsV)), nil
		case "%":
			return reflect.ValueOf(toInt64(lhsV) % toInt64(rhsV)), nil
		case "==":
			return reflect.ValueOf(equal(lhsV, rhsV)), nil
		case "!=":
			return reflect.ValueOf(equal(lhsV, rhsV) == false), nil
		case ">":
			return reflect.ValueOf(toFloat64(lhsV) > toFloat64(rhsV)), nil
		case ">=":
			return reflect.ValueOf(toFloat64(lhsV) >= toFloat64(rhsV)), nil
		case "<":
			return reflect.ValueOf(toFloat64(lhsV) < toFloat64(rhsV)), nil
		case "<=":
			return reflect.ValueOf(toFloat64(lhsV) <= toFloat64(rhsV)), nil
		case "|":
			return reflect.ValueOf(toInt64(lhsV) | toInt64(rhsV)), nil
		case "||":
			if toBool(lhsV) {
				return lhsV, nil
			}
			return rhsV, nil
		case "&":
			return reflect.ValueOf(toInt64(lhsV) & toInt64(rhsV)), nil
		case "&&":
			if toBool(lhsV) {
				return rhsV, nil
			}
			return lhsV, nil
		case "**":
			if lhsV.Kind() == reflect.Float64 {
				return reflect.ValueOf(math.Pow(toFloat64(lhsV), toFloat64(rhsV))), nil
			}
			return reflect.ValueOf(int64(math.Pow(toFloat64(lhsV), toFloat64(rhsV)))), nil
		case ">>":
			return reflect.ValueOf(toInt64(lhsV) >> uint64(toInt64(rhsV))), nil
		case "<<":
			return reflect.ValueOf(toInt64(lhsV) << uint64(toInt64(rhsV))), nil
		default:
			return NilValue, NewStringError(expr, "Unknown operator")
		}
	case *ast.ConstExpr:
		return reflect.ValueOf(e.Value), nil
	case *ast.AnonCallExpr:
		f, err := invokeExpr(e.Expr, env)
		if err != nil {
			return f, NewError(expr, err)
		}
		if f.Kind() != reflect.Func {
			return f, NewStringError(expr, "Unknown function")
		}
		return invokeExpr(&ast.CallExpr{Func: f, SubExprs: e.SubExprs}, env)
	case *ast.CallExpr:
		f := NilValue

		if e.Func != nil {
			f = e.Func.(reflect.Value)
		} else {
			var err error
			ff, err := env.Get(e.Name)
			if err != nil {
				return f, err
			}
			f = ff
		}
		_, isReflect := f.Interface().(Func)

		args := []reflect.Value{}
		for i, expr := range e.SubExprs {
			arg, err := invokeExpr(expr, env)
			if err != nil {
				return arg, NewError(expr, err)
			}

			if !arg.IsValid() {
				arg = NilValue
			} else if i < f.Type().NumIn() {
				if !f.Type().IsVariadic() {
					it := f.Type().In(i)
					if arg.Kind().String() == "unsafe.Pointer" {
						arg = reflect.New(it).Elem()
					} else if arg.Kind() != it.Kind() && arg.Type().ConvertibleTo(f.Type().In(i)) {
						arg = arg.Convert(it)
					}
				}
			}

			if !isReflect {
				args = append(args, arg)
			} else {
				if arg.Kind() == reflect.Interface {
					arg = arg.Elem()
				}
				args = append(args, reflect.ValueOf(arg))
			}
		}
		ret := NilValue
		var err error
		func() {
			defer func() {
				if os.Getenv("ANKO_DEBUG") == "" {
					if ex := recover(); ex != nil {
						err = errors.New(fmt.Sprint(ex))
					}
				}
			}()
			if f.Kind() == reflect.Interface {
				f = f.Elem()
			}
			rets := f.Call(args)
			if isReflect {
				ev := rets[1].Interface()
				if ev != nil {
					err = ev.(error)
				}
				ret = rets[0].Interface().(reflect.Value)
			} else {
				if f.Type().NumOut() == 1 {
					ret = rets[0]
				} else {
					var result []interface{}
					for _, r := range rets {
						result = append(result, r.Interface())
					}
					ret = reflect.ValueOf(result)
				}
			}
		}()
		if err != nil {
			return ret, NewError(expr, err)
		}
		return ret, nil
	case *ast.TernaryOpExpr:
		rv, err := invokeExpr(e.Expr, env)
		if err != nil {
			return rv, NewError(expr, err)
		}
		if toBool(rv) {
			lhsV, err := invokeExpr(e.Lhs, env)
			if err != nil {
				return lhsV, NewError(expr, err)
			}
			return lhsV, nil
		}
		rhsV, err := invokeExpr(e.Rhs, env)
		if err != nil {
			return rhsV, NewError(expr, err)
		}
		return rhsV, nil
	default:
		return NilValue, NewStringError(expr, "Unknown expression")
	}
}
