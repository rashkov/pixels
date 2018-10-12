package main

import (
	//"fmt"
	"golang.org/x/net/websocket"
	"log"
	"net/http"
	"os"
)

// @TODO: Add a keepalive heartbeat using time.Tick in a goroutine
//        Add a type field to ws frames -- "string" or "json"

// Fun with messages
// messages := make(chan string)
// go func() { messages <- "ping" }()
// go func() {
// 	msg := <-messages
// 	fmt.Println(msg)
// }()
// // Pause for input, otherwise program exits before we can print
// var input string
// fmt.Scanln(&input)

var (
	logger *log.Logger
)

type message struct {
	// the json tag means this will serialize the
	// value of Message with the key of "message" instead of "Message"
	Message string `json:"message"`
}

func main() {
	logger = log.New(os.Stdout, "web ", log.LstdFlags)

	server := &http.Server{
		Addr:    ":8080",
		Handler: routes(),
	}

	server.ListenAndServe()
}

func routes() *http.ServeMux {
	r := http.NewServeMux()

	r.HandleFunc("/hello", hello)

	// Can't use HandleFunc for websockets. Use Handle()
	//r.HandleFunc("/ws", websocket.Handler(wsSocket))
	r.Handle("/ws", websocket.Handler(wsSocket))

	return r
}

func hello(w http.ResponseWriter, r *http.Request) {
	logger.Println("this is the earth")
}

func wsSocket(ws *websocket.Conn) {
	for {
		// allocate our container struct
		var m message

		// receive a message using the codec
		if err := websocket.JSON.Receive(ws, &m); err != nil {
			log.Println(err)
			break
		}

		log.Println("Received message:", m.Message)

		// send a response
		m2 := message{"Thanks for the message!"}
		if err := websocket.JSON.Send(ws, m2); err != nil {
			log.Println(err)
			break
		}
	}
}
