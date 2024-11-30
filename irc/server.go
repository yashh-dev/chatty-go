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
	rooms map[string]*Room
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func NewIRCServer() *IRCServer {
	return &IRCServer{
		rooms: make(map[string]*Room),
	}
}

func Start(server *IRCServer) {

	http.HandleFunc("/", server.HTTPHandler)

	fmt.Fprintf(os.Stdout, "irc server starting on %s:%s...\n", HOSTNAME, PORT)

	if err := http.ListenAndServe(fmt.Sprintf("%s:%s", HOSTNAME, PORT), nil); err != nil {
		fmt.Fprint(os.Stderr, err.Error())
	}
}

func (s IRCServer) JoinRoom(name string, member *Member) {
	room, ok := s.rooms[name]
	if !ok {
		room = NewRoom(name)
		s.rooms[name] = room
	}
	room.Join(member)
}
