package atualizar_status_solicitacao

import (
	"errors"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/solicitacao"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/status"
)

type AtualizarStatusSolicitacao struct {
	repositorySolicitacao solicitacao.InterfaceSolicitacaoRepository
}

func NovoAtualizarStatusSolicitacao(
	solicitacaoRepository solicitacao.InterfaceSolicitacaoRepository,
) *AtualizarStatusSolicitacao {
	return &AtualizarStatusSolicitacao{
		repositorySolicitacao: solicitacaoRepository,
	}
}

func (s *AtualizarStatusSolicitacao) Execute(id uint) (AtualizarStatusSolicitacaoOutput, error) {
	//Pegar solicitacao
	solicitacaoBusca, errBuscaSolicitacao := s.repositorySolicitacao.BuscarPeloId(id)
	statusSolicitacaoBusca := solicitacaoBusca.PegandoStatusSolicitacao()
	if errBuscaSolicitacao != nil {
		return AtualizarStatusSolicitacaoOutput{}, errors.New(errBuscaSolicitacao.Error())
	}

	//Verificar o status
	EntityStatus := status.NovoStatus()
	EntityStatus.TendoStatusDesejado(solicitacaoBusca.PegandoStatusSolicitacao().Id)

	//Atualizar o status
	EntityStatus.ProximoStatus()
	EntitySolicitacao := solicitacao.NovaSolicitacao(
		solicitacaoBusca.PegandoServicoSolicitacao(),
		EntityStatus,
		solicitacaoBusca.VerificacaoSeEstaConcluida(),
		solicitacaoBusca.PegandoSolicitanteId(),
	)
	EntitySolicitacao.SetandoId(solicitacaoBusca.PegandoId())

	//Solicitacao concluida ?
	if EntityStatus.VerificaUltimoStatus() {
		EntitySolicitacao.EstaConcluida()
	}

	//Salvar no banco "update" caso ainda nao finalizado
	if statusSolicitacaoBusca.VerificaUltimoStatus() == false {
		errUpdate := s.repositorySolicitacao.AtualizarSolicitacao(*EntitySolicitacao)
		if errUpdate != nil {
			return AtualizarStatusSolicitacaoOutput{}, errors.New("Erro em atualizar solicitacao")
		}
	}

	return AtualizarStatusSolicitacaoOutput{
		Id:        EntitySolicitacao.PegandoId(),
		Concluida: EntitySolicitacao.VerificacaoSeEstaConcluida(),
		Status: AtualizarStatusSolicitacaoStatusOutput{
			Id:   EntitySolicitacao.PegandoStatusSolicitacao().Id,
			Nome: EntitySolicitacao.PegandoStatusSolicitacao().Nome,
		},
	}, nil
}
