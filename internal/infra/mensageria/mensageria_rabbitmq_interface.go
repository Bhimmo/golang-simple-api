package mensageria

type MensagemEnviarRabbitmq struct {
	SolicitanteId uint   `json:"solicitante_id"`
	Conteudo      string `json:"conteudo"`
}
type InterfaceMensageria interface {
	EnviarEmail(
		queueEnviar string,
		bodyMensagem MensagemEnviarRabbitmq,
	)
}
