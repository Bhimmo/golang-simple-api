package atualizar_status_solicitacao_test

import (
	"github.com/Bhimmo/golang-simple-api/adapter/repository/campo"
	"github.com/Bhimmo/golang-simple-api/adapter/repository/servico"
	"github.com/Bhimmo/golang-simple-api/adapter/repository/solicitacao"
	"github.com/Bhimmo/golang-simple-api/internal/domain/useCase/solicitacao/atualizar_status_solicitacao"
	"github.com/Bhimmo/golang-simple-api/internal/domain/useCase/solicitacao/salvar_solicitacao"
	"testing"
)

var r solicitacao.InMemorySolicitacaoRepository

func Setup() {
	r = solicitacao.InMemorySolicitacaoRepository{}
	input := salvar_solicitacao.SalvarSolicitacaoInput{
		ServicoId:     1,
		SolicitanteId: 123,
		Campos: []salvar_solicitacao.SalvarSolicitacaoCampoOutput{
			{Id: 1, Valor: "teste"},
		},
	}
	rs := servico.InMemoryServicoRepository{}
	rc := campo.InMemoryCampoRepository{}
	useCase := salvar_solicitacao.NovoSalvarSolicitacao(&r, &rs, &rc)
	_, _ = useCase.Execute(input)
}

func TestPrimeiroStatusAtualizar(t *testing.T) {
	Setup()

	useCase := atualizar_status_solicitacao.NovoAtualizarStatusSolicitacao(&r)

	result, errExec := useCase.Execute(1)

	if errExec != nil {
		t.Errorf("Caso de uso falhou")
	}
	if result.Concluida != false {
		t.Errorf("Caso de uso falhou na atualizacao de status")
	}
}

func TestAtualizarParaUltimoStatusMostrarConcluidoComTrue(t *testing.T) {
	useCase := atualizar_status_solicitacao.NovoAtualizarStatusSolicitacao(&r)

	result, errExec := useCase.Execute(1)

	if errExec != nil {
		t.Errorf("Caso de uso falhou")
	}
	if result.Concluida != true {
		t.Errorf("Caso de uso falhou na atualizacao de status")
	}
}

func TestForUltimoStatusNaoAtualizar(t *testing.T) {
	useCase := atualizar_status_solicitacao.NovoAtualizarStatusSolicitacao(&r)

	result, errExec := useCase.Execute(1)

	if errExec != nil {
		t.Errorf("Caso de uso falhou")
	}
	if result.Concluida != true {
		t.Errorf("Caso de uso falhou na atualizacao de status")
	}
	if result.Status.Id != 3 {
		t.Errorf("Caso de uso falhou status nao permitido nessa etapa")
	}
}
