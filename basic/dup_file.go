package main

import (
	"fmt"
	"os"
	"bufio"
)

func main()  {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	}
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			// %v 变量的自然形式
			fmt.Fprintf(os.Stderr, "dup file errors: %v\n", err)
			continue
		}
		countLines(f, counts)

	}
	for line, count := range counts {
		fmt.Printf("%d\t%s\n", count, line)
	}
}

func countLines(f *os.File, counts map[string] int)  {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
