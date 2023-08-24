package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Bhimmo/golang-simple-api/adapter/repository/campo"
	"github.com/Bhimmo/golang-simple-api/internal/domain/useCase/campo/cadastrar_campo"
	"github.com/Bhimmo/golang-simple-api/internal/domain/useCase/campo/pegando_campo_pelo_id"
	"github.com/Bhimmo/golang-simple-api/pkg/sqlite"
)

func NovoCampo(body []byte) (cadastrar_campo.CadastrarCampoOutput, int) {
	var input cadastrar_campo.CadastrarCampoInput

	errUm := json.Unmarshal(body, &input)
	if errUm != nil {
		return cadastrar_campo.CadastrarCampoOutput{}, http.StatusBadRequest
	}

	rc := campo.NovoRepositoryCampo(sqlite.Db)
	useCase := cadastrar_campo.NovoCadastrarCampo(rc)
	result, _ := useCase.Execute(input)

	return result, http.StatusCreated
}

func PegandoCampoById(id string) ([]byte, int) {
	idInt, errInt := strconv.Atoi(id)
	if errInt != nil {
		return []byte(errInt.Error()), http.StatusBadRequest
	}

	input := pegando_campo_pelo_id.PegandoCampoPeloIdInput{
		Id: uint(idInt),
	}
	rc := campo.NovoRepositoryCampo(sqlite.Db)
	useCase := pegando_campo_pelo_id.NewPegandoCampoPeloId(rc)
	result, errExec := useCase.Execute(input)

	if errExec != nil {
		return []byte(errExec.Error()), http.StatusInternalServerError
	}

	reBy, _ := json.Marshal(&result)
	return reBy, http.StatusOK
}
