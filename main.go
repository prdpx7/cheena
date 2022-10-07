package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/prdpx7/sqlite/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is SQLite\n", user.Username)
	fmt.Println("Starting REPL...")
	repl.Start(os.Stdin, os.Stdout)
}
