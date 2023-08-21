package mensageria

import (
	"encoding/json"
	"fmt"
	"github.com/Bhimmo/golang-simple-api/internal/infra/mensageria"
	"github.com/streadway/amqp"
)

type RabbitMq struct {
	connection *amqp.Connection
}

func NovoRabbitMq(conn *amqp.Connection) *RabbitMq {
	return &RabbitMq{
		connection: conn,
	}
}

func (m *RabbitMq) EnviarEmail(
	queueEnviar string,
	bodyMensagem mensageria.MensagemEnviarRabbitmq,
) {
	channel, errChanel := m.connection.Channel()
	erroRabbit(errChanel)
	defer channel.Close()

	_, errQueue := channel.QueueDeclare(
		queueEnviar,
		false,
		false,
		false,
		false,
		nil,
	)
	erroRabbit(errQueue)

	// publishing a message
	msgEnviar, errTransformar := json.Marshal(&bodyMensagem)
	erroRabbit(errTransformar)

	errPublish := channel.Publish(
		"",
		queueEnviar,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        msgEnviar,
		},
	)
	erroRabbit(errPublish)

	fmt.Println("Successfully published message em " + queueEnviar)
}

func erroRabbit(err error) {
	if err != nil {
		panic(err)
	}
}
