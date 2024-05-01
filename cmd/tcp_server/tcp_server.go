package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
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

	fmt.Println("Listening on port 8000")

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

			reader := bufio.NewReader(c)

			for {
				message, err := reader.ReadString('\n')
				switch {
				case err == io.EOF:
					fmt.Println("Client connection terminated")
					return
				case err != nil:
					fmt.Printf("An error occurred while reading from client %v\n", err)
					return
				}

				msgType, msgVal := extractMsg(message)

				fmt.Printf("Message type: %v\n", strings.TrimSpace(msgType))

				// Then handle the message according to its type
				switch msgType {
				case "custom":
					m := Message{}
					json.Unmarshal([]byte(msgVal), &m)
					fmt.Printf("custom is %v\n", m)
				case "close":
					return
				case "terminate":
					listener.Close()
					return
				default:
					c.Write([]byte("echo: " + message + "\n"))
				}
			}
		}(conn)
	}
}

func extractMsg(msg string) (string, string) {
	splitMsg := strings.Split(msg, ";")

	msgType := splitMsg[0]

	msgVal := ""

	if len(splitMsg) > 1 {
		msgVal = splitMsg[1]
	}

	return msgType, msgVal
}
