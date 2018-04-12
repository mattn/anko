// +build !appengine

package main

import (
	"os"
	"path/filepath"
	"runtime"
	"testing"
)

func TestMain(m *testing.M) {
	parseFlags()
	code := m.Run()
	os.Exit(code)
}

func TestRunNonInteractiveFile(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	testDir := filepath.Dir(filename) + "/core/testdata/"
	t.Log(testDir)
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
