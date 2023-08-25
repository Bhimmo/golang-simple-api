package routes

import (
	"io"
	"net/http"
	"strings"

	"github.com/Bhimmo/golang-simple-api/adapter/controller"
)

func NovoServico(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	resp, statusCode := controller.NovoServico(body)

	w.WriteHeader(statusCode)
	w.Write(resp)
}

func PegandoServicoPeloId(w http.ResponseWriter, r *http.Request) {
	paramsReplace := strings.Replace(r.RequestURI, "/", "", -1)
	paramId := strings.Replace(paramsReplace, "servico", "", -1)

	resp, statusCode := controller.PegandoServicoPeloId(paramId)

	w.WriteHeader(statusCode)
	w.Write(resp)
}

func AdicionandoCampos(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	resp, statusCode := controller.AdicionandoCamposServico(body)

	w.WriteHeader(statusCode)
	w.Write(resp)
}
