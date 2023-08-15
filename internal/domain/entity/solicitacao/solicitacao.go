package solicitacao

import (
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/servico"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/status"
)

type Solicitacao struct {
	id      uint
	servico servico.Servico
	status  status.Status
}

func NovaSolicitacao(servicoInput servico.Servico, statusInput status.Status) *Solicitacao {
	return &Solicitacao{
		servico: servicoInput,
		status:  statusInput,
	}
}

func (s *Solicitacao) PegandoIdDoServicoDaSolicitacao() uint {
	return s.servico.Id
}
func (s *Solicitacao) PegandoIdDoStatusDaSolicitacao() uint {
	return s.status.Id
}
