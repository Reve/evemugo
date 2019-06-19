package main

import (
	"flag"
	"net"
)

var host = flag.String("host", "0.0.0.0", "ip address to listen on")
var port = flag.String("port", "26000", "port to listen on")

func main() {
	flag.Parse()

	ln, err := net.Listen("tcp", *host+":"+*port)

	if err != nil {
		println(err.Error())
	}

	for {
		conn, err := ln.Accept()

		if err != nil {
			println(err.Error())
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// TODO Handle connection
}
