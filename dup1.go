package main

import "fmt"
import "bufio"
import "os"

func main()  {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}
	for line, count := range counts {
		fmt.Printf("%d\t%s\n", count, line)
	}
}
