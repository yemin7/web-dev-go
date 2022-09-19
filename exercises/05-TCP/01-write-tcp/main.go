package main

import (
	"io"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		io.WriteString(conn, "Hello from TCP Server ...")
		conn.Close()
	}
}
