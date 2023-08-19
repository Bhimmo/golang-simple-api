package salvar_solicitacao

import (
	"errors"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/servico"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/solicitacao"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/status"
)

type SalvarSolicitacaoUseCase struct {
	repositorySolicitacao solicitacao.InterfaceSolicitacaoRepository
	repositoryServico     servico.InterfaceServicoRepository
}

func NovoSalvarSolicitacao(
	solicitacaoRepository solicitacao.InterfaceSolicitacaoRepository,
	servicoRepository servico.InterfaceServicoRepository,
) *SalvarSolicitacaoUseCase {
	return &SalvarSolicitacaoUseCase{
		repositorySolicitacao: solicitacaoRepository,
		repositoryServico:     servicoRepository,
	}
}

func (s *SalvarSolicitacaoUseCase) Execute(input SalvarSolicitacaoInput) (*solicitacao.Solicitacao, error) {
	servicoBusca, errServico := s.repositoryServico.PegandoPeloId(input.ServicoId)
	if errServico != nil {
		return nil, errors.New("Erro em encontrar servico")
	}
	newStatus := status.NovoStatus()

	newSolicitacao := solicitacao.NovaSolicitacao(servicoBusca, newStatus, false, input.SolicitanteId)

	errSalvarSolicitacao := s.repositorySolicitacao.Salvar(
		newSolicitacao.PegandoIdDoServicoDaSolicitacao(),
		newSolicitacao.PegandoIdDoStatusDaSolicitacao(),
	)

	if errSalvarSolicitacao != nil {
		return nil, errors.New("Erro em salvar solicitacao")
	}

	return newSolicitacao, nil
}
