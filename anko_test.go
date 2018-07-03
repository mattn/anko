// +build !appengine

package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"
)

var logger *log.Logger

func TestMain(m *testing.M) {
	parseFlags()
	code := m.Run()
	os.Exit(code)
}

func TestRunNonInteractiveFile(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	testDir := filepath.Join(filepath.Dir(filename), "core", "testdata")
	setupEnv()

	file = filepath.Join(testDir, "not-found.ank")
	exitCode := runNonInteractive()
	if exitCode != 2 {
		t.Fatalf("exitCode - received: %v - expected: %v", exitCode, 2)
	}

	file = filepath.Join(testDir, "broken.ank")
	exitCode = runNonInteractive()
	os.Args = []string{os.Args[0]}
	if exitCode != 4 {
		t.Fatalf("exitCode - received: %v - expected: %v", exitCode, 4)
	}

	file = filepath.Join(testDir, "test.ank")
	exitCode = runNonInteractive()
	os.Args = []string{os.Args[0]}
	if exitCode != 0 {
		t.Fatalf("exitCode - received: %v - expected: %v", exitCode, 0)
	}

	file = ""
}

func TestRunNonInteractiveExecute(t *testing.T) {
	setupEnv()

	flagExecute = "1 + 1"
	exitCode := runNonInteractive()
	if exitCode != 0 {
		t.Fatalf("exitCode - received: %v - expected: %v", exitCode, 0)
	}

	flagExecute = "1++"
	exitCode = runNonInteractive()
	if exitCode != 4 {
		t.Fatalf("exitCode - received: %v - expected: %v", exitCode, 4)
	}

	flagExecute = ""
}

type testInteractive struct {
	runLines   []string
	runOutputs []string
	runError   string
}

func TestRunInteractive(t *testing.T) {
	// empty strings in runOutputs will ignore read timeouts
	tests := []testInteractive{
		{runLines: []string{".."}, runError: "1:1 syntax error on '.' at 1:1"},
		{runLines: []string{"1++"}, runError: "1:1 invalid operation"},
		{runLines: []string{"var , b = 1, 2"}, runError: "1:7 syntax error: unexpected ','"},

		{runLines: []string{"", "1"}, runOutputs: []string{"", "1"}},
		{runLines: []string{"1 + 1"}, runOutputs: []string{"2"}},
		{runLines: []string{"a = 1", "b = 2", "a + b"}, runOutputs: []string{"1", "2", "3"}},
		{runLines: []string{"a = 1", "if a == 1 {", "b = 1", "b = 2", "}", "a"}, runOutputs: []string{"1", "", "", "", "2", "1"}},
		{runLines: []string{"a = 1", "for i = 0; i < 2; i++ {", "a++", "}", "a"}, runOutputs: []string{"1", "", "", "<nil>", "3"}},
	}
	runInteractiveTests(t, tests)
}

