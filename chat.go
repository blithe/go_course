package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		connection, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}

		go handleConn(connection)
	}
}


type client chan<- string
var (
	entering = make(chan client)
	leaving = make(chan client)
	messages = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
			case msg := <-messages:
				for cli := range clients {
					cli <- msg
				}
			case cli := <-entering:
				clients[cli] = true
			case cli := <-leaving:
				delete(clients, cli)
				close(cli)
		}
	}
}

func handleConn(connection net.Conn) {
	channel := make(chan string)
	go clientWriter(connection, channel)
	who := connection.RemoteAddr().String()
	channel <- "You are " + who
	messages <- who +  " has arrived"
	entering <- channel

	input := bufio.NewScanner(connection)
	for input.Scan() {
		messages <- who + ": " + input.Text()
	}

	leaving <- channel
	messages <- who + " has left"
	connection.Close()
}

func clientWriter(connection net.Conn, channel <-chan string) {
	for message := range channel {
	fmt.Fprintln(connection, message)
	}
}
