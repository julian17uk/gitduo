package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"../internal/utils"
)

func main() {
	argCount := len(os.Args)
	if argCount == 1 {
		fmt.Println("No valid arguments")
	
		return
	}
	argsWithoutProg := os.Args[1:]
	arg := os.Args[1]
	workingrepo := utils.Repocheck()

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
		utils.Which()
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





func helpOLD() {
	file, err := os.Open("doc.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	b, err := ioutil.ReadAll(file)
	fmt.Print(string(b))
}



