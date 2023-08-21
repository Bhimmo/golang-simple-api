package mensageria

import (
	"fmt"
	"github.com/Bhimmo/golang-simple-api/internal/infra/mensageria"
)

type RabbitMqInMemory struct{}

func NovoRabbitMqInMemory() *RabbitMqInMemory {
	return &RabbitMqInMemory{}
}

func (m *RabbitMqInMemory) EnviarEmail(
	queueEnviar string,
	bodyMensagem mensageria.MensagemEnviarRabbitmq,
) {
	fmt.Println("ENVIANDO MESNAGEM PARA: " + queueEnviar)
}
