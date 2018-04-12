// +build !appengine

package main

import (
	"bufio"
	"io"
	"log"
	"os"
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
	testDir := filepath.Dir(filename) + "/core/testdata/"
	setupEnv()

	file = testDir + "not-found.ank"
	exitCode := runNonInteractive()
	if exitCode != 2 {
		t.Fatalf("exitCode - received: %v - expected: %v", exitCode, 2)
	}

	file = testDir + "broken.ank"
	exitCode = runNonInteractive()
	os.Args = []string{os.Args[0]}
	if exitCode != 4 {
		t.Fatalf("exitCode - received: %v - expected: %v", exitCode, 4)
	}

	file = testDir + "test.ank"
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
	runSource string
	runError  string
	runOutput string
}

func TestRunInteractive(t *testing.T) {
	tests := []testInteractive{
		testInteractive{runSource: "1 + 1", runOutput: "2"},
		testInteractive{runSource: "1++", runError: "1:1 invalid operation"},
		// TODO: add more tests
	}
	runInteractiveTests(t, tests)
}

func runInteractiveTests(t *testing.T, tests []testInteractive) {
	// Note: logger is only used for debugging since real stdout cannot be used
	logger = log.New(os.Stderr, "", log.Ldate|log.Ltime|log.LUTC|log.Llongfile)
	gopath := os.Getenv("GOPATH")
	filehandle, err := os.OpenFile(gopath+"/bin/anko_test.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		t.Fatal("OpenFile error: ", err)
	}
	defer filehandle.Close()
	logger.SetOutput(filehandle)

	realStdin := os.Stdin
	readFromIn, writeToIn, err := os.Pipe()
	if err != nil {
		t.Fatal("Pipe error: ", err)
	}
	os.Stdin = readFromIn

	realStderr := os.Stderr
	readFromErr, writeToErr, err := os.Pipe()
	if err != nil {
		os.Stdin = realStdin
		t.Fatal("Pipe error: ", err)
	}
	os.Stderr = writeToErr

	realStdout := os.Stdout
	readFromOut, writeToOut, err := os.Pipe()
	if err != nil {
		os.Stdin = realStdin
		os.Stderr = realStderr
		t.Fatal("Pipe error: ", err)
	}
	os.Stdout = writeToOut

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

	for _, test := range tests {
		go runInteractive()

		select {
		case dataErr = <-chanErr:
			os.Stdin = realStdin
			os.Stderr = realStderr
			os.Stdout = realStdout
			t.Fatalf("Stderr - received: %v - expected: %v - runSource: %v", dataErr, "", test.runSource)
		case <-time.After(readTimeout):
		}

		select {
		case dataOut = <-chanOut:
			os.Stdin = realStdin
			os.Stderr = realStderr
			os.Stdout = realStdout
			t.Fatalf("Stdout - received: %v - expected: %v - runSource: %v", dataOut, "", test.runSource)
		case <-time.After(readTimeout):
		}

		_, err = writeToIn.WriteString(test.runSource + "\n")
		if err != nil {
			os.Stdin = realStdin
			os.Stderr = realStderr
			os.Stdout = realStdout
			t.Fatal("Stdin WriteString error: ", err)
		}

		select {
		case dataErr = <-chanErr:
			dataErr = strings.TrimSpace(dataErr)
			if dataErr != test.runError {
				os.Stdin = realStdin
				os.Stderr = realStderr
				os.Stdout = realStdout
				t.Fatalf("Stderr - received: %v - expected: %v - runSource: %v", dataErr, test.runError, test.runSource)
			}
		case <-time.After(readTimeout):
			if test.runError != "" {
				os.Stdin = realStdin
				os.Stderr = realStderr
				os.Stdout = realStdout
				t.Fatalf("Stderr - received: %v - expected: %v - runSource: %v", "", test.runError, test.runSource)
			}
		}

		select {
		case dataOut = <-chanOut:
			if len(dataOut) > 0 && dataOut[0] == '>' {
				dataOut = dataOut[1:]
			}
			dataOut = strings.TrimSpace(dataOut)
			if dataOut != test.runOutput {
				os.Stdin = realStdin
				os.Stderr = realStderr
				os.Stdout = realStdout
				t.Fatalf("Stdout - received: %v - expected: %v - runSource: %v", dataOut, test.runOutput, test.runSource)
			}
		case <-time.After(readTimeout):
			if test.runOutput != "" {
				os.Stdin = realStdin
				os.Stderr = realStderr
				os.Stdout = realStdout
				t.Fatalf("Stdout - received: %v - expected: %v - runSource: %v", "", test.runOutput, test.runSource)
			}
		}

	}

	close(chanQuit)

	_, err = writeToIn.WriteString("\n")
	if err != nil {
		os.Stdin = realStdin
		os.Stderr = realStderr
		os.Stdout = realStdout
		t.Fatal("Stdin WriteString error: ", err)
	}
	_, err = writeToIn.WriteString("quit()\n")
	if err != nil {
		os.Stdin = realStdin
		os.Stderr = realStderr
		os.Stdout = realStdout
		t.Fatal("Stdin WriteString error: ", err)
	}

	writeToErr.Close()
	writeToOut.Close()

	os.Stdin = realStdin
	os.Stderr = realStderr
	os.Stdout = realStdout
}

func readerToChan(t *testing.T, chanQuit chan struct{}, reader *bufio.Reader, chanOut chan string) {
	for {
		data, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			t.Fatal("readerToChan ReadString error: ", err)
		}
		select {
		case <-chanQuit:
			return
		default:
		}
		chanOut <- data
	}
}
