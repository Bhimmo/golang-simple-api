package solicitacao

import (
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/servico"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/solicitacao"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/status"
)

type InMemorySolicitacaoRepository struct {
	Solicitacao []any
}

func (r *InMemorySolicitacaoRepository) Salvar(
	servicoId uint,
	statusId uint,
	concluida bool,
	solicitanteId uint,
) (uint, error) {
	itemSalvar := map[uint][]any{
		1: {servicoId, statusId, concluida, solicitanteId},
	}
	r.Solicitacao = append(r.Solicitacao, itemSalvar)
	return 1, nil
}

func (r *InMemorySolicitacaoRepository) BuscarPeloId(id uint) (solicitacao.Solicitacao, error) {
	solicitacaoReturn := solicitacao.NovaSolicitacao(
		servico.Servico{Id: 1, Nome: "Test"}, status.Status{Id: 1}, false, 4,
	)
	solicitacaoReturn.SetandoId(id)
	return *solicitacaoReturn, nil
}
