package irc

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

func (srv *IRCServer) HTTPHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	member := NewMember(conn)

	for {
		_, reader, err := conn.NextReader()

		if err != nil {
			fmt.Println(err)
			continue
		}

		msgBytes, err := io.ReadAll(reader)

		if err != nil {
			fmt.Println(err)
			continue
		}

		if len(msgBytes) > 0 {
			// fmt.Println("read : ", count)
			srv.HandleMessage(msgBytes, conn, member)
		}

	}

}

func (s *IRCServer) HandleMessage(msg []byte, conn *websocket.Conn, member *Member) {
	message := string(msg)

	if len(msg) > 0 {
		if msg[0] == '/' {

			messageBlocks := strings.Split(message, " ")
			if len(messageBlocks) < 2 {
				conn.WriteMessage(1, []byte("enter room name sucker"))
				return
			}

			switch messageBlocks[0] {
			case "/join":
				roomName := messageBlocks[1]
				member.room = roomName
				s.JoinRoom(roomName, member)
			}

		} else {
			if member.room != "" {
				room := s.rooms[member.room]
				room.channel <- message
			}
		}
	}
}
