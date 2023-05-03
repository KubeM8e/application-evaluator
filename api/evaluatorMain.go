package main

import (
	"application-evaluator/api/handlers"
	"application-evaluator/models"
	"application-evaluator/pkg/evaluators"
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

func Evaluate() {
	// uses RabbitMQ for messaging
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		log.Printf("Could not connect to RabbitMQ: %s", err)
	}
	defer conn.Close()

	// creates a channel for publishing messages
	ch, err := conn.Channel()
	if err != nil {
		log.Printf("Could not open channel: %s", err)
	}
	defer ch.Close()

	// upload API publishes the message and this consume
	messages, err := ch.Consume(handlers.UploadSourcecodeQueue, "", true, false, false, false, nil)
	if err != nil {
		log.Printf("Could not consume message: %s", err)
	}

	for msg := range messages {
		var projectSource models.ProjectSource
		errP := json.Unmarshal(msg.Body, &projectSource)
		evaluators.EvaluateSourcecode(projectSource)
		if errP != nil {
			log.Printf("Could not unmarshal project source: %s", errP)
		}
	}

	if _, errPurge := ch.QueuePurge(handlers.UploadSourcecodeQueue, false); err != nil {
		log.Printf("Could not clear queue: %s", errPurge)
	}

}
