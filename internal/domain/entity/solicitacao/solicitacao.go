package solicitacao

import (
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/servico"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/status"
)

type solicitacao struct {
	id      uint
	servico servico.Servico
	status  status.Status
}

func NovaSolicitacao(servicoInput servico.Servico, statusInput status.Status) *solicitacao {
	return &solicitacao{
		servico: servicoInput,
		status:  statusInput,
	}
}
