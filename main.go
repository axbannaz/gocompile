package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	dotfiles := os.Getenv("dotfilesdir")
	fmt.Printf("dotfiles: %s\n", dotfiles)
	if len(dotfiles) == 0 {
		df, err := os.UserHomeDir()
		if err != nil {
			dotfiles = "~\\.files"
		} else {
			dotfiles = df + "\\.files"
		}
	}

	command := dotfiles + "\\bin\\go-install"
	bash, err := exec.LookPath("bash")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("exec: %s %s\n", bash, command)

	cmd := exec.Command(bash, command)
	err = cmd.Start()
	if err != nil {
		fmt.Println(err)
		os.Exit(254)
	}

	err = cmd.Wait()

	if err != nil {
		fmt.Println(err)
		os.Exit(255)
	}
}
