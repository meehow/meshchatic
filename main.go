package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"meshchatic/messagehub"
	"net/http"
	"os"
	"strings"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/gorilla/websocket"
)

var (
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     checkOrigin,
	}
	hub = messagehub.New()
)

func main() {
	hostname, _ := os.Hostname()
	clientID := fmt.Sprintf("meshchatic-%x", md5.Sum([]byte(hostname)))
	opts := mqtt.NewClientOptions().AddBroker("tcp://mqtt.meshtastic.org:1883").SetClientID(clientID)
	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	log.Println("Connected")
	if token := c.Subscribe("msh/+/json/#", 0, mqttHandler); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}
	log.Println("Subscribed")
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe(":1985", nil)
}

func checkOrigin(r *http.Request) bool {
	return true
}

func mqttHandler(client mqtt.Client, msg mqtt.Message) {
	log.Println("Received message", msg.Topic())
	topic := strings.Split(msg.Topic(), "/")
	if len(topic) < 6 {
		return
	}
	hub.Broadcast <- messagehub.Message{
		Channel: topic[3],
		NodeID:  topic[4],
		App:     topic[5],
		Payload: msg.Payload(),
	}
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print(err)
		return
	}
	log.Println("New client", conn.RemoteAddr())
	hub.Register <- conn
}
