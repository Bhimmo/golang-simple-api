package solicitacao

import (
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/servico"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/status"
)

type Solicitacao struct {
	id            uint
	servico       servico.Servico
	status        status.Status
	concluida     bool
	solicitanteId uint
}

func NovaSolicitacao(
	servicoInput servico.Servico,
	statusInput status.Status,
	concluida bool,
	solicitanteId uint,
) *Solicitacao {
	return &Solicitacao{
		servico:       servicoInput,
		status:        statusInput,
		concluida:     concluida,
		solicitanteId: solicitanteId,
	}
}

func (s *Solicitacao) SetandoId(id uint) {
	s.id = id
}
func (s *Solicitacao) PegandoId() uint {
	return s.id
}
func (s *Solicitacao) PegandoIdDoServicoDaSolicitacao() uint {
	return s.servico.Id
}
func (s *Solicitacao) PegandoIdDoStatusDaSolicitacao() uint {
	return s.status.Id
}
func (s *Solicitacao) VerificacaoSeEstaConcluida() bool {
	return s.concluida
}
func (s *Solicitacao) PegandoSolicitanteId() uint {
	return s.solicitanteId
}
