package atualizar_status_solicitacao_test

import (
	"testing"

	"github.com/Bhimmo/golang-simple-api/adapter/repository/mensageria"
	"github.com/Bhimmo/golang-simple-api/adapter/repository/solicitacao"
	"github.com/Bhimmo/golang-simple-api/internal/domain/useCase/solicitacao/atualizar_status_solicitacao"
)

var r solicitacao.InMemorySolicitacaoRepository
var rm mensageria.RabbitMqInMemory

func Setup() {
	r = solicitacao.InMemorySolicitacaoRepository{}
	rm = mensageria.RabbitMqInMemory{}

	r.Salvar(1, 1, false, 123)
}

func TestPrimeiroStatusAtualizar(t *testing.T) {
	Setup()

	useCase := atualizar_status_solicitacao.NovoAtualizarStatusSolicitacao(&r, &rm)

	result, errExec := useCase.Execute(1)

	if errExec != nil {
		t.Errorf("Caso de uso falhou")
	}
	if result.Concluida != false {
		t.Errorf("Caso de uso falhou na atualizacao de status")
	}
}

func TestAtualizarParaUltimoStatusMostrarConcluidoComTrue(t *testing.T) {
	useCase := atualizar_status_solicitacao.NovoAtualizarStatusSolicitacao(&r, &rm)

	result, errExec := useCase.Execute(1)

	if errExec != nil {
		t.Errorf("Caso de uso falhou")
	}
	if result.Concluida != true {
		t.Errorf("Caso de uso falhou na atualizacao de status")
	}
}

func TestForUltimoStatusNaoAtualizar(t *testing.T) {
	useCase := atualizar_status_solicitacao.NovoAtualizarStatusSolicitacao(&r, &rm)

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
