package main

import (
	"fmt"
	"github.com/SanduCondorache/Interpreter/internals/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! This is the Loh programming language!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
