package message

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"his6/base/config"
)

var(
	connNats *nats.Conn
	subNats *nats.Subscription
)

type MessageData struct {
	Code string
	Info string
	SendIP string
	SendEmpId int
}

func init(){
	// Connect to a server
	natsUrl := config.GetConfigString("nats", "url", "nats://127.0.0.1:4222")

	connNats, _ = nats.Connect(natsUrl)

	// Simple Publisher
	//connNats.Publish("foo", []byte("Hello World"))

	//// Simple Async Subscriber
	//nc.Subscribe("foo", func(m *nats.Msg) {
	//	fmt.Printf("Received a message: %s\n", string(m.Data))
	//})
	//
	//// Responding to a request message
	//nc.Subscribe("request", func(m *nats.Msg) {
	//	m.Respond([]byte("answer is 42"))
	//})
	//
	//// Simple Sync Subscriber
	//sub, _ := nc.SubscribeSync("foo")
	//m, err := sub.NextMsg(timeout)
	//
	//// Channel Subscriber
	//ch := make(chan *nats.Msg, 64)
	//sub, err := nc.ChanSubscribe("foo", ch)
	//msg := <- ch
	//
	//// Unsubscribe
	//sub.Unsubscribe()
	//
	//// Drain
	//sub.Drain()
	//
	//// Requests
	//msg, err := nc.Request("help", []byte("help me"), 10*time.Millisecond)
	//
	//// Replies
	//nc.Subscribe("help", func(m *nats.Msg) {
	//	nc.Publish(m.Reply, []byte("I can help!"))
	//})

	//// Drain connection (Preferred for responders)
	//// Close() not needed if this is called.
	//connNats.Drain()
	//
	//// Close connection
	//connNats.Close()
}

func Publish(topic, content string) {
	md := MessageData{}
	md.Code = topic
	md.Info = content
	rst, _ := json.Marshal(md)
	connNats.Publish(topic, rst)
}

func Subscribe(topic string) {
	subNats, _ = connNats.Subscribe(topic, func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
}

func Unsubscribe() {
	subNats.Unsubscribe()
}