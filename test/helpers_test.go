package main

import (
	"testing"
	"../internal/utils"
)


func TestFind(t *testing.T) {
	validCommands := []string{"help", "which", "main", "work", "set", "pat"}
	
	var tests = []struct {
		command	string
		expected bool
	}{
		{ "", false},
		{ "main", true},
		{ "pat", true},
		{ "invalid", false},
	}

	for _, test := range tests {
		if actual := utils.Find(validCommands, test.command); actual != test.expected {
			t.Error("Find test failed, unexpected output")
		}
	}
}

func TestFilterArgs(t *testing.T) {
	var testargs1 = []string{ "" }
	var testargs2 = []string{ "main" }
	var testargs3 = []string{ "set", "1" }
	var testargs4 = []string{ "set", "1", "2", "3"}
	var tests = []struct {
		args []string
		command string
		paramcount int
	}{
		{
			testargs1, "", 0,
		},
		{
			testargs2, "main", 0,
		},
		{
			testargs3, "set", 1,
		},
		{
			testargs4, "set", 3,
		},
	}

	for _, test := range tests {
		command, paramcount, _ := utils.FilterArgs(test.args) 
		if command != test.command {
			t.Errorf("TestFilterArgs failed, unexpected command")
		}
		if paramcount != test.paramcount {
			t.Errorf("TestFilterArgs failed, unexpected number of params")
		}
	}
}

// func InputRequest(inputname string) string {
// 	fmt.Println("Please enter a name for the " + inputname + ":")
// 	reader := bufio.NewReader(os.Stdin)
// 	input, _ := reader.ReadString('\n')
// 	input = strings.TrimSuffix(input, "\n")
// 	input = strings.TrimSpace(input)
// 	return input
// }
