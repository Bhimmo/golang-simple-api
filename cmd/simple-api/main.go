package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.AllowContentType("application/json"))

	//Atendimento
	//r.Post("/atendimento", atendimento.NovoAtendimento)

	http.ListenAndServe(":3000", r)
}
