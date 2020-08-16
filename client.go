package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"pipetest/namedpipe"
)

type Message struct {
	Message string
}

func SendAndReceiveMessage() {
	fmt.Println("Connecting to server...")

	pipeName := "mytestpipe"
	pipeArgs := []string{}

	pipe, err := namedpipe.NewNamedPipeClient(pipeName, pipeArgs...)
	if err != nil {
		fmt.Printf("Error connecting to named pipe %s - %v", pipeName, err)
		os.Exit(1)
	}
	fmt.Println("Connected to server")

	var count int

	for {
		/*
			Uncomment below lines to receive messages from named-pipe server

			// reader, err := pipe.GetReader()
			// if err != nil {
			// 	fmt.Printf("Error in getting pipe reader %v", err)
			// 	os.Exit(1)
			// }
			// r := bufio.NewReader(reader)
			// msg, err := r.ReadString('\n')
			// if err != nil {
			// 	fmt.Printf("Unable to read: %+v\n", err)
			// 	return
			// }
			// fmt.Println("Message received: %s", msg)

		*/
		if count > 100 {
			break
		}

		writer, err := pipe.GetWriter()
		if err != nil {
			fmt.Printf("Error in getting pipe writer %v", err)
			os.Exit(1)
		}

		w := bufio.NewWriter(writer)
		_, err = w.WriteString("hello world\n")
		if err != nil {
			fmt.Printf("Error in writing to pipe %v", err)
			os.Exit(1)
		}
		fmt.Println("Wrote to pipe")
		err = w.Flush()
		if err != nil {
			fmt.Printf("Error in flushing pipe %v", err)
			os.Exit(1)
		}
		count = count + 1

		time.Sleep(5 * time.Second)
	}
}

func main() {
	fmt.Println("Created client")
	SendAndReceiveMessage()
}
