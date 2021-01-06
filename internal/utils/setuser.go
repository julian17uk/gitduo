package utils

import(
	"fmt"
	"os/exec"
)

func (a *ActiveUser) SetUser() {
    cmd := exec.Command("git", "config", "user.name", a.Name)
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	cmd = exec.Command("git", "config", "user.email", a.Email)
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}

	reponame := RemoteRepoName()
	url := "git@" + a.Host + ":" + reponame
	cmd = exec.Command("git", "remote", "set-url", "origin", url)
	err = cmd.Run()
	if err != nil {
		fmt.Println(err)
		return
	}
}