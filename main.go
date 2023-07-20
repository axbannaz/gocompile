package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	dotfiles := os.Getenv("dotfilesdir")
	if len(dotfiles) == 0 {
		df, err := os.UserHomeDir()
		if err != nil {
			dotfiles = "~/.files"
		} else {
			dotfiles = df + "/.files"
		}
	}

	command := dotfiles + "/bin/go-install"
	fmt.Printf("exec: %v\n", command)
	cmd := exec.Command(command)
	err := cmd.Wait()

	if err != nil {
		fmt.Print(err)
		os.Exit(255)
	}
}
