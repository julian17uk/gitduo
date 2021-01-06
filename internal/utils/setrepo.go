package utils

import(
	"fmt"
	"os/exec"
	"io/ioutil"
)

func (h *Handle) SetRepo(private bool) string {
	files, _ := ioutil.ReadDir(".")
	if len(files) != 0 {
		fmt.Println("directory is not empty")
		return ""
	}
	reponame := InputRequest("new repo")

	token := h.GetToken()
	if token == "" {
		fmt.Println("Private Access Token missing")
		token = InputRequest("github private access token")
	}
	h.Token = token

	err := localGitSetup()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	result, err := h.CreateRemoteRepo(reponame, private)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	err = h.populateRemoteRepo(reponame)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	fmt.Println(result)
	return result
}

func localGitSetup() error {
	cmd := exec.Command("touch", "README.md")
	err := cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("git", "init")
	err = cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("git", "add", "README.md")
	err = cmd.Run()
	if err != nil {
		return err
	}

	cmd = exec.Command("git", "commit", "-m", "first commit")
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
