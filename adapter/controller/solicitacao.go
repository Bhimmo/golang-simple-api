package controller

import (
	"encoding/json"
	"github.com/Bhimmo/golang-simple-api/adapter/repository/campo"
	"github.com/Bhimmo/golang-simple-api/adapter/repository/servico"
	"github.com/Bhimmo/golang-simple-api/adapter/repository/solicitacao"
	"github.com/Bhimmo/golang-simple-api/internal/domain/useCase/solicitacao/salvar_solicitacao"
	"github.com/Bhimmo/golang-simple-api/pkg/sqlite"
	"net/http"
)

func SalvarSolicitacao(body []byte) ([]byte, int) {
	var input salvar_solicitacao.SalvarSolicitacaoInput

	errBody := json.Unmarshal(body, &input)
	if errBody != nil || len(input.Campos) <= 0 {
		return []byte("Body invalido"), http.StatusBadRequest
	}

	r := solicitacao.NovoRepositorySolicitacao(sqlite.Db)
	rs := servico.NovoRepositoryServico(sqlite.Db)
	rc := campo.NovoRepositoryCampo(sqlite.Db)
	useCase := salvar_solicitacao.NovoSalvarSolicitacao(r, rs, rc)

	result, errUseCase := useCase.Execute(input)
	if errUseCase != nil {
		return []byte(errUseCase.Error()), http.StatusInternalServerError
	}

	re, _ := json.Marshal(&result)
	return re, http.StatusCreated
}
