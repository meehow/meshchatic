package messagehub

import (
	"encoding/json"
	"fmt"
	"log"
	"meshchatic/github.com/meshtastic/go/generated"
	"regexp"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
)

const historyLimit = 256

var rangetest = regexp.MustCompile(`^seq \d+$`)

type Hub struct {
	clients      map[*websocket.Conn]struct{}
	Broadcast    chan mqtt.Message
	Register     chan *websocket.Conn
	users        map[string]*generated.User
	positions    map[string]*generated.Position
	textMessages []string
}

type Message struct {
	Type    string          `json:"type"`
	Channel string          `json:"channel"`
	NodeID  string          `json:"nodeID"`
	Payload json.RawMessage `json:"payload"`
}

func New() *Hub {
	hub := &Hub{
		Broadcast:    make(chan mqtt.Message),
		Register:     make(chan *websocket.Conn),
		clients:      make(map[*websocket.Conn]struct{}),
		users:        make(map[string]*generated.User),
		positions:    make(map[string]*generated.Position),
		textMessages: make([]string, 0, historyLimit),
	}
	go hub.run()
	return hub
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.Register:
			h.clients[client] = struct{}{}
		case msg := <-h.Broadcast:
			envelope := new(generated.ServiceEnvelope)
			err := proto.Unmarshal(msg.Payload(), envelope)
			if err != nil {
				log.Println(err)
				continue
			}
			fmt.Printf("%s %s %s\n", msg.Topic(), envelope.ChannelId, envelope.GatewayId)
			switch payload := envelope.Packet.PayloadVariant.(type) {
			case *generated.MeshPacket_Decoded:
				switch payload.Decoded.Portnum {
				case generated.PortNum_NODEINFO_APP:
					user := new(generated.User)
					err = proto.Unmarshal(payload.Decoded.Payload, user)
					h.users[envelope.GatewayId] = user
				case generated.PortNum_POSITION_APP:
					position := new(generated.Position)
					err = proto.Unmarshal(payload.Decoded.Payload, position)
					h.positions[envelope.GatewayId] = position
				case generated.PortNum_TEXT_MESSAGE_APP:
					text := string(payload.Decoded.Payload)
					if rangetest.MatchString(text) {
						continue
					}
					if len(h.textMessages) >= historyLimit {
						h.textMessages = append(h.textMessages[1:], text)
					} else {
						h.textMessages = append(h.textMessages, text)
					}
				default:
					continue
				}
				if err != nil {
					fmt.Printf("%s %s %s %v\n", envelope.ChannelId, envelope.GatewayId, payload.Decoded, err)
				}
			case *generated.MeshPacket_Encrypted:
				// fmt.Printf("Encrypted: %#v\n", payload.Encrypted)
				// json.NewEncoder(os.Stdout).Encode(payload.Encrypted)
			default:
				log.Printf("PayloadVariant %T not implemented", payload)
				continue
			}
			for client := range h.clients {
				// err := client.WriteJSON(struct {
				// 	App     string
				// 	Payload any
				// }{
				// 	App: fmt.Sprintf("%s", envelope.Packet.PayloadVariant),
				// })
				err := client.WriteJSON(msg)
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
	return nil
	// messages := make([]Message, 0, len(h.nodeinfos)+len(h.positions)+len(h.textMessages))
	//
	//	for _, message := range h.nodeinfos {
	//		messages = append(messages, message)
	//	}
	//
	//	for _, message := range h.positions {
	//		messages = append(messages, message)
	//	}
	//
	// return append(messages, h.textMessages...)
}
