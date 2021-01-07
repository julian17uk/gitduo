package main

import (
	"testing"
	"../internal/testutils"
	"strings"

)

type TestRunnerRepoCheckOK struct{}

func (r TestRunnerRepoCheckOK) Run(command string, args ...string) ([]byte, error) {
	out := []byte(`git@github.com:john345/repo123`)
	return out, nil
}

type TestRunnerRepoCheckNotOK struct{}

func (r TestRunnerRepoCheckNotOK) Run(command string, args ...string) ([]byte, error) {
	out := []byte(`git@github.com:john345/repo123`)
	return out, nil
}

func TestHandleFuncRepoOK(t *testing.T) {
	var testRunner TestRunnerRepoCheckOK
	h := testutils.CreateTestHandler(testRunner)
	var params []string

	var tests = []struct {
		command			string
		paramcount		int
		params			[]string
		expectedOutput	string
	}{
		{ "invalid", 0, params, "not a valid command please see help" },
		{ "set", 0, params, "invalid arguments for set command, expecting boolean 1 = private, 0 = public" },
		{ "which", 4, params, "Too many arguments" },
		{ "set", 1, params, "set not valid inside a working git directory" },
		{ "main", 7, params, "Too many arguments" },
	}

	for _, test := range tests {
		out := testutils.CaptureOutput(func() { h.HandleFunc(test.command, test.paramcount, test.params) })
		out = strings.TrimSuffix(out, "\n")
		out = strings.TrimSpace(out)
		if out != test.expectedOutput {
			t.Errorf("Find test failed, expected output %v, actual output %v ", test.expectedOutput, out)
		}
	}
}

func TestHandleFuncRepoNotOK(t *testing.T) {
	var testRunner TestRunnerRepoCheckNotOK
	h := testutils.CreateTestHandler(testRunner)
	var params []string
	var tests = []struct {
		command			string
		paramcount		int
		params			[]string
		expectedOutput	string
	}{
		{ "invalid", 0, params, "not a valid command please see help" },
		{ "set", 0, params, "invalid arguments for set command, expecting boolean 1 = private, 0 = public" },
		{ "which", 4, params, "Too many arguments" },
		{ "set", 1, params, "set not valid inside a working git directory" },
		{ "main", 7, params, "Too many arguments" },
	}

	for _, test := range tests {
		out := testutils.CaptureOutput(func() { h.HandleFunc(test.command, test.paramcount, test.params) })
		out = strings.TrimSuffix(out, "\n")
		out = strings.TrimSpace(out)
		if out != test.expectedOutput {
			t.Errorf("Find test failed, expected output %v, actual output %v ", test.expectedOutput, out)
		}
	}	
}
