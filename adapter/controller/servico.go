package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Bhimmo/golang-simple-api/adapter/repository/campo"
	"github.com/Bhimmo/golang-simple-api/adapter/repository/servico"
	"github.com/Bhimmo/golang-simple-api/adapter/repository/servico_campo"
	cadastrando_campo_servico "github.com/Bhimmo/golang-simple-api/internal/domain/useCase/servico/cadastrando_campo"
	"github.com/Bhimmo/golang-simple-api/internal/domain/useCase/servico/criar_servico"
	"github.com/Bhimmo/golang-simple-api/internal/domain/useCase/servico/pegando_pelo_id"
	"github.com/Bhimmo/golang-simple-api/pkg/sqlite"
)

func NovoServico(body []byte) ([]byte, int) {
	var input criar_servico.CriarServicoInput
	var inputVazio criar_servico.CriarServicoInput

	err := json.Unmarshal(body, &input)
	if err != nil || input == inputVazio {
		return []byte("Erro input data"), http.StatusBadRequest
	}

	rs := servico.NovoRepositoryServico(sqlite.Db)
	useCase := criar_servico.NovoCriarServicoUseCase(rs)
	result, errUseCase := useCase.Execute(input)

	if errUseCase != nil {
		return []byte(errUseCase.Error()), http.StatusInternalServerError
	}

	re, _ := json.Marshal(result)
	return re, http.StatusCreated
}

func PegandoServicoPeloId(id string) ([]byte, int) {
	var input pegando_pelo_id.PegandoPeloIdInput
	idInt, errInt := strconv.Atoi(id)
	if errInt != nil {
		return []byte("Parametros invalidos"), http.StatusBadRequest
	}
	input.Id = uint(idInt)

	rs := servico.NovoRepositoryServico(sqlite.Db)
	useCase := pegando_pelo_id.NovoPegandoPeloId(rs)
	result, errUseCase := useCase.Execute(input)
	if errUseCase != nil {
		return []byte(errUseCase.Error()), http.StatusInternalServerError
	}

	re, _ := json.Marshal(result)
	return re, http.StatusOK
}

func AdicionandoCamposServico(body []byte) ([]byte, int) {
	var input cadastrando_campo_servico.CadastrandoCampoServicoInput

	errUm := json.Unmarshal(body, &input)
	if errUm != nil {
		return []byte(errUm.Error()), http.StatusBadRequest
	}

	rsc := servico_campo.NewRepositoryServicoCampo(sqlite.Db)
	rs := servico.NovoRepositoryServico(sqlite.Db)
	rc := campo.NovoRepositoryCampo(sqlite.Db)
	useCase := cadastrando_campo_servico.NewCadastrandoCampoServico(rsc, rs, rc)

	errorUseCase := useCase.Execute(input)
	if errorUseCase != nil {
		return []byte(errorUseCase.Error()), http.StatusInternalServerError
	}

	return []byte(""), http.StatusNoContent
}
