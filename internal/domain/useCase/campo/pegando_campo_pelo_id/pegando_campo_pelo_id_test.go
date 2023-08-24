package pegando_campo_pelo_id_test

import (
	"testing"

	"github.com/Bhimmo/golang-simple-api/adapter/repository/campo"
	"github.com/Bhimmo/golang-simple-api/internal/domain/useCase/campo/cadastrar_campo"
	"github.com/Bhimmo/golang-simple-api/internal/domain/useCase/campo/pegando_campo_pelo_id"
)

var rc campo.InMemoryCampoRepository

func Setup() {
	//Create campo for validate
	input := cadastrar_campo.CadastrarCampoInput{
		Nome: "Teste daniel",
	}

	useCase := cadastrar_campo.NovoCadastrarCampo(&rc)
	useCase.Execute(input)
}

func TestPegandoCampoPeloId(t *testing.T) {
	Setup()

	input := pegando_campo_pelo_id.PegandoCampoPeloIdInput{
		Id: 1,
	}

	useCase := pegando_campo_pelo_id.NewPegandoCampoPeloId(&rc)
	result, errExec := useCase.Execute(input)

	if errExec != nil {
		t.Errorf(errExec.Error())
	}
	if result.Nome == "" {
		t.Errorf("Nome vazio test campo pelo id")
	}
}
