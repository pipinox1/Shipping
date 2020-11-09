package publisher

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

type Publisher interface {
	PublishMessage(message interface{},queue string)
}

type publisher struct {
	*amqp.Channel
}

func NewPublisher(channel *amqp.Channel) Publisher {
	return &publisher{channel}
}

func (p *publisher) PublishMessage(message interface{},queue string) {
	ch := p.Channel
	q, err := ch.QueueDeclare(
		"shipment", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "Failed to declare a queue")

	if err != nil {
		panic("BAD MESSAGE")
	}
	bodyOfBytes, err := json.Marshal(message)
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:       bodyOfBytes ,
		})
	log.Printf(" [x] Sent %s", bodyOfBytes)
	failOnError(err, "Failed to publish a message")
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}
