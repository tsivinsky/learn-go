package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func broadcastMessage(conns []net.Conn, currentConn net.Conn, data []byte) {
	for _, conn := range conns {
		if conn.RemoteAddr().String() == currentConn.RemoteAddr().String() {
			continue
		}

		_, err := conn.Write(data)
		if err != nil {
			continue
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatal(err)
	}
	defer ln.Close()

	var conns []net.Conn

	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		defer conn.Close()

		go func(conn net.Conn) {
			for {
				r := bufio.NewReader(conn)
				s, _ := r.ReadString('\n')

				for _, c := range conns {
					if c.RemoteAddr().String() == conn.RemoteAddr().String() {
						continue
					}

					fmt.Fprintf(c, "%s: %s", conn.RemoteAddr().String(), s)
				}
			}
		}(conn)

		fmt.Printf("%s connected\n", conn.RemoteAddr().String())

		conns = append(conns, conn)
	}
}
