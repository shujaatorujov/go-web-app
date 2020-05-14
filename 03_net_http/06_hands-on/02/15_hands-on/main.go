package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func serve(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	i:=0
	var url, meth string
	for scanner.Scan() {
		ln := scanner.Text()
		if i==0 {
			urls:=strings.Fields(ln)
			meth=urls[0]
			url=urls[1]
			fmt.Println("Request method: "+meth)
			fmt.Println("Request url: "+url)
		}
		if ln == "" {
			break
		}
		i++
	}
	body:=`<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Code Gangsta</title>
		</head>
		<body>
			<h1>`+`Request url is: `+url+` Request method is: `+meth+`</h1>
		</body>
		</html>`
c
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	io.WriteString(conn, "\r\n")
	io.WriteString(conn, body)
}

func main() {

	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err!=nil{
			log.Println(err)
			continue
		}
		serve(conn)
	}
}
