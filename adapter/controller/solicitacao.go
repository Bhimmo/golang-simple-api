package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Bhimmo/golang-simple-api/adapter/repository/campo"
	"github.com/Bhimmo/golang-simple-api/adapter/repository/mensageria"
	"github.com/Bhimmo/golang-simple-api/adapter/repository/servico"
	"github.com/Bhimmo/golang-simple-api/adapter/repository/solicitacao"
	"github.com/Bhimmo/golang-simple-api/adapter/repository/solicitacao_campo"
	"github.com/Bhimmo/golang-simple-api/internal/domain/useCase/solicitacao/atualizar_status_solicitacao"
	"github.com/Bhimmo/golang-simple-api/internal/domain/useCase/solicitacao/pegando_solicitacao_pelo_id"
	"github.com/Bhimmo/golang-simple-api/internal/domain/useCase/solicitacao/salvar_solicitacao"
	rabbitmq2 "github.com/Bhimmo/golang-simple-api/pkg/rabbitmq"
	"github.com/Bhimmo/golang-simple-api/pkg/sqlite"
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
	rsc := solicitacao_campo.NewRepositorySolicitacaoCampo(sqlite.Db)
	useCase := salvar_solicitacao.NovoSalvarSolicitacao(r, rs, rc, rsc)

	result, errUseCase := useCase.Execute(input)
	if errUseCase != nil {
		return []byte(errUseCase.Error()), http.StatusInternalServerError
	}

	re, _ := json.Marshal(&result)
	return re, http.StatusCreated
}

func PegandoSolicitacaoPeloId(id string) ([]byte, int) {
	idInt, errConv := strconv.Atoi(id)
	if errConv != nil {
		return []byte("params invalidos"), http.StatusBadRequest
	}

	r := solicitacao.NovoRepositorySolicitacao(sqlite.Db)
	rc := campo.NovoRepositoryCampo(sqlite.Db)
	rsc := solicitacao_campo.NewRepositorySolicitacaoCampo(sqlite.Db)
	useCase := pegando_solicitacao_pelo_id.NovoPegandoSolicitacaoPeloId(r, rsc, rc)

	result, errUseCase := useCase.Execute(uint(idInt))
	if errUseCase != nil {
		return []byte(errUseCase.Error()), http.StatusInternalServerError
	}

	re, _ := json.Marshal(&result)
	return re, http.StatusOK
}

func AtualizandoStatusSolicitacao(id string) ([]byte, int) {
	idInt, errConv := strconv.Atoi(id)
	if errConv != nil {
		return []byte("params invalidos"), http.StatusBadRequest
	}

	r := solicitacao.NovoRepositorySolicitacao(sqlite.Db)
	rMensageria := mensageria.NovoRabbitMq(rabbitmq2.Rabbitmq)
	useCase := atualizar_status_solicitacao.NovoAtualizarStatusSolicitacao(r, rMensageria)
	result, errUseCase := useCase.Execute(uint(idInt))
	if errUseCase != nil {
		return []byte(errUseCase.Error()), http.StatusInternalServerError
	}

	re, _ := json.Marshal(&result)
	return re, http.StatusOK
}
