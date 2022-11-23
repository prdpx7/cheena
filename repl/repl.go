package repl

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

const PROMPT = "> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprint(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := strings.TrimSpace(strings.ToLower(scanner.Text()))
		if line == "exit" || line == "quit" {
			fmt.Fprintln(out, "Goodbye!")
			return
		}
		ProcessInput(line, out)
	}
}

func ProcessInput(input string, out io.Writer) {
	fmt.Fprintln(out, "Received input: "+input)
}
