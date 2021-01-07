package utils

import(
	"strings"
	"fmt"
)

var gitapiurl = "https://api.github.com/user/repos"

func (h *Handle)CreateRemoteRepo(reponame string, private bool) (string, error) {
	var result string
	privateflag := "true"
	if private != true {
		privateflag = "false"
	} 
	username := h.Conf.Main.Username
	input := `{"name":"` + reponame + `","private":` + privateflag + `}`
	user := username+":"+ h.Token

	out, err := h.Runner.Run("curl", "-i", "-u", user, gitapiurl, "-d", input)
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	s := string(out)
	if n := strings.IndexByte(s, '\n'); n >=0 {
		result = s[:n]
	} else {
		result = s
	}
	return result, nil
}

func (h *Handle)populateRemoteRepo(reponame string) error {
	host := h.Conf.Main.Host
	username := h.Conf.Main.Username
	ssh := "git@" + host + ":" + username + "/" + reponame + ".git"
	
	_, err := h.Runner.Run("git", "remote", "add", "origin", ssh)
	if err != nil {
		return err
	}
	_, err = h.Runner.Run("git", "push", "--set-upstream", "origin", "master")
	if err != nil {
		return err
	}
	return nil
}