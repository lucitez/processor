package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"

	"github.com/lucitez/processor/benchmarker"
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
			fmt.Println("Connection closed")
			os.Exit(1)
		}

		go handleConnection(listener, conn)
	}
}

func handleConnection(listener net.Listener, conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

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
		case "benchmark":
			handleBenchmark(conn, msgVal)
		case "close":
			return
		case "terminate":
			listener.Close()
			return
		default:
			conn.Write([]byte("echo: " + message + "\n"))
		}
	}
}

func handleBenchmark(conn net.Conn, url string) {
	b := benchmarker.New(url)

	conn.Write([]byte("Benchmarking " + url + "\n"))

	b.BenchmarkWebsite(func(p benchmarker.Performance) {
		asJson, _ := json.Marshal(p)
		conn.Write(append(asJson, '\n'))
	})

	conn.Write([]byte("Benchmarking complete\n"))
}

func extractMsg(msg string) (string, string) {
	splitMsg := strings.Split(msg, ";")

	msgType := splitMsg[0]

	msgVal := ""

	if len(splitMsg) > 1 {
		msgVal = splitMsg[1]
	}

	return strings.TrimSpace(msgType), strings.TrimSpace(msgVal)
}
