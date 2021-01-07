package utils

import (
	"fmt"
	"os/user"
	"os"
	"strings"
	"io/ioutil"
	"encoding/json"
	"os/exec"
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
	Runner  Runner
}

type RealRunner struct{}

type Runner interface {
	Run(string, ...string) ([]byte, error)
}

func (r RealRunner) Run(command string, args ...string) ([]byte, error) {
	out, err := exec.Command(command, args...).Output()
	return out, err
}

func Handler(f Files, r Runner) (*Handle, error) {
	var h Handle
	c, err := GetConf(f)
	if err != nil {
		return &h, err
	}
	h.Conf = c
	h.Files = f
	h.Runner = r
	return &h, nil
}

func (h *Handle)InitialiseActiveUser(name string) (ActiveUser) {
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
	return a
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

func GetConf(f Files) (Conf, error) {
	var c Conf

	usr, _ := user.Current()
	originaldir, _ := os.Getwd()
	err := os.Chdir(usr.HomeDir + f.ConfDir)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Expecting file " + f.ConfFileName + " in directory " + f.ConfDir + "P lease set README.md for details")
		return c, err
	}

	if _, err = os.Stat(f.ConfFileName); os.IsNotExist(err) {
		_ = os.Chdir(originaldir)
		fmt.Println("Error: Conf file missing")
		fmt.Println("Expecting file " + f.ConfFileName + " in directory ~" + f.ConfDir + " Please set README.md for details")
		return c, err

	} else {
		b, err := ioutil.ReadFile(f.ConfFileName)
		if err != nil {
			fmt.Println(err)
			return c, err
		}
		err = os.Chdir(originaldir)
		if err != nil {
			return c, err
		}

		err = json.Unmarshal(b, &c)
		if err != nil {
			fmt.Println("Error: " + f.ConfFileName + " unexpected format")
			fmt.Println("For expected format see gitconf.sample")
			var c1 Conf
			return c1, err
		}
		return c, nil
	}
}