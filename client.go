package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:12345")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	consoleScanner := bufio.NewScanner(os.Stdin)

	go readLoop(conn)

	for consoleScanner.Scan() { // รอรับ input console
		fmt.Fprintln(conn, consoleScanner.Text()) // ส่งให้ server
	}
}

func readLoop(conn net.Conn) {
	scanner := bufio.NewScanner(conn) // รอ server ตอบกลับ
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
