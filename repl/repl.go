package repl

import (
	"bufio"
	"fmt"
	"io"
)

const PROMPT = "> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		if line == "exit" {
			fmt.Fprintln(out, "Goodbye!")
			return
		}
		ProcessInput(line, out)
	}
}

func ProcessInput(input string, out io.Writer) {
	fmt.Fprintln(out, "Received input: "+input)
}
