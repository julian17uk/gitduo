package utils

import (
	"fmt"
	"os/exec"
	"os"
)

var usage = `usage: gitduo <command> [<args>]

These are the Gitduo commands available:

    help    list functionality
    
In a Git working directory:

    which   shows which is the active git user
    main    set main git user in the current location
    work    set work git user in the current location
	
In a non Git working directory:

    set     initailize a repository in an empty location and provisions it on github. This takes 1 arg of value 1 for private and 0 for public
			(example gitduo set 1)`


func Repocheck() bool {
	// fmt.Println("Inside repocheck function")
	cmd := exec.Command("git", "rev-parse", "--is-inside-work-tree")
	// cmd.Stdout = os.Stdout
	// cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
//		fmt.Println(err)
		return (false)
	}
	return true
}

func Help() {
	fmt.Println(usage)
}

func Which() {
	fmt.Println("Inside which function")
	cmd := exec.Command("git", "config", "user.email")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func Work() {
	fmt.Println("Inside work function")
}


func Mainfunc() {
	fmt.Println("Inside mainfunction")

}


func Set(value bool) {
	fmt.Printf("Inside set function %v", value)

}