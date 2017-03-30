package main

import (
	"bufio"
	"fmt"
	"os"
)

// go没有set
func main() {
	sets := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !sets[line] {
			sets[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v \n", err)
		os.Exit(1)
	}

}
