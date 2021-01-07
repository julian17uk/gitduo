package main

import (
	"testing"
	"../internal/testutils"
	"../internal/utils"
	"strings"

)

type TestRunnerSetRepo struct{}

func (r TestRunnerSetRepo) Run(command string, args ...string) ([]byte, error) {
	var out []byte
	return out, nil
}

func TestSetRepoNoErrorOnOutput(t *testing.T) {
	var testRunner TestRunnerSetRepo
	h := testutils.CreateTestHandler(testRunner)
	a := utils.ActiveUser{ "john", "john@email.com", "john1", "github.com"}

	out := testutils.CaptureOutput(func() { h.SetUser(a) })
	out = strings.TrimSuffix(out, "\n")
	out = strings.TrimSpace(out)
	expectedOutput := ""
	if out != expectedOutput {
		t.Errorf("Find test failed, expected output %v, actual output %v ", expectedOutput, out)
	}
}