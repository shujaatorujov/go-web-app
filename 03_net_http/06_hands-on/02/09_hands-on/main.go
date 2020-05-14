package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func serve(conn net.Conn) {
	defer conn.Close()
	scanner := bufio.NewScanner(conn)
	fmt.Fprint(conn, "sucu>")
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		fmt.Fprintln(conn,"You wrote: "+ln)
		if ln == "exit" {
			break
		}
		fmt.Fprint(conn, "sucu>")
	}
	io.WriteString(conn, "Connection lost")
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
