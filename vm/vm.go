package vm

import (
	"errors"
	"fmt"
	"github.com/mattn/anko/ast"
	"reflect"
	"strconv"
	"strings"
)

var Nil *interface{}
var NilValue = reflect.ValueOf(&Nil)
var TrueValue = reflect.ValueOf(true)
var FalseValue = reflect.ValueOf(false)

// Error provides a convenient interface for handling runtime error.
// It can be Error interface with type cast which can call Pos().
type Error struct {
	message string
	pos     ast.Position
}

var BreakError = errors.New("break")
var ContinueError = errors.New("continue")

// NewErrorString makes error interface with message.
func NewErrorString(err string, pos ast.Pos) error {
	return &Error{message: err, pos: pos.GetPos()}
}

// NewError makes error interface with message. This doesn't overwrite last error.
func NewError(err error, pos ast.Pos) error {
	if err == nil {
		return nil
	}
	if err == BreakError || err == ContinueError {
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
		if err != nil {
			return NilValue, NewError(err, stmt)
		}
		if _, ok := stmt.(*ast.ReturnStmt); ok {
			return rv, nil
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
			return NilValue, NewError(err, stmt)
		}
		return rv, nil
	case *ast.VarStmt:
		rv := NilValue
		var err error
		rvs := []reflect.Value{}
		for _, expr := range stmt.Exprs {
			rv, err = invokeExpr(expr, env)
			if err != nil {
				return NilValue, NewError(err, expr)
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
			return NilValue, NewError(err, stmt)
		}
		if toBool(rv) {
			// Then
			rv, err = Run(stmt.Then, env.NewEnv())
			if err != nil {
				return NilValue, NewError(err, stmt)
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
					return NilValue, NewError(err, stmt)
				}
				if !toBool(rv) {
					continue
				}
				// ElseIf Then
				done = true
				rv, err = Run(stmt_if.Then, env)
				if err != nil {
					return NilValue, NewError(err, stmt)
				}
				break
			}
		}
		if !done && len(stmt.Else) > 0 {
			// Else
			rv, err = Run(stmt.Else, env.NewEnv())
			if err != nil {
				return NilValue, NewError(err, stmt)
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
				err = NewError(e1, stmt.Catch[0])
			} else {
				err = nil
			}
		}
		if len(stmt.Finally) > 0 {
			_, e2 := Run(stmt.Finally, env.NewEnv())
			if e2 != nil {
				err = NewError(e2, stmt.Finally[0])
			}
		}
		return NilValue, NewError(err, stmt)
	case *ast.ForStmt:
		val, ee := invokeExpr(stmt.Value, env)
		if ee != nil {
			return NilValue, ee
		}
		if val.Kind() != reflect.Array && val.Kind() != reflect.Slice {
			return NilValue, NewErrorString("Invalid operation", stmt)
		}
		newenv := env.NewEnv()
		for i := 0; i < val.Len(); i++ {
			newenv.Define(stmt.Var, val.Index(i))
			_, err := Run(stmt.Stmts, newenv)
			if err != nil {
				if err == BreakError {
					err = nil
					break
				}
				if err == ContinueError {
					err = nil
					continue
				}
				return NilValue, NewError(err, stmt)
			}
		}
		return NilValue, nil
	case *ast.ReturnStmt:
		rv, err := invokeExpr(stmt.Expr, env)
		if err != nil {
			return NilValue, NewError(err, stmt)
		}
		return rv, nil
	case *ast.ThrowStmt:
		rv, err := invokeExpr(stmt.Expr, env)
		if err != nil {
			return NilValue, NewError(err, stmt)
		}
		return NilValue, NewErrorString(fmt.Sprint(rv.Interface()), stmt)
	case *ast.ModuleStmt:
		newenv := env.NewEnv()
		newenv.SetName(stmt.Name)
		rv, err := Run(stmt.Stmts, newenv)
		if err != nil {
			return NilValue, NewError(err, stmt)
		}
		env.DefineGlobal(stmt.Name, reflect.ValueOf(newenv))
		return rv, nil
	default:
		return NilValue, NewErrorString("Unknown statement", stmt)
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

// equal return true when lhsV and rhsV is same value.
func equal(lhsV, rhsV reflect.Value) bool {
	if lhsV.Kind() == reflect.Interface {
		lhsV = lhsV.Elem()
	}
	if rhsV.Kind() == reflect.Interface {
		rhsV = rhsV.Elem()
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

// invokeExpr evaluate one expression.
func invokeExpr(expr ast.Expr, env *Env) (reflect.Value, error) {
	switch e := expr.(type) {
	case *ast.NumberExpr:
		if strings.Contains(e.Lit, ".") {
			v, err := strconv.ParseFloat(e.Lit, 64)
			if err != nil {
				return NilValue, NewError(err, expr)
			}
			return reflect.ValueOf(float64(v)), nil
		}
		v, err := strconv.Atoi(e.Lit)
		if err != nil {
			return NilValue, NewError(err, expr)
		}
		return reflect.ValueOf(int64(v)), nil
	case *ast.IdentExpr:
		if v, ok := env.Get(e.Lit); ok {
			return v, nil
		} else {
			return NilValue, NewErrorString(fmt.Sprintf("Undefined variable: %s", e.Lit), expr)
		}
	case *ast.StringExpr:
		return reflect.ValueOf(e.Lit), nil
	case *ast.ArrayExpr:
		a := make([]interface{}, len(e.Exprs))
		for i, expr := range e.Exprs {
			arg, err := invokeExpr(expr, env)
			if err != nil {
				return NilValue, NewError(err, expr)
			}
			a[i] = arg.Interface()
		}
		return reflect.ValueOf(a), nil
	case *ast.MapExpr:
		m := make(map[string]interface{})
		for k, expr := range e.MapExpr {
			v, err := invokeExpr(expr, env)
			if err != nil {
				return NilValue, NewError(err, expr)
			}
			m[k] = v.Interface()
		}
		return reflect.ValueOf(m), nil
	case *ast.UnaryMinusExpr:
		v, err := invokeExpr(e.SubExpr, env)
		if err != nil {
			return NilValue, NewError(err, expr)
		}
		if v.Kind() == reflect.Float64 {
			return reflect.ValueOf(-v.Float()), nil
		}
		return reflect.ValueOf(-v.Int()), nil
	case *ast.ParenExpr:
		v, err := invokeExpr(e.SubExpr, env)
		if err != nil {
			return NilValue, NewError(err, expr)
		}
		return v, nil
	case *ast.FuncExpr:
		f := reflect.ValueOf(func(expr *ast.FuncExpr, env *Env) Func {
			return func(args ...reflect.Value) (reflect.Value, error) {
				if !expr.VarArg {
					if len(args) != len(expr.Args) {
						return NilValue, NewErrorString("Arguments Number of mismatch", expr)
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
				return Run(expr.Stmts, newenv)
			}
		}(e, env))
		env.Define(e.Name, f)
		return f, nil
	case *ast.ItemExpr:
		v, err := invokeExpr(e.Value, env)
		if err != nil {
			return NilValue, NewError(err, expr)
		}
		i, err := invokeExpr(e.Index, env)
		if err != nil {
			return NilValue, NewError(err, expr)
		}
		if v.Kind() == reflect.Interface {
			v = v.Elem()
		}
		if v.Kind() == reflect.Array || v.Kind() == reflect.Slice {
			if i.Kind() != reflect.Int && i.Kind() != reflect.Int64 {
				return NilValue, NewErrorString("Array index should be int", expr)
			}
			ii := int(i.Int())
			if ii < 0 || ii >= v.Len() {
				return NilValue, nil
			}
			return v.Index(ii), nil
		}
		if v.Kind() == reflect.Map {
			if i.Kind() != reflect.String {
				return NilValue, NewErrorString("Map key should be string", expr)
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
		return NilValue, NewErrorString("Invalid operation", expr)
	case *ast.MemberExpr:
		v, err := invokeExpr(e.Expr, env)
		if err != nil {
			return NilValue, NewError(err, expr)
		}
		if v.Kind() == reflect.Interface {
			v = v.Elem()
		}
		if v.Kind() == reflect.Slice {
			v = v.Index(0)
		}
		if v.CanInterface() {
			if vme, ok := v.Interface().(*Env); ok {
				m, ok := vme.Get(e.Method)
				if !m.IsValid() || !ok {
					return NilValue, NewErrorString("Invalid operation", expr)
				}
				return m, nil
			}
		}

		m := v.MethodByName(e.Method)
		if !m.IsValid() {
			if v.Kind() == reflect.Ptr {
				v = v.Elem()
			}
			if v.Kind() == reflect.Struct {
				m = v.FieldByName(e.Method)
				if !m.IsValid() {
					return NilValue, NewErrorString("Invalid operation", expr)
				}
			} else if v.Kind() == reflect.Map {
				m = v.MapIndex(reflect.ValueOf(e.Method))
				if !m.IsValid() {
					return NilValue, NewErrorString("Invalid operation", expr)
				}
			} else {
				return NilValue, NewErrorString("Invalid operation", expr)
			}
		}
		return m, nil
	case *ast.LetExpr:
		rv := NilValue
		var err error
		vs := []interface{}{}
		for i, ee := range e.Exprs {
			if e.Operator == "=" {
				rv, err = invokeExpr(ee, env)
			} else {
				rv, err = invokeExpr(&ast.BinOpExpr{Lhs: &ast.IdentExpr{Lit: e.Names[i]}, Operator: e.Operator, Rhs: ee}, env)
			}
			if err != nil {
				return NilValue, NewError(err, ee)
			}
			if rv == NilValue {
				vs = append(vs, nil)
			} else {
				vs = append(vs, rv.Interface())
			}
		}
		rvs := reflect.ValueOf(vs)
		if len(e.Names) > 1 && rvs.Len() == 1 {
			item := rvs.Index(0)
			if item.Kind() == reflect.Interface {
				item = item.Elem()
			}
			if item.Kind() == reflect.Slice {
				rvs = item
			}
		}
		for i, name := range e.Names {
			if i >= rvs.Len() {
				continue
			}
			v := rvs.Index(i)
			if v.Kind() == reflect.Interface {
				v = v.Elem()
			}
			if env.Set(name, v) != nil {
				env.Define(name, v)
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
		lhsV, err := invokeExpr(e.Lhs, env)
		if err != nil {
			return NilValue, NewError(err, expr)
		}
		rhsV, err := invokeExpr(e.Rhs, env)
		if err != nil {
			return NilValue, NewError(err, expr)
		}
		if lhsV.Kind() == reflect.Interface {
			lhsV = lhsV.Elem()
		}
		if rhsV.Kind() == reflect.Interface {
			rhsV = rhsV.Elem()
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
			return reflect.ValueOf(!equal(lhsV, rhsV)), nil
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
			if toBool(rhsV) {
				return rhsV, nil
			}
			return FalseValue, nil
		case "&":
			return reflect.ValueOf(toInt64(lhsV) & toInt64(rhsV)), nil
		case "&&":
			if toBool(lhsV) {
				return rhsV, nil
			}
			return lhsV, nil
		default:
			return NilValue, NewErrorString("Unknown operator", expr)
		}
	case *ast.ConstExpr:
		return reflect.ValueOf(e.Value), nil
	case *ast.AnonCallExpr:
		f, err := invokeExpr(e.Expr, env)
		if err != nil {
			return NilValue, NewError(err, expr)
		}
		if f.Kind() != reflect.Func {
			return NilValue, NewErrorString("Unknown function", expr)
		}
		return invokeExpr(&ast.CallExpr{Func: f, SubExprs: e.SubExprs}, env)
	case *ast.CallExpr:
		f := NilValue
		if e.Func != nil {
			f = e.Func.(reflect.Value)
		} else {
			var ok bool
			f, ok = env.Get(e.Name)
			if !ok {
				return NilValue, NewError(fmt.Errorf("Undefined function '%s'", e.Name), expr)
			}
		}
		_, isReflect := f.Interface().(Func)

		args := []reflect.Value{}
		for i, expr := range e.SubExprs {
			arg, err := invokeExpr(expr, env)
			if err != nil {
				return NilValue, NewError(err, expr)
			}
			if (arg.Kind() == arg.Interface() || arg.Kind() == reflect.Ptr) && arg.IsNil() && !f.Type().IsVariadic() {
				arg = reflect.New(f.Type().In(i)).Elem()
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
				if ex := recover(); ex != nil {
					err = errors.New(fmt.Sprint(ex))
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
			return NilValue, NewError(err, expr)
		}
		return ret, nil
	case *ast.TernaryOpExpr:
		rv, err := invokeExpr(e.Expr, env)
		if err != nil {
			return NilValue, NewError(err, expr)
		}
		if toBool(rv) {
			lhsV, err := invokeExpr(e.Lhs, env)
			if err != nil {
				return NilValue, NewError(err, expr)
			}
			return lhsV, nil
		}
		rhsV, err := invokeExpr(e.Rhs, env)
		if err != nil {
			return NilValue, NewError(err, expr)
		}
		return rhsV, nil
	default:
		return NilValue, NewErrorString("Unknown expression", expr)
	}
}
