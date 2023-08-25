package solicitacao_campo

import (
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/solicitacao_campo"
)

type RepositorySolicitacaoCampoInMemory struct {
	SolicitacaoCampo []solicitacao_campo.SolicitacaoCampo
}

func (r *RepositorySolicitacaoCampoInMemory) SalvarCamposDaSolicitacao(
	campoId uint,
	solicitacaoId uint,
	valor string,
) error {
	newSolicitacaoCampo := solicitacao_campo.SolicitacaoCampo{
		Id:            uint(len(r.SolicitacaoCampo) - 1),
		CampoId:       campoId,
		SolicitacaoId: solicitacaoId,
		Valor:         valor,
	}
	r.SolicitacaoCampo = append(r.SolicitacaoCampo, newSolicitacaoCampo)

	return nil
}

func (r *RepositorySolicitacaoCampoInMemory) BuscarCamposPelaSolicitacao(
	solicitacaoId uint,
) ([]solicitacao_campo.SolicitacaoCampo, error) {
	var listaSolicitacao []solicitacao_campo.SolicitacaoCampo
	for _, itemSolicitacao := range r.SolicitacaoCampo {
		if itemSolicitacao.SolicitacaoId == solicitacaoId {
			listaSolicitacao = append(listaSolicitacao, itemSolicitacao)
		}
	}

	return listaSolicitacao, nil
}
