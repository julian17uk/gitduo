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
	validCommands := []string{"help", "which", "main", "work", "set", "pat"}
	repoCommands := []string{"which", "main", "work"}
	command, paramcount, params := utils.FilterArgs(os.Args[1:])
	workingrepo := utils.Repocheck()
	filedata := initialiseFiles()
	h := utils.Handler(filedata)

	if !utils.Find(validCommands, command) {
		fmt.Println("not a valid command please see help")
		return
	}

	if paramcount == 0 {
		switch command {
		case "pat":
			fmt.Println("Personal access token:" + h.GetToken())
			return
		case "help":
			utils.Help()
			return
		case "set":
			fmt.Println("invalid arguments for set command, expecting boolean 1 = private, 0 = public")
			return
		}

		if workingrepo == true {
			switch command {
			case "which":
				h.Which()
				return
			case "main":
				a := h.InitialiseActiveUser("main")
				a.SetUser()
				return
			case "work":
				a := h.InitialiseActiveUser("work")
				a.SetUser()
				return
			}
		} else {
			if utils.Find(repoCommands, command) {
				fmt.Println("not inside a working git directory")
			}
		}
	}
	
	if paramcount == 1 {
		if workingrepo == true {
			if command == "set" {
				fmt.Println("set not valid inside a working git directory")
				return
			}
		} else {
			if command == "set" {
				switch params[0] {
				case "0":
					h.SetRepo(false)
				case "1":
					h.SetRepo(true)
				default:
					fmt.Println("invalid arguments for set command, expecting boolean 1 = private, 0 = public")
				}
				return	 
			}
		}
	}

	fmt.Println("Too many arguments")
}




