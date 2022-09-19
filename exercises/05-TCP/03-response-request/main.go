package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	i := 0
	var m, url string
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)

		if i == 0 {
			m = strings.Fields(ln)[0]
			url = strings.Fields(ln)[1]
			fmt.Println("METHOD: ", m)
			fmt.Println("URL: ", url)
		}
		if ln == "" {
			fmt.Println("THIS IS THE END OF THE HTTP REQUEST HEADERS")
			break
		}
		i++
	}
	//	body := "CHECK OUT THE RESPONSE BODY PAYLOAD\n"
	//	body += m + "\n"
	//	body += url

	body := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<title>Code Gangsta</title>
	</head>
	<body>
		<h1>"HOLY COW THIS IS LOW LEVEL"</h1>
	</body>
	</html>
`

	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}
