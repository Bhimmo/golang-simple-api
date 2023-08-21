package solicitacao

import (
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/servico"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/solicitacao"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/status"
)

type InMemorySolicitacaoRepositoryInput struct {
	Id            uint
	ServicoId     uint
	StatusId      uint
	Concluida     bool
	SolicitanteId uint
}
type InMemorySolicitacaoRepository struct {
	Solicitacao []InMemorySolicitacaoRepositoryInput
}

func (r *InMemorySolicitacaoRepository) Salvar(
	servicoId uint,
	statusId uint,
	concluida bool,
	solicitanteId uint,
) (uint, error) {
	idItem := uint(1)
	itemSalvar := InMemorySolicitacaoRepositoryInput{
		Id:            idItem,
		ServicoId:     servicoId,
		StatusId:      statusId,
		Concluida:     concluida,
		SolicitanteId: solicitanteId,
	}
	r.Solicitacao = append(r.Solicitacao, itemSalvar)
	return idItem, nil
}

func (r *InMemorySolicitacaoRepository) BuscarPeloId(id uint) (solicitacao.Solicitacao, error) {
	item := r.Solicitacao[id-1]
	Entityservico := servico.NovoServico()
	Entityservico.Id = item.ServicoId

	EntityStatus := status.NovoStatus()
	EntityStatus.TendoStatusDesejado(item.StatusId)

	s := solicitacao.NovaSolicitacao(
		*Entityservico,
		EntityStatus,
		item.Concluida,
		item.SolicitanteId,
	)
	s.SetandoId(item.Id)
	return *s, nil
}

func (r *InMemorySolicitacaoRepository) AtualizarSolicitacao(solicitacao solicitacao.Solicitacao) error {
	updateItem := InMemorySolicitacaoRepositoryInput{
		Id:            solicitacao.PegandoId(),
		StatusId:      solicitacao.PegandoStatusSolicitacao().Id,
		ServicoId:     solicitacao.PegandoServicoSolicitacao().Id,
		Concluida:     solicitacao.VerificacaoSeEstaConcluida(),
		SolicitanteId: solicitacao.PegandoSolicitanteId(),
	}
	r.Solicitacao[solicitacao.PegandoId()-1] = updateItem
	return nil
}
