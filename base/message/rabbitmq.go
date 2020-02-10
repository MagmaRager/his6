package message

import (
	"fmt"
	"his6/base/config"
	"log"

	"github.com/streadway/amqp"
)

var (
	on   bool = false
	conn *amqp.Connection
	ch   *amqp.Channel
)

func init() {
	createConnAndChannel()
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Println("%s: %s", msg, err) //.Fatalf
	}
}

//Send 发送rabbitmq消息
func Send(queue string, msg string) {

	if !on {
		return
	}
	/*
		q, err := ch.QueueDeclare(
			queue, // name
			true,  // durable
			false, // delete when unused
			false, // exclusive
			false, // no-wait
			nil,   // arguments
		)
		failOnError(err, "Failed to declare a queue")
	*/

	err := ch.Publish(
		"",    // exchange
		queue, // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(msg),
		})
	//log.Printf(" [x] Sent %s", msg)
	if err != nil {
		failOnError(err, "Failed to publish a message")
		createConnAndChannel()
	}
}

func createConnAndChannel() {
	url := config.GetConfigString("mq", "url", "")
	if url == "" {
		return
	}
	port := config.GetConfigInt("mq", "port", 5672)
	user := config.GetConfigString("mq", "user", "guest")
	pwd := config.GetConfigString("mq", "password", "guest")

	on = true

	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", user, pwd, url, port))
	failOnError(err, "Failed to connect to RabbitMQ")

	ch, err = conn.Channel()
	failOnError(err, "Failed to open a channel")
}
