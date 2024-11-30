package irc

import (
	"bufio"
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

	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(string(message))
		}
	}()

	for {

		ioReader := bufio.NewReader(os.Stdin)
		message, err := ioReader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}

		if len(message) < 1 {
			continue
		}

		fmt.Printf("message entered: %v.\n", message)
		err = conn.WriteMessage(1, []byte(message))
		if err != nil {
			fmt.Fprint(os.Stderr, err.Error())
			os.Exit(1)
			//TODO: reconnect if connection issues
		}
	}
}
