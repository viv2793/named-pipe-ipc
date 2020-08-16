package main

import (
	// "bufio"

	"bufio"
	"fmt"
	"os"

	"pipetest/namedpipe"
)

type Message struct {
	Message string
}

func acceptPipeConnections(pipe namedpipe.NamedPipeServer) {
	reqId := "dummyreq"
	fmt.Println("Listening for pipe connections...")
	err := pipe.NewClient(reqId)
	if err != nil {
		fmt.Printf("Error: %+v", err)
		return
	}

	fmt.Println("New client connected...")
	/*
		Uncomment below lines to send messages to client

		// message := Message{
		// 	Message: "Hello, World from server !",
		// }
		// writer, err := pipe.GetWriter(reqId)
		// if err != nil {
		// 	fmt.Printf("Unable to get writer: %+v\n", err)
		// 	return
		// }
		// enc := gob.NewEncoder(writer)
		// err = enc.Encode(message)
		// if err != nil {
		// 	fmt.Printf("Unable to write: %+v\n", err)
		// 	return
		// }

		// fmt.Println("Sent message to client...")
	*/
	for {
		reader, err := pipe.GetReader(reqId)
		if err != nil {
			fmt.Printf("Unable to get reader: %+v", err)
			return
		}
		fmt.Println("got reader...")

		r := bufio.NewReader(reader)
		msg, err := r.ReadString('\n')
		if err != nil {
			// handle error
			fmt.Printf("error 1: %s\n", err)
			return
		}
		fmt.Printf("Received response: %+v", msg)
	}
}

func CreateServer() {
	fmt.Println("Attempting to create pipe")
	pipeName := "mytestpipe"
	pipe, err := namedpipe.NewNamedPipeServer(pipeName)
	if err != nil {
		fmt.Printf("Error: %+v\n", err)
		os.Exit(1)
	}
	fmt.Println("Created pipe")
	acceptPipeConnections(pipe)
}

func main() {
	CreateServer()
}
