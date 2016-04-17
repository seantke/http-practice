package main

import (
	"io"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		// handle error
	}
	defer ln.Close()
	for {
		// Wait for a connection.
		conn, err := ln.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// Handle the connection in a new goroutine.
		// The loop then returns to accepting, so that
		// multiple connections may be served concurrently.
		go handleConnection(conn)
	}
}

func handleConnection(c net.Conn) {
	// Echo all incoming data.
	io.Copy(c, c)
	// Shut down the connection.
	c.Close()
}
