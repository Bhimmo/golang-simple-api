package salvar_solicitacao_test

import (
	"github.com/Bhimmo/golang-simple-api/adapter/repository/servico"
	"github.com/Bhimmo/golang-simple-api/adapter/repository/solicitacao"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/campo"
	"github.com/Bhimmo/golang-simple-api/internal/domain/useCase/solicitacao/salvar_solicitacao"
	"testing"
)

func TestNovoSalvarSolicitacao(t *testing.T) {
	input := salvar_solicitacao.SalvarSolicitacaoInput{
		ServicoId:     0,
		SolicitanteId: 123,
		Campos:        []campo.Campo{{Id: uint(1), Valor: "Nao gostei do atendimento"}},
	}
	r := solicitacao.InMemorySolicitacaoRepository{}
	rs := servico.InMemoryServicoRepository{}

	useCase := salvar_solicitacao.NovoSalvarSolicitacao(&r, &rs)
	_, errSolicitacao := useCase.Execute(input)

	if errSolicitacao != nil {
		t.Errorf(error.Error(errSolicitacao))
	}
}

func TestSalvarSolcitacaoNaoPodeEstarConcluido(t *testing.T) {
	input := salvar_solicitacao.SalvarSolicitacaoInput{
		ServicoId:     0,
		SolicitanteId: 123,
		Campos:        []campo.Campo{{Id: uint(1), Valor: "Nao gostei do atendimento"}},
	}
	r := solicitacao.InMemorySolicitacaoRepository{}
	rs := servico.InMemoryServicoRepository{}

	useCase := salvar_solicitacao.NovoSalvarSolicitacao(&r, &rs)
	s, errSolicitacao := useCase.Execute(input)

	if errSolicitacao != nil {
		t.Errorf(error.Error(errSolicitacao))
	}
	if s.VerificacaoSeEstaConcluida() != false {
		t.Errorf("Solicitacao esta concluida na criacao")
	}
}
