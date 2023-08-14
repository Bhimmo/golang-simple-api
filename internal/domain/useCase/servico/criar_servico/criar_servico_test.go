package criar_servico_test

import (
	"github.com/Bhimmo/golang-simple-api/adapter/repository/servico"
	"github.com/Bhimmo/golang-simple-api/internal/domain/useCase/servico/criar_servico"
	"testing"
)

func TestCriarNovoServico(t *testing.T) {
	input := criar_servico.CriarServicoInput{
		Nome: "Servico test",
	}
	r := servico.InMemoryServicoRepository{}
	criarServicoUseCase := criar_servico.NovoCriarServicoUseCase(&r)
	s, err := criarServicoUseCase.Execute(input)

	if err != nil {
		t.Errorf("Error na criar do servico")
	}
	if s.Nome != "Servico test" {
		t.Errorf("Error na criar do servico, nome diferente")
	}
}
