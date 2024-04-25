package main

import (
	"bufio"
	"context"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"
)

type message struct {
	mType string
	value string
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

	enc := gob.NewEncoder(conn)

	defer conn.Close()

	for {
		// pipe stdin to tcp connection
		if scanner.Scan() {
			if err := enc.Encode(message{"bing", scanner.Text()}); err != nil {
				log.Fatal(err)
			}
		}

		status, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			if err == io.EOF {
				fmt.Println("Connection terminated")
			} else {
				fmt.Printf("An error occurred while reading from server %v\n", err)
			}
			return
		}

		fmt.Println(status)
	}
}
