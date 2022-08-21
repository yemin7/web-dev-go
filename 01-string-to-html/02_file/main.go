package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	name := "Hello World"
	str := fmt.Sprint(`
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	</html>
	`)

	file, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Error creating file", err)
	}
	defer file.Close()

	io.Copy(file, strings.NewReader(str))
}
