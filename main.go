package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "list requires a delimiter")
		return
	}

	delim := os.Args[1][0]

	stdin := bufio.NewReader(os.Stdin)

	for {
		line, err := stdin.ReadString(delim)
		if err != nil && err != io.EOF {
			fmt.Fprintln(os.Stderr, "Failed while reading from stdin: %s", err)
			return
		}

		if len(line) == 0 {
			return
		}

		fmt.Println(line)

		if err == io.EOF {
			return
		}
	}
}
