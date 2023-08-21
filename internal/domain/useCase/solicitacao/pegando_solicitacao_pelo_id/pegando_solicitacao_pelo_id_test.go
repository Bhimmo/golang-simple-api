package pegando_solicitacao_pelo_id_test

import (
	"github.com/Bhimmo/golang-simple-api/adapter/repository/campo"
	"github.com/Bhimmo/golang-simple-api/adapter/repository/servico"
	"github.com/Bhimmo/golang-simple-api/adapter/repository/solicitacao"
	"github.com/Bhimmo/golang-simple-api/internal/domain/useCase/solicitacao/pegando_solicitacao_pelo_id"
	"github.com/Bhimmo/golang-simple-api/internal/domain/useCase/solicitacao/salvar_solicitacao"
	"testing"
)

var r solicitacao.InMemorySolicitacaoRepository
var rc campo.InMemoryCampoRepository

func Setup() {
	r = solicitacao.InMemorySolicitacaoRepository{}
	rc = campo.InMemoryCampoRepository{}
	rs := servico.InMemoryServicoRepository{}

	input := salvar_solicitacao.SalvarSolicitacaoInput{
		ServicoId:     1,
		SolicitanteId: 123,
		Campos: []salvar_solicitacao.SalvarSolicitacaoCampoOutput{
			{Id: 1, Valor: "teste"},
		},
	}
	useCase := salvar_solicitacao.NovoSalvarSolicitacao(&r, &rs, &rc)
	_, _ = useCase.Execute(input)
}

func TestPegandoSolicitacaoPeloId(t *testing.T) {
	Setup()

	input := uint(1)
	useCase := pegando_solicitacao_pelo_id.NovoPegandoSolicitacaoPeloId(&r, &rc)
	result, err := useCase.Execute(input)

	if err != nil {
		t.Errorf("Erro em pegar solicitacao")
	}
	if result.Id != input {
		t.Errorf("Erro em pegar id solicitacao")
	}
	if len(result.Campos) <= 0 {
		t.Errorf("Erro em pegar campos solicitacao")
	}
}
