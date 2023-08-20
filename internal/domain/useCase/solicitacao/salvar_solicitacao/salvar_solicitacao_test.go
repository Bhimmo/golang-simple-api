package salvar_solicitacao_test

import (
	campoRepo "github.com/Bhimmo/golang-simple-api/adapter/repository/campo"
	"github.com/Bhimmo/golang-simple-api/adapter/repository/servico"
	"github.com/Bhimmo/golang-simple-api/adapter/repository/solicitacao"
	"github.com/Bhimmo/golang-simple-api/internal/domain/useCase/solicitacao/salvar_solicitacao"
	"testing"
)

func TestNovoSalvarSolicitacao(t *testing.T) {
	input := salvar_solicitacao.SalvarSolicitacaoInput{
		ServicoId:     0,
		SolicitanteId: 123,
		Campos: []struct {
			Id    uint   `json:"id"`
			Valor string `json:"valor"`
		}{
			{Id: 1, Valor: "teste"},
		},
	}

	r := solicitacao.InMemorySolicitacaoRepository{}
	rs := servico.InMemoryServicoRepository{}
	rc := campoRepo.InMemoryCampoRepository{}

	useCase := salvar_solicitacao.NovoSalvarSolicitacao(&r, &rs, &rc)
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
		ServicoId:     0,
		SolicitanteId: 123,
		Campos: []struct {
			Id    uint   `json:"id"`
			Valor string `json:"valor"`
		}{
			{Id: 1, Valor: "teste"},
		},
	}
	r := solicitacao.InMemorySolicitacaoRepository{}
	rs := servico.InMemoryServicoRepository{}
	rc := campoRepo.InMemoryCampoRepository{}

	useCase := salvar_solicitacao.NovoSalvarSolicitacao(&r, &rs, &rc)
	s, errSolicitacao := useCase.Execute(input)

	if errSolicitacao != nil {
		t.Errorf(error.Error(errSolicitacao))
	}
	if s.Concluida != false {
		t.Errorf("Solicitacao esta concluida na criacao")
	}
}
