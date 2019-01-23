package vm

import (
	"context"
	"math"
	"reflect"
	"strings"

	"github.com/mattn/anko/ast"
)

// invokeOperator evaluates one Operator.
func invokeOperator(ctx context.Context, operator ast.Operator, env *Env) (reflect.Value, error) {
	switch op := operator.(type) {

	// BinaryOperator
	case *ast.BinaryOperator:
		rv, err := invokeExpr(ctx, op.LHS, env)
		if err != nil {
			return nilValue, newError(op.LHS, err)
		}
		if rv.Kind() == reflect.Interface && !rv.IsNil() {
			rv = rv.Elem()
		}

		switch op.Operator {
		case "||":
			if toBool(rv) {
				return rv, nil
			}
		case "&&":
			if !toBool(rv) {
				return rv, nil
			}
		default:
			return nilValue, newStringError(op, "unknown operator")
		}

		if op.RHS == nil {
			return nilValue, nil
		}

		rv, err = invokeExpr(ctx, op.RHS, env)
		if err != nil {
			return nilValue, newError(op.RHS, err)
		}
		if rv.Kind() == reflect.Interface && !rv.IsNil() {
			rv = rv.Elem()
		}
		return rv, nil

	// ComparisonOperator
	case *ast.ComparisonOperator:
		lhsV, err := invokeExpr(ctx, op.LHS, env)
		if err != nil {
			return nilValue, newError(op.LHS, err)
		}
		if lhsV.Kind() == reflect.Interface && !lhsV.IsNil() {
			lhsV = lhsV.Elem()
		}

		var rhsV reflect.Value
		rhsV, err = invokeExpr(ctx, op.RHS, env)
		if err != nil {
			return nilValue, newError(op.RHS, err)
		}
		if rhsV.Kind() == reflect.Interface && !rhsV.IsNil() {
			rhsV = rhsV.Elem()
		}

		switch op.Operator {
		case "==":
			return reflect.ValueOf(equal(lhsV, rhsV)), nil
		case "!=":
			return reflect.ValueOf(equal(lhsV, rhsV) == false), nil
		case "<":
			return reflect.ValueOf(toFloat64(lhsV) < toFloat64(rhsV)), nil
		case "<=":
			return reflect.ValueOf(toFloat64(lhsV) <= toFloat64(rhsV)), nil
		case ">":
			return reflect.ValueOf(toFloat64(lhsV) > toFloat64(rhsV)), nil
		case ">=":
			return reflect.ValueOf(toFloat64(lhsV) >= toFloat64(rhsV)), nil
		}

		return nilValue, newStringError(op, "unknown operator")

	// AddOperator
	case *ast.AddOperator:
		lhsV, err := invokeExpr(ctx, op.LHS, env)
		if err != nil {
			return nilValue, newError(op.LHS, err)
		}
		if lhsV.Kind() == reflect.Interface && !lhsV.IsNil() {
			lhsV = lhsV.Elem()
		}

		var rhsV reflect.Value
		rhsV, err = invokeExpr(ctx, op.RHS, env)
		if err != nil {
			return nilValue, newError(op.RHS, err)
		}
		if rhsV.Kind() == reflect.Interface && !rhsV.IsNil() {
			rhsV = rhsV.Elem()
		}

		switch op.Operator {
		case "+":
			if (lhsV.Kind() == reflect.Slice || lhsV.Kind() == reflect.Array) && (rhsV.Kind() != reflect.Slice && rhsV.Kind() != reflect.Array) {
				rhsT := rhsV.Type()
				lhsT := lhsV.Type().Elem()
				if lhsT.Kind() != rhsT.Kind() {
					if !rhsT.ConvertibleTo(lhsT) {
						return nilValue, newStringError(op, "invalid type conversion")
					}
					rhsV = rhsV.Convert(lhsT)
				}
				return reflect.Append(lhsV, rhsV), nil
			}
			if (lhsV.Kind() == reflect.Slice || lhsV.Kind() == reflect.Array) && (rhsV.Kind() == reflect.Slice || rhsV.Kind() == reflect.Array) {
				return appendSlice(op, lhsV, rhsV)
			}
			if lhsV.Kind() == reflect.String || rhsV.Kind() == reflect.String {
				return reflect.ValueOf(toString(lhsV) + toString(rhsV)), nil
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
		case "|":
			return reflect.ValueOf(toInt64(lhsV) | toInt64(rhsV)), nil
		}

		return nilValue, newStringError(op, "unknown operator")

	// MultiplyOperator
	case *ast.MultiplyOperator:
		lhsV, err := invokeExpr(ctx, op.LHS, env)
		if err != nil {
			return nilValue, newError(op.LHS, err)
		}
		if lhsV.Kind() == reflect.Interface && !lhsV.IsNil() {
			lhsV = lhsV.Elem()
		}

		var rhsV reflect.Value
		rhsV, err = invokeExpr(ctx, op.RHS, env)
		if err != nil {
			return nilValue, newError(op.RHS, err)
		}
		if rhsV.Kind() == reflect.Interface && !rhsV.IsNil() {
			rhsV = rhsV.Elem()
		}

		switch op.Operator {
		case "*":
			if lhsV.Kind() == reflect.String && (rhsV.Kind() == reflect.Int || rhsV.Kind() == reflect.Int32 || rhsV.Kind() == reflect.Int64) {
				return reflect.ValueOf(strings.Repeat(toString(lhsV), int(toInt64(rhsV)))), nil
			}
			if lhsV.Kind() == reflect.Float64 || rhsV.Kind() == reflect.Float64 {
				return reflect.ValueOf(toFloat64(lhsV) * toFloat64(rhsV)), nil
			}
			return reflect.ValueOf(toInt64(lhsV) * toInt64(rhsV)), nil
		case "/":
			return reflect.ValueOf(toFloat64(lhsV) / toFloat64(rhsV)), nil
		case "%":
			return reflect.ValueOf(toInt64(lhsV) % toInt64(rhsV)), nil
		case ">>":
			return reflect.ValueOf(toInt64(lhsV) >> uint64(toInt64(rhsV))), nil
		case "<<":
			return reflect.ValueOf(toInt64(lhsV) << uint64(toInt64(rhsV))), nil
		case "&":
			return reflect.ValueOf(toInt64(lhsV) & toInt64(rhsV)), nil
		case "**":
			if lhsV.Kind() == reflect.Float64 {
				return reflect.ValueOf(math.Pow(lhsV.Float(), toFloat64(rhsV))), nil
			}
			return reflect.ValueOf(int64(math.Pow(toFloat64(lhsV), toFloat64(rhsV)))), nil
		}

		return nilValue, newStringError(op, "unknown operator")

	default:
		return nilValue, newStringError(op, "unknown operator")

	}
}
