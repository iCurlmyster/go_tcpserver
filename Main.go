package main

import (
	"./src/listener"
	"log"
	"net"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "8080"
	CONN_TYPE = "tcp"
)

func main() {
	l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	log.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	// spin off group listener
	listener.GetWorldInstance()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Println("Acceptance error: " + err.Error())
		}
		go listener.UserListenerLoop(conn)
	}
}
