package utils

import (
	"fmt"
	"os/user"
	"os/exec"
	"os"
	"regexp"
	"bytes"
	"strings"
	"bufio"
	"io/ioutil"
	"encoding/json"
)

type Conf struct {
	Main struct {
		Email    string `json:"Email"`
		Host     string `json:"Host"`
		Name     string `json:"Name"`
		Username string `json:"Username"`
	} `json:"Main"`
	Work struct {
		Email    string `json:"Email"`
		Host     string `json:"Host"`
		Name     string `json:"Name"`
		Username string `json:"Username"`
	} `json:"Work"`
}

type ActiveUser struct {
	Name 	string
	Email	string
	Username string
	Host	string
}

type Files struct {
	TokenDir 		string
	TokenFileName 	string
	ConfDir 		string
	ConfFileName 	string
}

type Handle struct {
	Conf 	Conf
	Files	Files
	Token	string
}

func Handler(f Files) (*Handle) {
	var h Handle
	h.Conf = GetConf(f)
	h.Files = f

	return &h
}

func (h *Handle)InitialiseActiveUser(name string) (*ActiveUser) {
	var a ActiveUser
	switch name {
	case "work":
		a.Name = h.Conf.Work.Name
		a.Email = h.Conf.Work.Email
		a.Host = h.Conf.Work.Host
		a.Username = h.Conf.Work.Username
	case "main":
		a.Name = h.Conf.Main.Name
		a.Email = h.Conf.Main.Email
		a.Host = h.Conf.Main.Host
		a.Username = h.Conf.Main.Username
	}
	return &a
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

func (h *Handle)GetToken() string{
	usr, _ := user.Current()
	originaldir, _ := os.Getwd()
	err := os.Chdir(usr.HomeDir + h.Files.TokenDir)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	if _, err = os.Stat(h.Files.TokenFileName); os.IsNotExist(err) {
		err = os.Chdir(originaldir)
		return ""
	} else {
		b, err := ioutil.ReadFile(h.Files.TokenFileName)
		if err != nil {
			fmt.Println(err)
			return ""
		}
		err = os.Chdir(originaldir)
		token := string(b)
		token = strings.TrimSuffix(token, "\n")
		token = strings.TrimSpace(token)
		return token
	}
}

func GetConf(f Files) (Conf) {
	var c Conf

	usr, _ := user.Current()
	originaldir, _ := os.Getwd()
	err := os.Chdir(usr.HomeDir + f.ConfDir)
	if err != nil {
		fmt.Println(err)
		return c
	}

	if _, err = os.Stat(f.ConfFileName); os.IsNotExist(err) {
		err = os.Chdir(originaldir)
		fmt.Println("Error: Conf file missing")
		return c
	} else {
		b, err := ioutil.ReadFile(f.ConfFileName)
		if err != nil {
			fmt.Println(err)
			return c
		}
		err = os.Chdir(originaldir)

		json.Unmarshal(b, &c)
		return c
	}
}