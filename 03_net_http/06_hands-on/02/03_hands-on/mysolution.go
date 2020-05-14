package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
)

func main()  {
	li, err:=net.Listen("tcp",":8080")
	if err!=nil{
		log.Fatalln(err)
	}
	defer li.Close()

	var conn net.Conn;
	for {
		conn, err = li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		scanner:=bufio.NewScanner(conn)
		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println(ln)
			io.WriteString(conn, "I see you connected. And Wrote: "+ln+"\n")
			if ln=="exit" {
				conn.Close()
			}
		}
		conn.Close()
		fmt.Println("Code got here.")
	}
	defer conn.Close()
}
