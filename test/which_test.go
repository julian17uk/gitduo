package main

import (
	"testing"
	"../internal/testutils"
	"strings"

)


type TestRunnerWhich struct{}

func (r TestRunnerWhich) Run(command string, args ...string) ([]byte, error) {
	out := []byte(`git@github.com:john345/repo123`)
	return out, nil
}

func TestWhich(t *testing.T) {
	var testRunner TestRunnerWhich
	h := testutils.CreateTestHandler(testRunner)

	out := testutils.CaptureOutput(func() { h.Which() })
	out = strings.TrimSuffix(out, "\n")
	out = strings.TrimSpace(out)
	expectedOutput := "main is active user with github username john1"
	if out != expectedOutput {
		t.Errorf("Find test failed, expected output %v, actual output %v ", expectedOutput, out)
	}
}
