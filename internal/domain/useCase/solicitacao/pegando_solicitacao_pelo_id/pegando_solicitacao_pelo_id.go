package pegando_solicitacao_pelo_id

import (
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/campo"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/solicitacao"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/solicitacao_campo"
)

type PegandoSolicitacaoPeloId struct {
	repositorySolicitacao      solicitacao.InterfaceSolicitacaoRepository
	repositorySolicitacaoCampo solicitacao_campo.SolicitacaoCampoInterface
	repositoryCampo            campo.InterfaceCampoRepository
}

func NovoPegandoSolicitacaoPeloId(
	solicitacaoRepository solicitacao.InterfaceSolicitacaoRepository,
	campoSolicitacaoRepository solicitacao_campo.SolicitacaoCampoInterface,
	campoRepository campo.InterfaceCampoRepository,
) *PegandoSolicitacaoPeloId {
	return &PegandoSolicitacaoPeloId{
		repositorySolicitacao:      solicitacaoRepository,
		repositorySolicitacaoCampo: campoSolicitacaoRepository,
		repositoryCampo:            campoRepository,
	}
}

func (s *PegandoSolicitacaoPeloId) Execute(id uint) (PegandoSolicitacaoPeloIdOutput, error) {
	solicitacaoBusca, errBuscaSolicitacao := s.repositorySolicitacao.BuscarPeloId(id)
	if errBuscaSolicitacao != nil {
		return PegandoSolicitacaoPeloIdOutput{}, errBuscaSolicitacao
	}
	//Campos
	campoBuscaLista, errBuscaCampo := s.repositorySolicitacaoCampo.BuscarCamposPelaSolicitacao(id)
	if errBuscaCampo != nil {
		return PegandoSolicitacaoPeloIdOutput{}, errBuscaCampo
	}

	var listaRetornoCampo []PegandoSolicitacaoPeloIdCampoOutput
	for _, itemCampo := range campoBuscaLista {
		newCampo, errCampo := s.repositoryCampo.BuscarPeloId(itemCampo.CampoId)
		if errCampo != nil {
			return PegandoSolicitacaoPeloIdOutput{}, errCampo
		}

		newItemCampo := PegandoSolicitacaoPeloIdCampoOutput{
			Id:    newCampo.Id,
			Nome:  newCampo.Nome,
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
		CreatedAt:     solicitacaoBusca.CreatedAt,
		Campos:        listaRetornoCampo,
	}, nil
}
