package salvar_solicitacao

import (
	"errors"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/campo"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/servico"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/solicitacao"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/status"
)

type SalvarSolicitacaoUseCase struct {
	repositorySolicitacao solicitacao.InterfaceSolicitacaoRepository
	repositoryServico     servico.InterfaceServicoRepository
	repositoryCampo       campo.InterfaceCampoRepository
}

func NovoSalvarSolicitacao(
	solicitacaoRepository solicitacao.InterfaceSolicitacaoRepository,
	servicoRepository servico.InterfaceServicoRepository,
	campoRepository campo.InterfaceCampoRepository,
) *SalvarSolicitacaoUseCase {
	return &SalvarSolicitacaoUseCase{
		repositorySolicitacao: solicitacaoRepository,
		repositoryServico:     servicoRepository,
		repositoryCampo:       campoRepository,
	}
}

func (s *SalvarSolicitacaoUseCase) Execute(input SalvarSolicitacaoInput) (*solicitacao.Solicitacao, error) {
	//Servico
	servicoBusca, errServico := s.repositoryServico.PegandoPeloId(input.ServicoId)
	if errServico != nil {
		return nil, errors.New("Erro em encontrar servico")
	}
	//Status
	newStatus := status.NovoStatus()
	//Solicitacao
	newSolicitacao := solicitacao.NovaSolicitacao(
		servicoBusca,
		newStatus,
		false,
		input.SolicitanteId,
	)

	idSolicitacao, errSalvarSolicitacao := s.repositorySolicitacao.Salvar(
		newSolicitacao.PegandoIdDoServicoDaSolicitacao(),
		newSolicitacao.PegandoIdDoStatusDaSolicitacao(),
		newSolicitacao.VerificacaoSeEstaConcluida(),
		newSolicitacao.PegandoSolicitanteId(),
	)
	if errSalvarSolicitacao != nil {
		return nil, errors.New("Erro em salvar solicitacao")
	}

	newSolicitacao.SetandoId(idSolicitacao)
	//Campos
	for _, itemCampo := range input.Campos {
		newCampo := campo.NovoCampo(itemCampo.Id, itemCampo.Valor, newSolicitacao.PegandoId())
		errCampo := s.repositoryCampo.Salvar(*newCampo)
		if errCampo != nil {
			return nil, errors.New("Erro em salvar campos")
		}
	}

	return newSolicitacao, nil
}
