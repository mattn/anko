package vm

import (
	"context"
	"reflect"
	"testing"

	"github.com/mattn/anko/env"
	"github.com/mattn/anko/parser"
)

func benchmarkExecute(b *testing.B, script string) {
	b.Helper()
	stmt, err := parser.ParseSrc(script)
	if err != nil {
		b.Fatal(err)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e := env.NewEnv()
		_, err := RunContext(context.Background(), e, nil, stmt)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkArithmeticInt(b *testing.B) {
	benchmarkExecute(b, `x = 1 + 2 + 3 + 4 + 5`)
}

func BenchmarkArithmeticFloat(b *testing.B) {
	benchmarkExecute(b, `x = 1.1 + 2.2 + 3.3 + 4.4 + 5.5`)
}

func BenchmarkComparison(b *testing.B) {
	benchmarkExecute(b, `x = 1 == 1; y = 2 != 3; z = 4 < 5`)
}

func BenchmarkLoop(b *testing.B) {
	benchmarkExecute(b, `x = 0; for i = 0; i < 100; i++ { x += i }`)
}

func BenchmarkFibonacci(b *testing.B) {
	benchmarkExecute(b, `
func fib(n) {
	if n <= 1 { return n }
	return fib(n-1) + fib(n-2)
}
fib(15)
`)
}

func BenchmarkStringConcat(b *testing.B) {
	benchmarkExecute(b, `x = "a" + "b" + "c" + "d" + "e"`)
}

func BenchmarkVarAccess(b *testing.B) {
	benchmarkExecute(b, `x = 1; y = x; z = y; w = z`)
}

func BenchmarkMapAccess(b *testing.B) {
	benchmarkExecute(b, `x = {"a": 1, "b": 2, "c": 3}; y = x["a"] + x["b"] + x["c"]`)
}

func BenchmarkArrayAccess(b *testing.B) {
	benchmarkExecute(b, `x = [1, 2, 3, 4, 5]; y = x[0] + x[1] + x[2] + x[3] + x[4]`)
}

func BenchmarkFuncCall(b *testing.B) {
	benchmarkExecute(b, `func add(a, b) { return a + b }; x = add(1, 2)`)
}

func BenchmarkEqual(b *testing.B) {
	b.Run("IntInt", func(b *testing.B) {
		lhs := int64(42)
		rhs := int64(42)
		lhsV := reflect.ValueOf(lhs)
		rhsV := reflect.ValueOf(rhs)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			equal(lhsV, rhsV)
		}
	})
	b.Run("FloatFloat", func(b *testing.B) {
		lhs := float64(3.14)
		rhs := float64(3.14)
		lhsV := reflect.ValueOf(lhs)
		rhsV := reflect.ValueOf(rhs)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			equal(lhsV, rhsV)
		}
	})
	b.Run("IntFloat", func(b *testing.B) {
		lhs := int64(42)
		rhs := float64(42.0)
		lhsV := reflect.ValueOf(lhs)
		rhsV := reflect.ValueOf(rhs)
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			equal(lhsV, rhsV)
		}
	})
}
