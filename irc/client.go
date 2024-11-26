package irc

import (
	"fmt"
	"net/url"
	"os"

	"github.com/gorilla/websocket"
)

type IRCclient struct {
}

// url := url.URL{Scheme: "ws", Host: HOSTNAME, Path: PATH}
func Connect() {
	fmt.Println("trying to ping....")
	dialer := websocket.DefaultDialer
	wsUrl := url.URL{Scheme: "ws", Host: "localhost:6697"}
	conn, _, err := dialer.Dial(wsUrl.String(), nil)
	if err != nil {
		fmt.Fprint(os.Stderr, err.Error())
		os.Exit(1)
	}

	defer conn.Close()

	// err = conn.WriteMessage(1, []byte("hello guys"))
	// if err != nil {
	// 	fmt.Fprint(os.Stderr, err.Error())
	// 	os.Exit(1)
	// }

	for {
		message := []byte{}
		_, err := fmt.Scanln(&message)
		if err != nil {
			panic(err)
		}
		err = conn.WriteMessage(1, message)
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			os.Exit(1)
			//TODO: reconnect if connection issues
		}
	}
}
