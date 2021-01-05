package main

import (
	"fmt"
	"os"
	"log"
//	"encoding/json"
	"io/ioutil"
	"../internal/utils"
	"os/user"
)

var gitDuoConfigLocation = ".ssh"
var gitDuoConfigFile = "gitduo.json"


type Handler interface {
}

func HandlerFunc() {

}


func initialiseUserData() (userData utils.UserData) {
	user, err := user.Current()
	originaldir, err := os.Getwd()
	err = os.Chdir(user.HomeDir)
	err = os.Chdir(gitDuoConfigLocation)
	dir, err := os.Getwd()
	fmt.Print(dir)

	fmt.Println()
	dir, err = os.Getwd()
	fmt.Print(dir)
	file, err := ioutil.ReadFile(gitDuoConfigFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(file))

	var h utils.UserData

	//_ = json.Unmarshal([]byte(file), &h)

	h.Userdata.Main.Name = "juliankarnik"
	h.Userdata.Main.Email = "julian.karnik@ecs-digital.co.uk"
	h.Userdata.Main.Username = "julian17uk"
	h.Userdata.Main.Host = "github.com"
	h.Userdata.Work.Name = "juliankarnik-work"
	h.Userdata.Work.Email = "julian.karnik@ecs.co.uk"
	h.Userdata.Work.Username = "julian19uk"
	h.Userdata.Work.Host = "github-work"

	err = os.Chdir(originaldir)
	return h
}

func main() {
	argCount := len(os.Args)
	if argCount == 1 {
		fmt.Println("No valid arguments")
	
		return
	}
	argsWithoutProg := os.Args[1:]
	arg := os.Args[1]
	workingrepo := utils.Repocheck()


	h := initialiseUserData()
	fmt.Println("Value of Work.Email in struct h: " + h.Userdata.Work.Email)

	fmt.Println(argsWithoutProg)
	fmt.Println(arg)

	if workingrepo == false {
		if arg == "set" {
			if argCount == 2 {
				fmt.Println("Not valid arguments for set command, expecting boolean 1 = private, 0 = public")
				utils.Help()
				return
			}
			if os.Args[2] == "1" {
				utils.Set(true)
			} else if os.Args[2] == "0" {
				utils.Set(false)
			}	
			return	
		} else {
			switch arg {
				case "which": fmt.Println("Not inside a working git directory")
				case "main": fmt.Println("Not inside a working git directory")
				case "work": fmt.Println("Not inside a working git directory")
				case "help": 
				default: fmt.Println("Not a valid command")
			}
			utils.Help()
			return
		}
	}

	switch arg {
	case "which":
		h.Which()
	case "help":
		utils.Help()
	case "main":
		utils.Mainfunc()
	case "work":
		utils.Work()
	case "set":
		fmt.Println("Command not valid inside a working git directory")
	default:
		fmt.Println("Not a valid command")
//		utils.Help()
	}
	
}




