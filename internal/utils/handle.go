package utils

import (
	"fmt"
)

func (h *Handle)HandleFunc(command string, paramcount int, params []string) {
	validCommands := []string{"help", "which", "main", "work", "set", "pat"}
	repoCommands := []string{"which", "main", "work"}
	workingrepo := Repocheck()

	if !Find(validCommands, command) {
		fmt.Println("not a valid command please see help")
		return
	}

	if paramcount == 0 {
		switch command {
		case "pat":
			fmt.Println("Personal access token:" + h.GetToken())
			return
		case "help":
			Help()
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
				h.SetUser(a)
				return
			case "work":
				a := h.InitialiseActiveUser("work")
				h.SetUser(a)
				return
			}
		} else {
			if Find(repoCommands, command) {
				fmt.Println("not inside a working git directory")
				return
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

