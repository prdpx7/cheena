package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/prdpx7/cheena/repl"
	"github.com/prdpx7/cheena/server"
)

func main() {

	arg := os.Args[1]
	if arg == "server" {
		server.Server()
	} else if arg == "cli" {
		user, err := user.Current()
		if err != nil {
			panic(err)
		}
		fmt.Printf("Hello %s! This is cheena\n", user.Username)
		fmt.Println("Starting REPL...")
		repl.Start(os.Stdin, os.Stdout)
	}

}
