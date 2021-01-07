package utils

import(
	"fmt"
)

func (h *Handle) SetUser(a ActiveUser) {
	_, err := h.Runner.Run("git", "config", "user.name", a.Name)
 	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = h.Runner.Run("git", "config", "user.email", a.Email)
	if err != nil {
		fmt.Println(err)
		return
	}

	reponame := RemoteRepoName()
	url := "git@" + a.Host + ":" + reponame
	_, err = h.Runner.Run("git", "remote", "set-url", "origin", url)
	if err != nil {
		fmt.Println(err)
		return
	}
}