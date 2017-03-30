package main

import (
	"net/http"
	"fmt"
	"log"
	"sync"
)

var mutex sync.Mutex

var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

func handler(writer http.ResponseWriter, reader *http.Request) {
	// %q 带双引号的字符串“abc” or 带单引号的字符‘x’
	fmt.Fprintf(writer, "URL.PATH = %q, URL.METHOD = %s \n", reader.URL.Path, reader.Method)
	mutex.Lock()
	count++
	mutex.Unlock()
}

func counter(writer http.ResponseWriter, read *http.Request) {
	mutex.Lock()
	fmt.Fprintf(writer, "current count: %d.", count)
	mutex.Unlock()
}