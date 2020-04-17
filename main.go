package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

var conns = make(map[string]net.Conn)

func main() {

	l, err := net.Listen("tcp", ":12345")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := l.Accept() // รอจนกว่ามีคน connect
		if err != nil {
			log.Fatal(err)
		}

		go Handler(conn)
	}

}

func Handler(conn net.Conn) {
	conns[conn.RemoteAddr().String()] = conn
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		input := scanner.Text()
		for _, c := range conns {
			fmt.Fprintln(c, input)
		}
	}
}
