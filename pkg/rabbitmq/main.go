package rabbitmq

import (
	"fmt"
	"github.com/streadway/amqp"
	"os"
)

var Rabbitmq *amqp.Connection

func Init() {
	amqpServerURL := os.Getenv("URL_RABBITMQ")

	connectRabbitMQ, err := amqp.Dial(amqpServerURL)
	if err != nil {
		fmt.Println("Nao conectado ao rabbit")
	}

	Rabbitmq = connectRabbitMQ
}
