package routes

import (
	"github.com/Bhimmo/golang-simple-api/adapter/controller"
	"io"
	"net/http"
	"strings"
)

func SalvarSolicitacao(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	resp, statusCode := controller.SalvarSolicitacao(body)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(resp)
}

func PegandoSolicitacaoPeloId(w http.ResponseWriter, r *http.Request) {
	paramsReplace := strings.Replace(r.RequestURI, "/", "", -1)
	paramId := strings.Replace(paramsReplace, "solicitacao", "", -1)

	resp, statusCode := controller.PegandoSolicitacaoPeloId(paramId)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(resp)
}

func AtualizarStatusSolicitacao(w http.ResponseWriter, r *http.Request) {
	paramsReplace := strings.Replace(r.RequestURI, "/", "", -1)
	paramReplaceRota := strings.Replace(paramsReplace, "solicitacao", "", -1)
	paramId := strings.Split(paramReplaceRota, "atualizar-status")[0]

	resp, statusCode := controller.AtualizandoStatusSolicitacao(paramId)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(resp)
}
