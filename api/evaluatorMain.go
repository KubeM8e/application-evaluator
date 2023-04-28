package main

import (
	"application-evaluator/api/handlers"
	"application-evaluator/pkg/evaluators"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func Evaluate() {
	fmt.Println("YAYYYYYY")
	// uses RabbitMQ for messaging
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Printf("Could not connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	// Create a channel for publishing messages
	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Could not open channel: %s", err)
	}
	defer ch.Close()

	messages, err := ch.Consume(handlers.UploadSourcecodeQueue, "", true, false, false, false, nil)
	if err != nil {
		log.Printf("Could not consume message: %s", err)
	}

	for msg := range messages {
		evaluators.EvaluateSourcecode(string(msg.Body))
	}

	if _, errPurge := ch.QueuePurge(handlers.UploadSourcecodeQueue, false); err != nil {
		log.Printf("Could not clear queue: %s", errPurge)
	}

}
