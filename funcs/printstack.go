package main

import (
	"runtime"
	"os"
	"fmt"
)

func PrintStack()  {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func f(value int)  {
	fmt.Println(value)
}

func main()  {
	defer PrintStack()
	f(1)
}
