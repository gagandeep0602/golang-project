package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/websocket"
)

type Display struct {
	Rollno         int
	Totalwords     int16
	Totalcharacter int
	Wordsminute    int
}

var y int

//read and write buffer size upgrade
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	// upgrade this connection to a WebSocket
	// connection
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	log.Println("Client Connected")

	reader(ws)
}

// ... Use conn to send and receive messages.

func reader(conn *websocket.Conn) {
	for {
		message, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)

			return
		}

		s := len(string(p))
		if s > 1 {
			g := len(string(p))

			fmt.Println(g) // prints the length of the text

		}

		fmt.Println(string(p))

		if err := conn.WriteMessage(message, p); err != nil {

			log.Println(err)
			return
		}

	}

}

// display of teacher//

func agg(w http.ResponseWriter, r *http.Request) {
	p := Display{Rollno: 1,
		Totalwords:     1,
		Totalcharacter: y,
		Wordsminute:    5}
	t, er := template.ParseFiles("basic.html")
	if er != nil {
		fmt.Println(er)
	}
	t.Execute(w, p)
}

func setupRoutes() {
	http.HandleFunc("/", agg)

	http.HandleFunc("/ws", wsEndpoint)

}

func main() {

	fmt.Println("Hello World")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":5500", nil))

}

// everything you can see in your terminal also
