package pegando_pelo_id_test

import (
	"github.com/Bhimmo/golang-simple-api/adapter/repository/servico"
	servico2 "github.com/Bhimmo/golang-simple-api/internal/domain/entity/servico"
	"github.com/Bhimmo/golang-simple-api/internal/domain/useCase/servico/pegando_pelo_id"
	"testing"
)

func TestPegandoUmServicoPeloId(t *testing.T) {
	input := pegando_pelo_id.PegandoPeloIdInput{
		Id: 0,
	}
	r := servico.InMemoryServicoRepository{}
	pegandoPeloIdUseCase := pegando_pelo_id.NovoPegandoPeloId(&r)
	s, err := pegandoPeloIdUseCase.Execute(input)

	if err != nil {
		t.Errorf("Servico nao encontrado")
	}
	if &s == servico2.NovoServico() {
		t.Errorf("Retornando uma instancia de servico vazia")
	}
}
