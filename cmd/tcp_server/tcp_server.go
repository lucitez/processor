package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
)

type Message struct {
	Type  string
	Value any
}

type CustomMessage struct {
	FieldA string
	FieldB string
}

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

			dec := json.NewDecoder(c)

			for {
				m := Message{}
				if err := dec.Decode(&m); err != nil {
					fmt.Printf("Error decoding message %v\n", err)
					return
				}

				fmt.Printf("%v\n", m.Value)

				switch m.Type {
				case "echo":
					out := []byte(m.Value.(string))
					c.Write(append(out, '\n'))
				case "custom":
					customMessage := CustomMessage(m.Value.(map[string]interface{}))
					fmt.Printf("FieldA is %s, FieldB is %s\n", customMessage.FieldA, customMessage.FieldB)
				default:

					if m.Value == "close" {
						return
					}

					if m.Value == "terminate" {
						listener.Close()
						return
					}
				}
			}
		}(conn)
	}
}
