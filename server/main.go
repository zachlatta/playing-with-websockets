package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var conns []*websocket.Conn

func main() {
	go serve()

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(`Welcome to the websockets demo!

Please enter coordinates of the circle separated by comma. Ex. 10,50

> `)
	for scanner.Scan() {
		for _, conn := range conns {
			go conn.WriteMessage(1, []byte(scanner.Text()))
		}

		fmt.Print("> ")
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Reading standard input:", err)
	}
}

func serve() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":4040", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(w, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		log.Println(err)
	}
	conns = append(conns, conn)
}
