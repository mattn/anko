package vm

import (
	"reflect"
	"strings"

	"github.com/mattn/anko/ast"
)

// invokeOperator evaluates one Operator.
func (runInfo *runInfoStruct) invokeOperator() {
	switch operator := runInfo.operator.(type) {

	// BinaryOperator
	case *ast.BinaryOperator:
		runInfo.expr = operator.LHS
		runInfo.invokeExpr()
		if runInfo.err != nil {
			return
		}
		if runInfo.rv.Kind() == reflect.Interface && !runInfo.rv.IsNil() {
			runInfo.rv = runInfo.rv.Elem()
		}

		switch operator.Operator {
		case "||":
			if toBool(runInfo.rv) {
				runInfo.rv = trueValue
				return
			}
		case "&&":
			if !toBool(runInfo.rv) {
				runInfo.rv = falseValue
				return
			}
		default:
			runInfo.err = newStringError(operator, "unknown operator")
			runInfo.rv = nilValue
			return
		}

		runInfo.expr = operator.RHS
		runInfo.invokeExpr()
		if runInfo.err != nil {
			return
		}
		if runInfo.rv.Kind() == reflect.Interface && !runInfo.rv.IsNil() {
			runInfo.rv = runInfo.rv.Elem()
		}

		if toBool(runInfo.rv) {
			runInfo.rv = trueValue
		} else {
			runInfo.rv = falseValue
		}

	// ComparisonOperator
	case *ast.ComparisonOperator:
		runInfo.expr = operator.LHS
		runInfo.invokeExpr()
		if runInfo.err != nil {
			return
		}
		if runInfo.rv.Kind() == reflect.Interface && !runInfo.rv.IsNil() {
			runInfo.rv = runInfo.rv.Elem()
		}
		lhsV := runInfo.rv

		runInfo.expr = operator.RHS
		runInfo.invokeExpr()
		if runInfo.err != nil {
			return
		}
		if runInfo.rv.Kind() == reflect.Interface && !runInfo.rv.IsNil() {
			runInfo.rv = runInfo.rv.Elem()
		}

		switch operator.Operator {
		case "==":
			if equal(lhsV, runInfo.rv) {
				runInfo.rv = trueValue
			} else {
				runInfo.rv = falseValue
			}
		case "!=":
			if !equal(lhsV, runInfo.rv) {
				runInfo.rv = trueValue
			} else {
				runInfo.rv = falseValue
			}
		case "<":
			// integers are compared as integers to keep full int64 precision
			var result bool
			if isIntKind(lhsV) && isIntKind(runInfo.rv) {
				result = lhsV.Int() < runInfo.rv.Int()
			} else {
				result = toFloat64(lhsV) < toFloat64(runInfo.rv)
			}
			if result {
				runInfo.rv = trueValue
			} else {
				runInfo.rv = falseValue
			}
		case "<=":
			var result bool
			if isIntKind(lhsV) && isIntKind(runInfo.rv) {
				result = lhsV.Int() <= runInfo.rv.Int()
			} else {
				result = toFloat64(lhsV) <= toFloat64(runInfo.rv)
			}
			if result {
				runInfo.rv = trueValue
			} else {
				runInfo.rv = falseValue
			}
		case ">":
			var result bool
			if isIntKind(lhsV) && isIntKind(runInfo.rv) {
				result = lhsV.Int() > runInfo.rv.Int()
			} else {
				result = toFloat64(lhsV) > toFloat64(runInfo.rv)
			}
			if result {
				runInfo.rv = trueValue
			} else {
				runInfo.rv = falseValue
			}
		case ">=":
			var result bool
			if isIntKind(lhsV) && isIntKind(runInfo.rv) {
				result = lhsV.Int() >= runInfo.rv.Int()
			} else {
				result = toFloat64(lhsV) >= toFloat64(runInfo.rv)
			}
			if result {
				runInfo.rv = trueValue
			} else {
				runInfo.rv = falseValue
			}
		default:
			runInfo.err = newStringError(operator, "unknown operator")
			runInfo.rv = nilValue
		}

	// AddOperator
	case *ast.AddOperator:
		runInfo.expr = operator.LHS
		runInfo.invokeExpr()
		if runInfo.err != nil {
			return
		}
		if runInfo.rv.Kind() == reflect.Interface && !runInfo.rv.IsNil() {
			runInfo.rv = runInfo.rv.Elem()
		}
		lhsV := runInfo.rv

		runInfo.expr = operator.RHS
		runInfo.invokeExpr()
		if runInfo.err != nil {
			return
		}
		if runInfo.rv.Kind() == reflect.Interface && !runInfo.rv.IsNil() {
			runInfo.rv = runInfo.rv.Elem()
		}

		switch operator.Operator {
		case "+":
			lhsKind := lhsV.Kind()
			rhsKind := runInfo.rv.Kind()

			if lhsKind == reflect.Slice || lhsKind == reflect.Array {
				if rhsKind == reflect.Slice || rhsKind == reflect.Array {
					// append slice to slice
					runInfo.rv, runInfo.err = appendSlice(operator, lhsV, runInfo.rv)
					return
				}
				// try to append rhs non-slice to lhs slice
				runInfo.rv, runInfo.err = convertReflectValueToType(runInfo.rv, lhsV.Type().Elem())
				if runInfo.err != nil {
					runInfo.err = newStringError(operator, "invalid type conversion")
					runInfo.rv = nilValue
					return
				}
				runInfo.rv = reflect.Append(lhsV, runInfo.rv)
				return
			}
			if rhsKind == reflect.Slice || rhsKind == reflect.Array {
				// can not append rhs slice to lhs non-slice
				runInfo.err = newStringError(operator, "invalid type conversion")
				runInfo.rv = nilValue
				return
			}

			kind := precedenceOfKinds(lhsKind, rhsKind)
			switch kind {
			case reflect.String:
				runInfo.rv = reflect.ValueOf(toString(lhsV) + toString(runInfo.rv))
			case reflect.Float64, reflect.Float32:
				runInfo.rv = float64Value(toFloat64(lhsV) + toFloat64(runInfo.rv))
			default:
				runInfo.rv = int64Value(toInt64(lhsV) + toInt64(runInfo.rv))
			}

		case "-":
			switch lhsV.Kind() {
			case reflect.Float64, reflect.Float32:
				runInfo.rv = float64Value(toFloat64(lhsV) - toFloat64(runInfo.rv))
				return
			}
			switch runInfo.rv.Kind() {
			case reflect.Float64, reflect.Float32:
				runInfo.rv = float64Value(toFloat64(lhsV) - toFloat64(runInfo.rv))
			default:
				runInfo.rv = int64Value(toInt64(lhsV) - toInt64(runInfo.rv))
			}

		case "|":
			runInfo.rv = int64Value(toInt64(lhsV) | toInt64(runInfo.rv))
		default:
			runInfo.err = newStringError(operator, "unknown operator")
			runInfo.rv = nilValue
		}

	// MultiplyOperator
	case *ast.MultiplyOperator:
		runInfo.expr = operator.LHS
		runInfo.invokeExpr()
		if runInfo.err != nil {
			return
		}
		if runInfo.rv.Kind() == reflect.Interface && !runInfo.rv.IsNil() {
			runInfo.rv = runInfo.rv.Elem()
		}
		lhsV := runInfo.rv

		runInfo.expr = operator.RHS
		runInfo.invokeExpr()
		if runInfo.err != nil {
			return
		}
		if runInfo.rv.Kind() == reflect.Interface && !runInfo.rv.IsNil() {
			runInfo.rv = runInfo.rv.Elem()
		}

		switch operator.Operator {
		case "*":
			if lhsV.Kind() == reflect.String && (runInfo.rv.Kind() == reflect.Int || runInfo.rv.Kind() == reflect.Int32 || runInfo.rv.Kind() == reflect.Int64) {
				count := toInt64(runInfo.rv)
				if count < 0 {
					runInfo.err = newStringError(operator, "negative repeat count")
					runInfo.rv = nilValue
					return
				}
				runInfo.rv = reflect.ValueOf(strings.Repeat(toString(lhsV), int(count)))
				return
			}
			if lhsV.Kind() == reflect.Float64 || runInfo.rv.Kind() == reflect.Float64 {
				runInfo.rv = float64Value(toFloat64(lhsV) * toFloat64(runInfo.rv))
				return
			}
			runInfo.rv = int64Value(toInt64(lhsV) * toInt64(runInfo.rv))
		case "/":
			runInfo.rv = float64Value(toFloat64(lhsV) / toFloat64(runInfo.rv))
		case "%":
			rhs := toInt64(runInfo.rv)
			if rhs == 0 {
				runInfo.err = newStringError(operator, "integer divide by zero")
				runInfo.rv = nilValue
				return
			}
			runInfo.rv = int64Value(toInt64(lhsV) % rhs)
		case ">>":
			runInfo.rv = int64Value(toInt64(lhsV) >> uint64(toInt64(runInfo.rv)))
		case "<<":
			runInfo.rv = int64Value(toInt64(lhsV) << uint64(toInt64(runInfo.rv)))
		case "&":
			runInfo.rv = int64Value(toInt64(lhsV) & toInt64(runInfo.rv))

		default:
			runInfo.err = newStringError(operator, "unknown operator")
			runInfo.rv = nilValue
		}

	default:
		runInfo.err = newStringError(operator, "unknown operator")
		runInfo.rv = nilValue

	}
}
