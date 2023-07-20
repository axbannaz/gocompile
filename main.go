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
	dotfiles := os.Getenv("dotfilesdir")
	dotfiles = strings.ReplaceAll(dotfiles, "\\", "/")
	fmt.Printf("dotfiles: %s\n", dotfiles)
	if len(dotfiles) == 0 {
		df, err := os.UserHomeDir()
		if err != nil {
			dotfiles = "~/.files"
		} else {
			dotfiles = df + "/.files"
		}
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

func print(stdout io.ReadCloser, err bool) {
	r := bufio.NewReader(stdout)
	for {
		line, _, errr := r.ReadLine()
		if errr != nil {
			break
		}
		fd := os.Stdout
		if err {
			fd = os.Stderr
		}
		fmt.Fprintf(fd, "%s\n", line)
	}
}
