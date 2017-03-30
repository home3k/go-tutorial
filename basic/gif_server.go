package main

import (
	"net/http"
	"log"
	. "./gif_module"

)

func main()  {
	http.HandleFunc("/", handleGif)
	log.Fatal(http.ListenAndServe("localhost:8888", nil))
}

func handleGif(w http.ResponseWriter, r *http.Request)  {
	Lissajous(w)
}
