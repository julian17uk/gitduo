package utils

import(
	"os"
	"bytes"
	"strings"
	"os/exec"
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

	cmd := exec.Command("curl", "-i", "-u", user, gitapiurl, "-d", input)
	var out bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	s := out.String()
	if n := strings.IndexByte(out.String(), '\n'); n >=0 {
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

	cmd := exec.Command("git", "remote", "add", "origin", ssh)
	err := cmd.Run()
	if err != nil {
		return err
	}
	cmd = exec.Command("git", "push", "--set-upstream", "origin", "master")
	err = cmd.Run()
	if err != nil {
		return err
	}
	return nil
}