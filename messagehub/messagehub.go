package messagehub

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

const historyLimit = 256

var topics = []string{"NODEINFO_APP", "POSITION_APP", "TEXT_MESSAGE_APP"}

type Hub struct {
	clients   map[*websocket.Conn]struct{}
	Broadcast chan Message
	Register  chan *websocket.Conn
	history   map[string][]json.RawMessage
}

type Message struct {
	Topic   string          `json:"topic"`
	Payload json.RawMessage `json:"payload"`
}

func New() *Hub {
	hub := &Hub{
		Broadcast: make(chan Message),
		Register:  make(chan *websocket.Conn),
		clients:   make(map[*websocket.Conn]struct{}),
		history:   make(map[string][]json.RawMessage),
	}
	go hub.run()
	return hub
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.Register:
			h.clients[client] = struct{}{}
			for _, topic := range topics {
				for _, payload := range h.history[topic] {
					client.WriteJSON(Message{
						Topic:   topic,
						Payload: payload,
					})
				}
			}
		case message := <-h.Broadcast:
			if len(h.history[message.Topic]) >= historyLimit {
				h.history[message.Topic] = append(h.history[message.Topic][:1], message.Payload)
			} else {
				h.history[message.Topic] = append(h.history[message.Topic], message.Payload)
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
