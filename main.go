package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"runtime"
	"strings"
)

func main() {
	dotfiles := os.Getenv("dotFilesDir")
	dotfiles = strings.ReplaceAll(dotfiles, "\\", "/")
	fmt.Printf("dotfiles: %s\n", dotfiles)
	if len(dotfiles) == 0 {
		fmt.Print("dotFilesDir env is not set, guessing ")
		df, err := os.UserHomeDir()
		if err != nil {
			dotfiles = "~/.files"
		} else {
			dotfiles = df + "/.files"
		}
		fmt.Printf("dotfiles as: %s\n", dotfiles)
	}

	command := fmt.Sprintf("%s/bin/go-install", dotfiles)
	var Args []string
	if runtime.GOOS == "windows" {
		Args = append(Args, "C:\\Program Files\\Git\\usr\\bin\\bash.exe")
	}

	Args = append(Args, command)
	Args = append(Args, os.Args[1:]...)

	fmt.Printf("exec: %v\n", Args)
	cmd := exec.Command(Args[0], Args[1:]...)
	stdout, _ := cmd.StdoutPipe()
	stderr, _ := cmd.StderrPipe()
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(254)
	}

	go print(stdout, false)
	go print(stderr, true)

	err = cmd.Wait()

	if err != nil {
		fmt.Println(err)
		os.Exit(255)
	}
}

func print(stdout io.ReadCloser, is_stderr bool) {
	r := bufio.NewReader(stdout)
	for {
		line, _, err := r.ReadLine()
		if err != nil {
			break
		}
		fd := os.Stdout
		if is_stderr {
			fd = os.Stderr
		}
		fmt.Fprintf(fd, "%s\n", line)
	}
}
