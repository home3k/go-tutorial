package main

import (
	"fmt"
	"runtime"
)

func main()  {
	fmt.Println(runtime.GOOS, runtime.GOARCH, runtime.NumGoroutine(), runtime.NumCPU())
}
