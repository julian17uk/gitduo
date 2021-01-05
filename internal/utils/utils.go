package utils

import (
	"fmt"
	"os/exec"
	"os"
	"regexp"
	"bytes"
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

// type UserData struct {
// 	Main	User `json:"Main"`
// 	Work	User `json:"Work"`
// }

// type User struct {
// 	Name	string `json:"Name"`
// 	Email	string `json:"Email"`
// 	Username string `json:"Username"`
// 	Host	string `json:"Hose"`
// }

type UserData struct {
	Userdata struct {
		Main struct {
			Name     string `json:"Name"`
			Email    string `json:"Email"`
			Username string `json:"Username"`
			Host     string `json:"Host"`
		} `json:"Main"`
		Work struct {
			Name     string `json:"Name"`
			Email    string `json:"Email"`
			Username string `json:"Username"`
			Host     string `json:"Host"`
		} `json:"Work"`
	} `json:"Userdata"`
}

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

func (h *UserData) Which() {
	fmt.Println("Inside which function")
	cmd := exec.Command("git", "remote", "-v")
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("output: " + out.String())
	// output := fmt.Sprint(cmd.Stdout)
	// fmt.Println(output)
	fmt.Println("Success using structs in other packages")
	// Add regex parser and use gitduo.config file to print out correct result
//	searchstring := `git@github.com:julian17uk/gitduo.git`
	regex2, err := regexp.Compile(`\@(.*?)\:`)
	regex, err := regexp.Compile(`\:(.*?)\/`)
fmt.Println("Here is tdfsdfsdfhe matched regexp: ")
// TODO: Find searchstring from cmd.STdout
	result := regex2.Find([]byte(out.String()))

	fmt.Println("Here is the matched regexp: ")
	fmt.Println(string(result))
	fmt.Println("Post result")
	sz := len(result)
	result = result[:sz-1]
	result = result[1:]
	fmt.Println(string(result))

	resultstr := string(result)
	username := regex.Find([]byte(out.String()))
	sy := len(username)
	username = username[:sy-1]
	username = username[1:]
	fmt.Println("username: ")
	fmt.Println(string(username))
	// re := regexp.MustCompile(`foo.?`)
	// fmt.Printf("%q\n", re.Find([]byte('seafood fool')))

	if resultstr == h.Userdata.Main.Host {
		fmt.Println("User is Main")
	}
	if resultstr == h.Userdata.Work.Host {
		fmt.Println("User is Work")
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