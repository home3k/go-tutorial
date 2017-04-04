package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8010")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)

	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		if _, err := io.WriteString(conn, time.Now().Format("2006-01-02 15:04:05\n")); err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
