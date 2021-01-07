package testutils

import (
	"../utils"
	"encoding/json"
	"bytes"
	"log"
	"os"
	"sync"
	"io"
)

type TestRunner struct{}

func (r TestRunner) Run(command string, args ...string) ([]byte, error) {
	out := []byte(`ok`)
	return out, nil
}

// func CaptureOutput(f func()) string {
// 	var buf bytes.Buffer
// 	log.SetOutput(&buf)
// 	f()
// 	log.SetOutput(os.Stderr)
// 	return buf.String()
// }

func CaptureOutput(f func()) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
		log.SetOutput(os.Stderr)
	}()
	os.Stdout = writer
	os.Stderr = writer
	log.SetOutput(writer)
	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait()
	f()
	writer.Close()
	return <-out
}

func CreateTestHandler(testRunner utils.Runner) (utils.Handle) {
	var f utils.Files
	token := "abc12325345dsdf"
	test1 := []byte(`{ "Main":
        { "Name": "john",
          "Email": "john@email.com",
          "Username": "john1",
          "Host": "github.com"
        },
  			"Work":
        { "Name": "johnb",
          "Email": "john@work.com",
          "Username": "john1-work",
          "Host": "github-work"
        }}`)

	var c1 utils.Conf
	json.Unmarshal(test1, &c1)
	h := utils.Handle{ c1, f, token, testRunner }
	return h
} 