func runInteractiveTests(t *testing.T, tests []testInteractive) {
	// create logger
	// Note: logger is used for debugging since real stdout cannot be used
	logger = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.LUTC|log.Llongfile)
	gopath := os.Getenv("GOPATH")
	if gopath == "" {
		b, err := exec.Command("go", "env", "GOPATH").CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}
		gopath = strings.TrimSpace(string(b))
	}
	filehandle, err := os.OpenFile(filepath.Join(gopath, "bin", "anko_test.log"), os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		t.Fatal("OpenFile error:", err)
	}
	defer filehandle.Close()
	logger.SetOutput(filehandle)
	logger.Print("logger file created")

	// defer sending std back to real
	realStdin := os.Stdin
	realStderr := os.Stderr
	realStdout := os.Stdout
	defer setStd(realStdin, realStderr, realStdout)

	// create pipes
	readFromIn, writeToIn, err := os.Pipe()
	if err != nil {
		t.Fatal("Pipe error:", err)
	}
	os.Stdin = readFromIn
	logger.Print("pipe in created")
	readFromErr, writeToErr, err := os.Pipe()
	if err != nil {
		t.Fatal("Pipe error:", err)
	}
	os.Stderr = writeToErr
	logger.Print("pipe err created")
	readFromOut, writeToOut, err := os.Pipe()
	if err != nil {
		t.Fatal("Pipe error:", err)
	}
	os.Stdout = writeToOut
	logger.Print("pipe out created")

	// setup reader
	readerErr := bufio.NewReaderSize(readFromErr, 256)
	readerOut := bufio.NewReaderSize(readFromOut, 256)
	chanQuit := make(chan struct{}, 1)
	chanErr := make(chan string, 3)
	chanOut := make(chan string, 3)
	readTimeout := 10 * time.Millisecond
	var dataErr string
	var dataOut string

	go readerToChan(t, chanQuit, readerErr, chanErr)
	go readerToChan(t, chanQuit, readerOut, chanOut)

	go runInteractive()

	time.Sleep(readTimeout)

	// basic read and write to make sure things are working
	_, err = writeToIn.WriteString("1\n")
	if err != nil {
		t.Fatal("Stdin WriteString error:", err)
	}
	select {
	case dataOut = <-chanOut:
		if len(dataOut) > 0 && dataOut[0] == '>' {
			dataOut = dataOut[1:]
			dataOut = strings.TrimSpace(dataOut)
		}
		if dataOut != "1" {
			t.Fatalf("Stdout - received: %v - expected: %v - basic test", dataOut, "1")
		}
	case dataErr = <-chanErr:
		dataErr = strings.TrimSpace(dataErr)
		if dataErr != "" {
			t.Fatalf("Stderr - received: %v - expected: %v - basic test", dataErr, "")
		}
	case <-time.After(readTimeout):
		t.Fatal("read timeout for basic test")
	}

	// run tests
	logger.Print("running tests start")
	for _, test := range tests {

		for i, runLine := range test.runLines {

			_, err = writeToIn.WriteString(runLine + "\n")
			if err != nil {
				t.Fatal("Stdin WriteString error:", err)
			}

			select {
			case <-time.After(readTimeout):
				if test.runOutputs[i] != "" {
					t.Fatalf("read timeout for i: %v - runLines: %v", i, test.runLines)
				}
			case dataOut = <-chanOut:
				for len(dataOut) > 0 && dataOut[0] == '>' {
					dataOut = dataOut[1:]
					dataOut = strings.TrimSpace(dataOut)
				}
				if dataOut != test.runOutputs[i] {
					t.Fatalf("Stdout - received: %v - expected: %v - i: %v - runLines: %v", dataOut, test.runOutputs[i], i, test.runLines)
				}
			case dataErr = <-chanErr:
				dataErr = strings.TrimSpace(dataErr)
				if dataErr != test.runError {
					t.Fatalf("Stderr - received: %v - expected: %v - i: %v - runLines: %v", dataErr, test.runError, i, test.runLines)
				}

				// to clean output from error
				_, err = writeToIn.WriteString("1\n")
				if err != nil {
					t.Fatal("Stdin WriteString error:", err)
				}

				select {
				case dataOut = <-chanOut:
					for len(dataOut) > 0 && dataOut[0] == '>' {
						dataOut = dataOut[1:]
						dataOut = strings.TrimSpace(dataOut)
					}
					if dataOut != "1" {
						t.Fatalf("Stdout - received: %v - expected: %v - i: %v - runLines: %v", dataOut, test.runOutputs[i], i, test.runLines)
					}
				case <-time.After(readTimeout):
					t.Fatalf("read timeout for i: %v - runLines: %v", i, test.runLines)
				}

			}

		}

	}
	logger.Print("running tests end")

	// quit
	_, err = writeToIn.WriteString("quit()\n")
	if err != nil {
		t.Fatal("Stdin WriteString error:", err)
	}
	logger.Print("quit() sent")

	close(chanQuit)
	logger.Print("chanQuit closed")

	writeToErr.Close()
	writeToOut.Close()
	logger.Print("pipes closed")
}

func readerToChan(t *testing.T, chanQuit chan struct{}, reader *bufio.Reader, chanOut chan string) {
	logger.Print("readerToChan start")
	for {
		data, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			t.Fatal("readerToChan ReadString error:", err)
		}
		select {
		case <-chanQuit:
			logger.Print("readerToChan end")
			return
		default:
		}
		chanOut <- data
	}
}

func setStd(stdin *os.File, stderr *os.File, stdout *os.File) {
	os.Stdin = stdin
	os.Stderr = stderr
	os.Stdout = stdout
}
