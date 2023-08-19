package pegando_solicitacao_pelo_id_test

import (
	"github.com/Bhimmo/golang-simple-api/adapter/repository/campo"
	"github.com/Bhimmo/golang-simple-api/adapter/repository/solicitacao"
	"github.com/Bhimmo/golang-simple-api/internal/domain/useCase/solicitacao/pegando_solicitacao_pelo_id"
	"testing"
)

func TestPegandoSolicitacaoPeloId(t *testing.T) {
	input := uint(1)
	r := solicitacao.InMemorySolicitacaoRepository{}
	rc := campo.InMemoryCampoRepository{}
	useCase := pegando_solicitacao_pelo_id.NovoPegandoSolicitacaoPeloId(&r, &rc)
	result, err := useCase.Execute(input)

	if err != nil {
		t.Errorf("Erro em pegar solicitacao")
	}
	if result.Solicitacao.PegandoId() != input {
		t.Errorf("Erro em pegar id solicitacao")
	}
	if len(result.Campos) <= 0 {
		t.Errorf("Erro em pegar campos solicitacao")
	}
}
