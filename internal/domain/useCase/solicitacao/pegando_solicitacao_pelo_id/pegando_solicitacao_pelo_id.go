package pegando_solicitacao_pelo_id

import (
	"errors"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/campo"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/solicitacao"
)

type PegandoSolicitacaoPeloId struct {
	repositorySolicitacao solicitacao.InterfaceSolicitacaoRepository
	repositoryCampo       campo.InterfaceCampoRepository
}

func NovoPegandoSolicitacaoPeloId(
	solicitacaoRepository solicitacao.InterfaceSolicitacaoRepository,
	campoRepository campo.InterfaceCampoRepository,
) *PegandoSolicitacaoPeloId {
	return &PegandoSolicitacaoPeloId{
		repositorySolicitacao: solicitacaoRepository,
		repositoryCampo:       campoRepository,
	}
}

func (s *PegandoSolicitacaoPeloId) Execute(id uint) (PegandoSolicitacaoPeloIdOutput, error) {
	solicitacaoBusca, errBuscaSolicitacao := s.repositorySolicitacao.BuscarPeloId(id)
	if errBuscaSolicitacao != nil {
		return PegandoSolicitacaoPeloIdOutput{}, errors.New("Erro em busca solicitacao")
	}
	//Campos
	campoBuscaLista, errBuscaCampo := s.repositoryCampo.BuscarCampoPeloSolicitanteId(solicitacaoBusca.PegandoSolicitanteId())
	if errBuscaCampo != nil {
		return PegandoSolicitacaoPeloIdOutput{}, errors.New("Erro em busca campos")
	}

	return PegandoSolicitacaoPeloIdOutput{
		Solicitacao: solicitacaoBusca,
		Campos:      campoBuscaLista,
	}, nil
}
