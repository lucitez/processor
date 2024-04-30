package main

import (
	"bufio"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
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
	var d net.Dialer
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	conn, err := d.DialContext(ctx, "tcp", "localhost:8000")
	if err != nil {
		log.Fatalf("Failed to dial: %v", err)
	}

	scanner := bufio.NewScanner(os.Stdin)

	enc := json.NewEncoder(conn)

	defer conn.Close()

	for {
		// pipe stdin to tcp connection
		if scanner.Scan() {
			input := scanner.Text()
			var message Message
			switch input {
			case "echo":
				message = Message{Type: "echo", Value: input}
			case "custom":
				message = Message{Type: "custom", Value: CustomMessage{FieldA: "foo", FieldB: "bar"}}
			default:
				message = Message{Type: "message", Value: input}
			}

			if err := enc.Encode(message); err != nil {
				log.Fatal(err)
			}
		}

		msg, err := bufio.NewReader(conn).ReadString('\n')
		switch {
		case err == io.EOF:
			fmt.Println("Connection terminated")
			return
		case err != nil:
			fmt.Printf("An error occurred while reading from server %v\n", err)
			return
		}

		msg = strings.Trim(msg, "\n")
		fmt.Println(msg)
	}
}
