package vm_test

import (
	"log"

	"github.com/mattn/anko/packages"
	"github.com/mattn/anko/vm"
)

func Example_vmRegexp() {
	env := vm.NewEnv()
	packages.DefineImport(env)

	script := `
fmt = import("fmt")
regexp = import("regexp")

re = regexp.MustCompile("^simple$")
result = re.MatchString("simple")
fmt.Println(result)
fmt.Println("")

re = regexp.MustCompile("simple")
result = re.FindString("This is a simple sentence")
fmt.Println(result)
fmt.Println("")

re = regexp.MustCompile(",")
result = re.Split("a,b,c", -1)
fmt.Println(result)
fmt.Println("")

re = regexp.MustCompile("foo")
result = re.ReplaceAllString("foo", "bar")
fmt.Println(result)
`

	_, err := env.Execute(script)
	if err != nil {
		log.Fatalf("execute error: %v\n", err)
	}

	// output:
	// true
	//
	// simple
	//
	// [a b c]
	//
	// bar
}
