package main

import (
	"bufio"
	"net"
)

func main() {
	listener, _ := net.Listen("tcp", ":5000")

	for {
		conn, _ := listener.Accept()

		go func(conn net.Conn) {
			defer conn.Close()

			sc := bufio.NewScanner(conn)

			for sc.Scan() {
				if sc.Text() == "Bye" {
					return
				}

				if sc.Text() == "Ping" {
					conn.Write([]byte("Pong\n"))
				}
			}
		}(conn)
	}
}
