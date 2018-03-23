package core

import (
	"os"
	"testing"
)

func TestToX(t *testing.T) {
	os.Setenv("ANKO_DEBUG", "1")
	tests := []testStruct{
		{script: `toBool(-2)`, runOutput: false},
		{script: `toBool(-1.5)`, runOutput: false},
		{script: `toBool(-1)`, runOutput: false},
		{script: `toBool(-0.4)`, runOutput: false},
		{script: `toBool(0)`, runOutput: false},
		{script: `toBool(0.4)`, runOutput: true},
		{script: `toBool(1)`, runOutput: true},
		{script: `toBool(1.5)`, runOutput: true},
		{script: `toBool(2)`, runOutput: true},
		{script: `toBool("true")`, runOutput: true},
		{script: `toBool("false")`, runOutput: false},
		{script: `toBool("yes")`, runOutput: true},
		{script: `toBool("ye")`, runOutput: false},
		{script: `toBool("y")`, runOutput: true},
		{script: `toBool("false")`, runOutput: false},
		{script: `toBool("f")`, runOutput: false},
		{script: `toBool("")`, runOutput: false},
		{script: `toBool(nil)`, runOutput: false},
		{script: `toBool({})`, runOutput: false},
		{script: `toBool([])`, runOutput: false},
		{script: `toBool([true])`, runOutput: false},
		{script: `toBool({"true": "true"})`, runOutput: false},
		{script: `toString(nil)`, runOutput: "<nil>"},
		{script: `toString("")`, runOutput: ""},
		{script: `toString(1)`, runOutput: "1"},
		{script: `toString(1.2)`, runOutput: "1.2"},
		{script: `toString(1/3)`, runOutput: "0.3333333333333333"},
		{script: `toString(false)`, runOutput: "false"},
		{script: `toString(true)`, runOutput: "true"},
		{script: `toString({})`, runOutput: "map[]"},
		{script: `toString({"foo": "bar"})`, runOutput: "map[foo:bar]"},
		{script: `toString([true,nil])`, runOutput: "[true <nil>]"},
	}
	runTests(t, tests)
}
