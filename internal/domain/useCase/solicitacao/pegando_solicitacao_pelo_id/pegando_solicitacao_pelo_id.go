package pegando_solicitacao_pelo_id

import (
	"errors"

	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/solicitacao"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/solicitacao_campo"
)

type PegandoSolicitacaoPeloId struct {
	repositorySolicitacao solicitacao.InterfaceSolicitacaoRepository
	repositoryCampo       solicitacao_campo.SolicitacaoCampoInterface
}

func NovoPegandoSolicitacaoPeloId(
	solicitacaoRepository solicitacao.InterfaceSolicitacaoRepository,
	campoRepository solicitacao_campo.SolicitacaoCampoInterface,
) *PegandoSolicitacaoPeloId {
	return &PegandoSolicitacaoPeloId{
		repositorySolicitacao: solicitacaoRepository,
		repositoryCampo:       campoRepository,
	}
}

func (s *PegandoSolicitacaoPeloId) Execute(id uint) (PegandoSolicitacaoPeloIdOutput, error) {
	solicitacaoBusca, errBuscaSolicitacao := s.repositorySolicitacao.BuscarPeloId(id)
	if errBuscaSolicitacao != nil {
		return PegandoSolicitacaoPeloIdOutput{}, errors.New(errBuscaSolicitacao.Error())
	}
	//Campos
	campoBuscaLista, errBuscaCampo := s.repositoryCampo.BuscarCamposPelaSolicitacao(solicitacaoBusca.PegandoId())
	if errBuscaCampo != nil {
		return PegandoSolicitacaoPeloIdOutput{}, errors.New(errBuscaCampo.Error())
	}

	var listaRetornoCampo []PegandoSolicitacaoPeloIdCampoOutput
	for _, itemCampo := range campoBuscaLista {
		newItemCampo := PegandoSolicitacaoPeloIdCampoOutput{
			Id:    itemCampo.Id,
			Valor: itemCampo.Valor,
		}
		listaRetornoCampo = append(listaRetornoCampo, newItemCampo)
	}

	return PegandoSolicitacaoPeloIdOutput{
		Id: solicitacaoBusca.PegandoId(),
		Servico: PegandoSolicitacaoPeloIdServicoOutput{
			Id:   solicitacaoBusca.PegandoServicoSolicitacao().Id,
			Nome: solicitacaoBusca.PegandoServicoSolicitacao().Nome,
		},
		Status: PegandoSolicitacaoPeloIdStatusOutput{
			Id:   solicitacaoBusca.PegandoStatusSolicitacao().Id,
			Nome: solicitacaoBusca.PegandoStatusSolicitacao().Nome,
		},
		Concluida:     solicitacaoBusca.VerificacaoSeEstaConcluida(),
		SolicitanteId: solicitacaoBusca.PegandoSolicitanteId(),
		Campos:        listaRetornoCampo,
	}, nil
}
