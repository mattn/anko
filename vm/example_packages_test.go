package vm_test

import (
	"log"

	"github.com/mattn/anko/env"
	_ "github.com/mattn/anko/packages"
	"github.com/mattn/anko/vm"
)

func Example_vmSort() {
	// _ "github.com/mattn/anko/packages"

	e := env.NewEnv()

	script := `
fmt = import("fmt")
sort = import("sort")
a = [5, 1.1, 3, "f", "2", "4.4"]
sortFuncs = make(sort.SortFuncsStruct)
sortFuncs.LenFunc = func() { return len(a) }
sortFuncs.LessFunc = func(i, j) { return a[i] < a[j] }
sortFuncs.SwapFunc = func(i, j) { temp = a[i]; a[i] = a[j]; a[j] = temp }
sort.Sort(sortFuncs)
fmt.Println(a)
`

	_, err := vm.Execute(e, nil, script)
	if err != nil {
		log.Fatalf("execute error: %v\n", err)
	}

	// output:
	// [f 1.1 2 3 4.4 5]
}

func Example_vmRegexp() {
	// _ "github.com/mattn/anko/packages"

	e := env.NewEnv()

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

	_, err := vm.Execute(e, nil, script)
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

func Example_vmHttp() {
	// _ "github.com/mattn/anko/packages"

	e := env.NewEnv()

	script := `
fmt = import("fmt")
io = import("io")
ioutil = import("io/ioutil")
net = import("net")
http = import("net/http")
time = import("time")

func handlerRoot(responseWriter, request) {
	io.WriteString(responseWriter, "Hello World :)")
}

serveMux = http.NewServeMux()
serveMux.HandleFunc("/", handlerRoot)
listener, err = net.Listen("tcp", ":8080")
if err != nil {
	fmt.Println(err)
	return
}
go http.Serve(listener, serveMux)

client = http.DefaultClient

response, err = client.Get("http://localhost:8080/")
if err != nil {
	fmt.Println(err)
	return
}

body, err = ioutil.ReadAll(response.Body)
if err != nil {
	fmt.Println(err)
}
response.Body.Close()

fmt.Printf("%s\n", body)
`

	_, err := vm.Execute(e, nil, script)
	if err != nil {
		log.Fatalf("execute error: %v\n", err)
	}

	// output:
	// Hello World :)
}
