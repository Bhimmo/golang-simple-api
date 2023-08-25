package salvar_solicitacao

import (
	"errors"
	"time"

	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/campo"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/servico"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/solicitacao"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/solicitacao_campo"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/status"
)

type SalvarSolicitacaoUseCase struct {
	repositorySolicitacao      solicitacao.InterfaceSolicitacaoRepository
	repositoryServico          servico.InterfaceServicoRepository
	repositoryCampo            campo.InterfaceCampoRepository
	repositorySolicitacaocampo solicitacao_campo.SolicitacaoCampoInterface
}

func NovoSalvarSolicitacao(
	solicitacaoRepository solicitacao.InterfaceSolicitacaoRepository,
	servicoRepository servico.InterfaceServicoRepository,
	campoRepository campo.InterfaceCampoRepository,
	solicitacaoCampoRepository solicitacao_campo.SolicitacaoCampoInterface,
) *SalvarSolicitacaoUseCase {
	return &SalvarSolicitacaoUseCase{
		repositorySolicitacao:      solicitacaoRepository,
		repositoryServico:          servicoRepository,
		repositoryCampo:            campoRepository,
		repositorySolicitacaocampo: solicitacaoCampoRepository,
	}
}

func (s *SalvarSolicitacaoUseCase) Execute(input SalvarSolicitacaoInput) (SalvarSolicitacaoOutput, error) {
	//Servico
	servicoBusca, errServico := s.repositoryServico.PegandoPeloId(input.ServicoId)
	if errServico != nil {
		return SalvarSolicitacaoOutput{}, errors.New("erro em encontrar servico")
	}
	//Status
	newStatus := status.NovoStatus()
	newStatus.TendoStatusInicial()
	//Solicitacao
	newSolicitacao := solicitacao.NovaSolicitacao(
		servicoBusca,
		newStatus,
		false,
		input.SolicitanteId,
		time.Now(),
	)

	idSolicitacao, errSalvarSolicitacao := s.repositorySolicitacao.Salvar(
		newSolicitacao.PegandoServicoSolicitacao().Id,
		newSolicitacao.PegandoStatusSolicitacao().Id,
		newSolicitacao.VerificacaoSeEstaConcluida(),
		newSolicitacao.PegandoSolicitanteId(),
	)
	if errSalvarSolicitacao != nil {
		return SalvarSolicitacaoOutput{}, errors.New("erro em salvar solicitacao")
	}
	newSolicitacao.SetandoId(idSolicitacao)

	//Campos
	for _, itemCampo := range input.Campos {
		newCampo := campo.NovoCampo()
		newCampo.Id = itemCampo.Id

		campoBusca, errCampo := s.repositoryCampo.BuscarPeloId(newCampo.Id)
		if errCampo != nil {
			return SalvarSolicitacaoOutput{}, errCampo
		}

		errCampoSolicitacao := s.repositorySolicitacaocampo.SalvarCamposDaSolicitacao(
			campoBusca.Id,
			newSolicitacao.PegandoId(),
			itemCampo.Valor,
		)
		if errCampoSolicitacao != nil {
			return SalvarSolicitacaoOutput{}, errCampoSolicitacao
		}
	}

	return SalvarSolicitacaoOutput{
		Id:            newSolicitacao.PegandoId(),
		Concluida:     newSolicitacao.VerificacaoSeEstaConcluida(),
		SolicitanteId: newSolicitacao.PegandoSolicitanteId(),
		ServicoId:     newSolicitacao.PegandoServicoSolicitacao().Id,
		Campos:        input.Campos,
	}, nil
}
