package atualizar_status_solicitacao

import (
	"errors"

	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/solicitacao"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/status"
	"github.com/Bhimmo/golang-simple-api/internal/infra/mensageria"
)

type AtualizarStatusSolicitacao struct {
	repositorySolicitacao solicitacao.InterfaceSolicitacaoRepository
	repositoryMensageria  mensageria.InterfaceMensageria
}

func NovoAtualizarStatusSolicitacao(
	solicitacaoRepository solicitacao.InterfaceSolicitacaoRepository,
	mensageriaRepository mensageria.InterfaceMensageria,
) *AtualizarStatusSolicitacao {
	return &AtualizarStatusSolicitacao{
		repositorySolicitacao: solicitacaoRepository,
		repositoryMensageria:  mensageriaRepository,
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
		solicitacaoBusca.CreatedAt,
	)
	EntitySolicitacao.SetandoId(solicitacaoBusca.PegandoId())

	//Solicitacao concluida e a solicitacao do banco nao pode estar como concluida
	if EntityStatus.VerificaUltimoStatus() && !statusSolicitacaoBusca.VerificaUltimoStatus() {
		EntitySolicitacao.EstaConcluida()
		s.repositoryMensageria.EnviarEmail(
			"EnviarEmail",
			mensageria.MensagemEnviarRabbitmq{
				SolicitanteId: EntitySolicitacao.PegandoSolicitanteId(),
				Conteudo:      "Solicitacao finalizada",
			},
		)
	}

	//Salvar no banco "update" caso ainda nao finalizado
	if !statusSolicitacaoBusca.VerificaUltimoStatus() {
		errUpdate := s.repositorySolicitacao.AtualizarSolicitacao(*EntitySolicitacao)
		if errUpdate != nil {
			return AtualizarStatusSolicitacaoOutput{}, errors.New("erro em atualizar solicitacao")
		}
	}

	return AtualizarStatusSolicitacaoOutput{
		Id:        EntitySolicitacao.PegandoId(),
		Concluida: EntitySolicitacao.VerificacaoSeEstaConcluida(),
		Status:    EntitySolicitacao.PegandoStatusSolicitacao().Nome,
	}, nil
}
