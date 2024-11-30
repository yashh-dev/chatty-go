package main

import (
	"chatty-go/irc"
	"fmt"
	"os"
)

var version = "0.0.0"

func main() {
	server := irc.NewIRCServer()
	if len(os.Args) < 2 {
		fmt.Println(HELP)
		os.Exit(0)
	}

	switch os.Args[1] {
	case "start":
		irc.Start(server)
	case "connect":
		irc.Connect()
	case "help":
		fmt.Println(HELP)
	default:
		fmt.Println(HELP)
	}
}

const HELP = `Welcome to IRC Chat!
You're now connected to [ServerName].
Type /help for a list of available commands.
Here are some common commands to get started:
1. Join a chat room: /join #channel_name, Example: /join #general
2. Leave a chat room: /part #channel_name, Example: /part #general
3. Send a private message to a user: /msg username message, Example: /msg Alice Hi Alice, how are you?
4. Change your nickname: /nick new_nickname, Example: /nick CoolCat
5. List available chat rooms: /list
6. Quit the server: /quit [optional_message] Example: /quit See you later!
Happy chatting! For more help, visit our documentation or type /help.`
