package main

import (
	"fmt"
	"strconv"
)

func countAndSay(n int) string {
	s :="1"
	if n == 1 {
		return s
	}
	for i:=1;i<n;i++{
		s = count(s)
	}
	return s
}

func count(s string) string {
	var d rune
	var c=0
	result := ""
	for _, item := range s {
		if item!=d {
			if c != 0 {
				result = result + strconv.Itoa(c) + string(d)
			}
			d = item
			c = 1
		} else {
			c++
		}
	}
	return result + strconv.Itoa(c) + string(d)
}

func main()  {
	fmt.Printf("%s\n", countAndSay(1))
	fmt.Printf("%s\n", countAndSay(2))
	fmt.Printf("%s\n", countAndSay(3))
	fmt.Printf("%s\n", countAndSay(4))
}


