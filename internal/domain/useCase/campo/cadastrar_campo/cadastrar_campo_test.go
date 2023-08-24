package cadastrar_campo_test

import (
	"testing"

	"github.com/Bhimmo/golang-simple-api/adapter/repository/campo"
	"github.com/Bhimmo/golang-simple-api/internal/domain/useCase/campo/cadastrar_campo"
)

func TestCadastrarCampoRetorneId(t *testing.T) {
	input := cadastrar_campo.CadastrarCampoInput{
		Nome: "Campo Teste",
	}

	rc := campo.InMemoryCampoRepository{}
	useCase := cadastrar_campo.NovoCadastrarCampo(&rc)
	result, errExec := useCase.Execute(input)

	if errExec != nil {
		t.Errorf(errExec.Error())
	}
	if result.Id != 0 {
		t.Errorf("Erro no id")
	}
}
