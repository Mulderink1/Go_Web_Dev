package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}
		go handle(conn, li)
	}
}

func handle(conn net.Conn, li net.Listener) {
	defer conn.Close()
	defer li.Close()
	response(conn)
}

func response(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		respond(conn)
		return
	}
}

func respond(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Your Connected</strong></body></html>`
	setHeaders(conn, body)
}

func setHeaders(conn net.Conn, body string) {
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}