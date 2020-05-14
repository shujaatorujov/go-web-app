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
	var url string
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i==0 {
			urls:=strings.Fields(ln)
			url=urls[1]
			fmt.Println("Request url: "+ln)
		}
		if ln == "" {
			break
		}
		i++
	}
	body := "Request url is: "+url
	io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/plain\r\n")
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
