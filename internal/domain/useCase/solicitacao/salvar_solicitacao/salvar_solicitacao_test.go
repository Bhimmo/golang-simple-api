package salvar_solicitacao_test

import (
	"testing"

	campoRepo "github.com/Bhimmo/golang-simple-api/adapter/repository/campo"
	"github.com/Bhimmo/golang-simple-api/adapter/repository/servico"
	"github.com/Bhimmo/golang-simple-api/adapter/repository/solicitacao"
	"github.com/Bhimmo/golang-simple-api/adapter/repository/solicitacao_campo"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/campo"
	"github.com/Bhimmo/golang-simple-api/internal/domain/useCase/solicitacao/salvar_solicitacao"
)

var rc campoRepo.InMemoryCampoRepository

func Setup() {
	//Salvar um campo
	_, _ = rc.Salvar(campo.Campo{
		Id: 1, Nome: "Teste daniel",
	})
}

func TestNovoSalvarSolicitacao(t *testing.T) {
	Setup()

	input := salvar_solicitacao.SalvarSolicitacaoInput{
		ServicoId:     1,
		SolicitanteId: 123,
		Campos: []salvar_solicitacao.SalvarSolicitacaoCampoOutput{
			{Id: 1, Valor: "teste"},
		},
	}

	r := solicitacao.InMemorySolicitacaoRepository{}
	rs := servico.InMemoryServicoRepository{}
	rsc := solicitacao_campo.RepositorySolicitacaoCampoInMemory{}

	useCase := salvar_solicitacao.NovoSalvarSolicitacao(&r, &rs, &rc, &rsc)
	s, errSolicitacao := useCase.Execute(input)

	if errSolicitacao != nil {
		t.Errorf(error.Error(errSolicitacao))
	}
	if s.Id != 1 {
		t.Errorf("Id nao encluso")
	}
}

func TestSalvarSolcitacaoNaoPodeEstarConcluido(t *testing.T) {
	input := salvar_solicitacao.SalvarSolicitacaoInput{
		ServicoId:     1,
		SolicitanteId: 123,
		Campos: []salvar_solicitacao.SalvarSolicitacaoCampoOutput{
			{Id: 1, Valor: "teste"},
		},
	}
	r := solicitacao.InMemorySolicitacaoRepository{}
	rs := servico.InMemoryServicoRepository{}
	rsc := solicitacao_campo.RepositorySolicitacaoCampoInMemory{}

	useCase := salvar_solicitacao.NovoSalvarSolicitacao(&r, &rs, &rc, &rsc)
	s, errSolicitacao := useCase.Execute(input)

	if errSolicitacao != nil {
		t.Errorf(error.Error(errSolicitacao))
	}
	if s.Concluida != false {
		t.Errorf("Solicitacao esta concluida na criacao")
	}
}
