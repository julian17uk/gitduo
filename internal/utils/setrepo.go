package utils

import(
	"fmt"
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

	err := h.localGitSetup()
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

func (h *Handle)localGitSetup() error {
	_, err := h.Runner.Run("touch", "README.md")
	if err != nil {
		return err
	}

	_, err = h.Runner.Run("git", "init")
	if err != nil {
		return err
	}
	_, err = h.Runner.Run("git", "add", "README.md")
	if err != nil {
		return err
	}
	_, err = h.Runner.Run("git", "commit", "-m", "first commit")
	if err != nil {
		return err
	}
	return nil
}
