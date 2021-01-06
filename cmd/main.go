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
	argCount := len(os.Args)
	if argCount == 1 {
		fmt.Println("no valid arguments please see help")
		return
	}
	arg := os.Args[1]
	workingrepo := utils.Repocheck()
	filedata := initialiseFiles()
	h := utils.Handler(filedata)

	if arg == "help" {
		utils.Help()
		return
	}

	if workingrepo == false {
		if arg == "set" {
			if argCount == 2 {
				fmt.Println("invalid arguments for set command, expecting boolean 1 = private, 0 = public")
				return
			}
			if os.Args[2] == "1" {
				h.SetRepo(true)
			} else if os.Args[2] == "0" {
				h.SetRepo(false)
			}	
			return	
		} else {
			fmt.Println("not inside a working git directory")
			return
		}
	}

	switch arg {
	case "which":
		h.Which()
	case "main":
		a := h.InitialiseActiveUser("main")
		a.SetUser()
	case "work":
		a := h.InitialiseActiveUser("work")
		a.SetUser()
	case "pat":
		fmt.Println("Personal access token:" + h.GetToken())
	case "set":
		fmt.Println("set not valid inside a working git directory")
	default:
		fmt.Println("not a valid command please see help")
	}
}




