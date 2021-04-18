package messagehub

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

const (
	historyLimit   = 256
	nodeinfoApp    = "NODEINFO_APP"
	positionApp    = "POSITION_APP"
	textMessageApp = "TEXT_MESSAGE_APP"
)

type Hub struct {
	clients      map[*websocket.Conn]struct{}
	Broadcast    chan Message
	Register     chan *websocket.Conn
	nodeinfos    map[string]Message
	positions    map[string]Message
	textMessages []Message
}

type Message struct {
	Channel string          `json:"channel"`
	NodeID  string          `json:"nodeID"`
	App     string          `json:"app"`
	Payload json.RawMessage `json:"payload"`
}

func New() *Hub {
	hub := &Hub{
		Broadcast:    make(chan Message),
		Register:     make(chan *websocket.Conn),
		clients:      make(map[*websocket.Conn]struct{}),
		nodeinfos:    make(map[string]Message),
		positions:    make(map[string]Message),
		textMessages: make([]Message, 0, historyLimit),
	}
	go hub.run()
	return hub
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.Register:
			h.clients[client] = struct{}{}
		case message := <-h.Broadcast:
			switch message.App {
			case nodeinfoApp:
				h.nodeinfos[message.NodeID] = message
			case positionApp:
				h.positions[message.NodeID] = message
			case textMessageApp:
				if len(h.textMessages) >= historyLimit {
					h.textMessages = append(h.textMessages[1:], message)
				} else {
					h.textMessages = append(h.textMessages, message)
				}
			default:
				log.Printf("App %q not implemented", message.App)
			}
			for client := range h.clients {
				err := client.WriteJSON(message)
				if err != nil {
					log.Println(err)
					delete(h.clients, client)
					client.Close()
				}
			}
		}
	}
}

func (h *Hub) History() []Message {
	messages := make([]Message, 0, len(h.nodeinfos)+len(h.positions)+len(h.textMessages))
	for _, message := range h.nodeinfos {
		messages = append(messages, message)
	}
	for _, message := range h.positions {
		messages = append(messages, message)
	}
	return append(messages, h.textMessages...)
}
