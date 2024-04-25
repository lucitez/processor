package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal(err)
	}

	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			if err == net.ErrClosed {
				fmt.Println("Connection closed")
				return
			}
			log.Fatal(err)
		}

		go func(c net.Conn) {
			defer c.Close()

			b := make([]byte, 1024)

			for {
				n, err := c.Read(b)

				if err != nil {
					fmt.Println(err)
					return
				}

				fmt.Printf("%s\n", b[:n])

				if string(b[:n]) == "close" {
					return
				}

				if string(b[:n]) == "terminate" {
					listener.Close()
					return
				}

				out := b[:n]
				out = append(out, '\n')
				c.Write(out)
			}
		}(conn)
	}
}
