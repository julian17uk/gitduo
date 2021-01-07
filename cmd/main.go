package main

import (
	"fmt"
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
	h := utils.Handler(filedata, runner)

	h.handleFunc(command, paramcount, params)
}




