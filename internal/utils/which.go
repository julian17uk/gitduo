package utils

import(
	"fmt"
	"regexp"
)

func (h *Handle) Which() {
	out, err := h.Runner.Run("git", "remote", "-v")
	if err != nil {
		fmt.Println(err)
		return
	}
	regex, err := regexp.Compile(`\@(.*?)\:`)
	result := regex.Find(out)
	sz := len(result)
	result = result[:sz-1]
	result = result[1:]

	if string(result) == h.Conf.Main.Host {
		fmt.Println("main is active user with github username " + h.Conf.Main.Username)
		return
	}
	if string(result) == h.Conf.Work.Host {
		fmt.Println("work is active user with github username " + h.Conf.Work.Username)
		return
	}
	fmt.Println("active user is neither main or work...")
	return
}