package irc

import (
	"fmt"

	"github.com/gorilla/websocket"
)

type Room struct {
	name    string
	channel chan string
	members []chan string
}

func NewRoom(name string) *Room {
	room := &Room{
		name:    name,
		channel: make(chan string),
	}
	go func() {
		for {
			msg := <-room.channel
			for _, member := range room.members {
				fmt.Printf("message recieved at %v : %v\n", room.name, msg)
				member <- msg
			}
		}
	}()

	return room
}

func (r *Room) Join(member *Member) {
	r.members = append(r.members, member.channel)
	fmt.Println(r.name, " has another member")
}

type Member struct {
	channel chan string
	room    string
	name    string
	conn    *websocket.Conn
}

func NewMember(conn *websocket.Conn) *Member {
	member := &Member{conn: conn, channel: make(chan string)}

	go func() {
		for msg := range member.channel {
			err := conn.WriteMessage(1, []byte(msg))
			if err != nil {
				fmt.Println(err)
			}
		}
	}()

	return member
}
