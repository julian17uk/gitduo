package main

import (
	"os"
	"../internal/utils"
)

var PersonalAccessTokenDirectory = "/.ssh"
var PersonalAccessTokenFileName = "gittoken"
var ConfUserDataDirectory = "/.ssh"
var ConfUserDataFileName = "gitconf"

func initialiseFiles() utils.Files {
	var f utils.Files
	f.TokenDir = PersonalAccessTokenDirectory
	f.TokenFileName = PersonalAccessTokenFileName
	f.ConfDir = ConfUserDataDirectory
	f.ConfFileName = ConfUserDataFileName
	return f
}

func main() {
	command, paramcount, params := utils.FilterArgs(os.Args[1:])
	filedata := initialiseFiles()
	runner := utils.RealRunner{}
	h, err := utils.Handler(filedata, runner)
	if err != nil {
		return
	}

	h.HandleFunc(command, paramcount, params)
}




