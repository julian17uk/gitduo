package utils

import(
    "fmt"
    "os/exec"
    "regexp"
    "bytes"
    "bufio"
    "os"
	"strings"
)

var usage = `usage: gitduo <command> [<args>]

These are the Gitduo commands available:

    help    list functionality
    pat     displays the github personal access token (if available)
    
In a Git working directory:

    which   shows which is the active git user
    main    set main git user in the current location
    work    set work git user in the current location
	
In a non Git working directory:

    set     initailize a repository in an empty location and provisions it on github. This takes 1 arg of value 1 for private and 0 for public
			(example gitduo set 1)`
			
func Help() {
	fmt.Println(usage)
}


func Repocheck() bool {
	cmd := exec.Command("git", "status")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return (false)
	}

	return true
}

func RemoteRepoName() string {
	cmd := exec.Command("git", "config", "remote.origin.url")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	Reg1, err := regexp.Compile(`\:(.*)`) 
	result := Reg1.Find([]byte(out.String()))
	result = result[1:]
	return string(result)
}

func InputRequest(inputname string) string {
	fmt.Println("Please enter a name for the " + inputname + ":")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSpace(input)
	return input
}