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

	fmt.Println("Connected to server")

	go readMessages(conn)

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		input := scanner.Text()

		if len(input) > 0 {
			switch input {
			case "custom;foo":
				message := Message{Type: "custom", Value: CustomMessage{FieldA: "foo", FieldB: "bar"}}
				asJson, _ := json.Marshal(message)
				conn.Write(append(asJson, '\n'))
			default:
				conn.Write(append([]byte(input), '\n'))
			}
		}
	}
}

func readMessages(conn net.Conn) {
	defer func() {
		conn.Close()
		os.Exit(1)
	}()

	for {
		msg, err := bufio.NewReader(conn).ReadString('\n')

		switch {
		case err == io.EOF:
			fmt.Println("Server connection terminated")
			return
		case err != nil:
			fmt.Printf("An error occurred while reading from server %v\n", err)
			return
		}

		fmt.Print(msg)
	}
}
