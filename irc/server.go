package irc

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var HOSTNAME string = "localhost"
var PORT string = "6697"

type IRCServer struct {
	wconn *websocket.Conn
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func Start(server *IRCServer) {

	http.HandleFunc("/", server.HTTPHandler)

	fmt.Fprintf(os.Stdout, "irc server starting on %s:%s...\n", HOSTNAME, PORT)

	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", HOSTNAME, PORT), nil); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
	}
}
