package routes

import (
	"github.com/Bhimmo/golang-simple-api/adapter/controller"
	"io"
	"net/http"
	"strings"
)

func NovoServico(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	resp, statusCode := controller.NovoServico(body)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(resp)
}

func PegandoServicoPeloId(w http.ResponseWriter, r *http.Request) {
	paramsReplace := strings.Replace(r.RequestURI, "/", "", -1)
	paramId := strings.Replace(paramsReplace, "servico", "", -1)

	resp, statusCode := controller.PegandoServicoPeloId(paramId)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(resp)
}
