package main

import (
	"testing"
	"fmt"
	"../internal/utils"
	"../internal/testutils"
)

type TestRunner struct{}

func (r TestRunner) Run(command string, args ...string) ([]byte, error) {
	out := []byte(`ok`)
	return out, nil
}

func TestInitialiseActiveUser(t *testing.T) {
	user1 := utils.ActiveUser{ "john", "john@email.com", "john1", "github.com"}
	user2 := utils.ActiveUser{ "johnb", "john@work.com", "john1-work", "github-work"}
	var testRunner TestRunner
	h := testutils.CreateTestHandler(testRunner)

	var tests = []struct {
		h			utils.Handle
		name		string
		expected	utils.ActiveUser
	}{
		{ h, "main", user1 },
		{ h, "work", user2 },
	}

	fmt.Println(user1.Name)

	for _, test := range tests {
		actual := test.h.InitialiseActiveUser(test.name)
		if actual.Name != test.expected.Name {
			t.Errorf("Initialise active user failed, found %v, expected %v", actual.Name, test.expected.Name)
		}
		if actual.Email != test.expected.Email {
			t.Errorf("Initialise active user failed, found %v, expected %v", actual.Name, test.expected.Name)
		}
		if actual.Username != test.expected.Username {
			t.Errorf("Initialise active user failed, found %v, expected %v", actual.Name, test.expected.Name)
		}
		if actual.Host != test.expected.Host {
			t.Errorf("Initialise active user failed, found %v, expected %v", actual.Name, test.expected.Name)
		}
	}
}
