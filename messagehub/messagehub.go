package messagehub

import (
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
	nodeinfos    map[string]*Message
	positions    map[string]*Message
	textMessages []*Message
}

type Message struct {
	ChannelId    string `json:"channel_id"`
	GatewayId    string `json:"gateway_id"`
	App          string `json:"app"`
	WantResponse bool   `json:"want_response,omitempty"`
	Dest         uint32 `json:"dest,omitempty"`
	Source       uint32 `json:"source,omitempty"`
	RequestId    uint32 `json:"request_id,omitempty"`
	ReplyId      uint32 `json:"reply_id,omitempty"`
	Emoji        uint32 `json:"emoji,omitempty"`
	Payload      any    `json:"payload"`
}

func New() *Hub {
	hub := &Hub{
		Broadcast:    make(chan mqtt.Message),
		Register:     make(chan *websocket.Conn),
		clients:      make(map[*websocket.Conn]struct{}),
		nodeinfos:    make(map[string]*Message),
		positions:    make(map[string]*Message),
		textMessages: make([]*Message, 0, historyLimit),
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
			message := &Message{
				ChannelId: envelope.ChannelId,
				GatewayId: envelope.GatewayId,
			}
			switch payload := envelope.Packet.PayloadVariant.(type) {
			case *generated.MeshPacket_Decoded:
				message.App = payload.Decoded.Portnum.String()
				message.WantResponse = payload.Decoded.WantResponse
				message.Dest = payload.Decoded.Dest
				message.Source = payload.Decoded.Source
				message.RequestId = payload.Decoded.RequestId
				message.ReplyId = payload.Decoded.ReplyId
				message.Emoji = payload.Decoded.Emoji
				fmt.Printf("%s %s %s %s\n", message.App, msg.Topic(), envelope.ChannelId, envelope.GatewayId)
				switch payload.Decoded.Portnum {
				case generated.PortNum_NODEINFO_APP:
					nodeinfo := new(generated.User)
					err = proto.Unmarshal(payload.Decoded.Payload, nodeinfo)
					message.Payload = nodeinfo
					h.nodeinfos[nodeinfo.Id] = message
				case generated.PortNum_POSITION_APP:
					position := new(generated.Position)
					err = proto.Unmarshal(payload.Decoded.Payload, position)
					message.Payload = position
					h.positions[envelope.GatewayId] = message
				case generated.PortNum_TEXT_MESSAGE_APP:
					text := string(payload.Decoded.Payload)
					if rangetest.MatchString(text) {
						continue
					}
					message.Payload = text
					if len(h.textMessages) >= historyLimit {
						h.textMessages = append(h.textMessages[1:], message)
					} else {
						h.textMessages = append(h.textMessages, message)
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

func (h *Hub) History() []*Message {
	messages := make([]*Message, 0, len(h.nodeinfos)+len(h.positions)+len(h.textMessages))

	for _, message := range h.nodeinfos {
		messages = append(messages, message)
	}

	for _, message := range h.positions {
		messages = append(messages, message)
	}

	return append(messages, h.textMessages...)
}
