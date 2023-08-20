package routes

import (
	"github.com/Bhimmo/golang-simple-api/adapter/controller"
	"io"
	"net/http"
)

func SalvarSolicitacao(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	resp, statusCode := controller.SalvarSolicitacao(body)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(resp)
}
