package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main (){
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
		go handle(conn)
	}
}

func handle (conn net.Conn) {
	defer conn.Close()
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line
			m := strings.Fields(ln)[0] // method
			u := strings.Fields(ln)[1] // uri
			if u == "/mike" {
				respondMike(conn)
			} else if u == "/winning" {
				respondWinning(conn)
			} else {
				respondHome(conn)
			}
			fmt.Println("***METHOD", m)
			fmt.Println("***URI", u)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
}

func respondMike(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Mike Route</strong></body></html>`
	setHeaders(conn, body)
}

func respondWinning(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Winning Route</strong></body></html>`
	setHeaders(conn, body)
}

func respondHome(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World From Home Route</strong></body></html>`
	setHeaders(conn, body)
}

func setHeaders(conn net.Conn, body string) {
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}