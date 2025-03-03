package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

// ! im using https://dev.to/neelp03/using-websockets-in-go-for-real-time-communication-4b3l as reference.
var clients = make(map[*websocket.Conn]bool)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

// for what i understand, writeResp is a somewhat like object-class (type struct in golang) for making the response to send the client.
// like a OOP builder, but for functional programming?
func handler(writeResp http.ResponseWriter, req *http.Request) {
	conn, err := upgrader.Upgrade(writeResp, req, nil)

	clients[conn] = true // hashMapping client connection

	defer delete(clients, conn)
	defer conn.Close()

	if err != nil {
		fmt.Println("errorHasOccurred", err)
		return
	}

	for {
		// Read message from browser
		_, msg, err := conn.ReadMessage()

		if err != nil {
			fmt.Println("read error:", err)
			break
		}
		fmt.Printf("Received: %s\n", msg)

		// Write message back to browser
		msgServer := []byte("this is what the client gets")
		if err := conn.WriteMessage(websocket.TextMessage, msgServer); err != nil {
			fmt.Println("write error:", err)
			break
		}
	}

	fmt.Println("connection closed")

}

/*

	// from what i learnt, the keyword "go" (known as goroutine) is making the func async
	// using another thread (since this is a multithreaded language) and to wait them to finish
	// we use things like "waitGroups"
	go messageHandler(p)

	func messageHandler(message []byte) {
		fmt.Println(string(message))
		for conn := range clients {
			conn.WriteMessage(websocket.TextMessage, message)
		}
}

*/
