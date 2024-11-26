package irc

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func (srv *IRCServer) HTTPHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
		}
		srv.HandleMessage(msg)
	}

}

func (*IRCServer) HandleMessage(msg []byte) {
	fmt.Println(string(msg))
	if len(msg) > 0 {
		if msg[0] == '/' {
			fmt.Println("yeah command")
		}
	}

}
