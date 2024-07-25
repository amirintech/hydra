package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/amirintech/hydra-compiler/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hey, %s! This is Hydra's Programming Language REPL!\n", user.Username)
	repl.Run(os.Stdin, os.Stdout)
}
