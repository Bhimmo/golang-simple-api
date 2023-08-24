package routes

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/Bhimmo/golang-simple-api/adapter/controller"
)

func TodosCampos(w http.ResponseWriter, r *http.Request) {
	resp, statusCode := controller.TodosCampos()

	w.WriteHeader(statusCode)
	w.Write(resp)
}

func NovoCampo(w http.ResponseWriter, r *http.Request) {
	body, _ := io.ReadAll(r.Body)
	defer r.Body.Close()

	campo, statusCode := controller.NovoCampo(body)

	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(campo)
}

func PegandoCampoById(w http.ResponseWriter, r *http.Request) {
	paramsReplace := strings.Replace(r.RequestURI, "/", "", -1)
	paramId := strings.Replace(paramsReplace, "campo", "", -1)

	resp, statusCode := controller.PegandoCampoById(paramId)

	w.WriteHeader(statusCode)
	w.Write(resp)
}
