package main

import (
	"fmt"
	"net/http"
	"os"

	my_middleware "github.com/Bhimmo/golang-simple-api/adapter/middleware"
	"github.com/Bhimmo/golang-simple-api/adapter/routes"
	"github.com/Bhimmo/golang-simple-api/pkg/rabbitmq"
	"github.com/Bhimmo/golang-simple-api/pkg/sqlite"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	// Enviroment
	errEnv := godotenv.Load(".env")
	if errEnv != nil {
		panic("Erro ao carregar variaveis de ambiente")
	}
	//Init database conection
	sqlite.Init()

	//Init rabbimq
	rabbitmq.Init()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.AllowContentType("application/json"))
	r.Use(middleware.SetHeader("Content-Type", "application/json"))

	//Auth
	r.Post("/access-token", routes.AccessToken)
	//Routes
	configRoutes(r)

	addr := os.Getenv("PORT")
	if addr == "" {
		addr = ":3000"
	}
	fmt.Println("Start api: lister on -> " + addr)
	http.ListenAndServe(addr, r)
}

func configRoutes(r *chi.Mux) {
	r.Group(func(r chi.Router) {
		//Validacao do token
		r.Use(my_middleware.ValidToken)

		//Campos
		r.Get("/campo", routes.TodosCampos)
		r.Post("/campo", routes.NovoCampo)
		r.Get("/campo/{id}", routes.PegandoCampoById)

		//Servico
		r.Post("/servico", routes.NovoServico)
		r.Get("/servico/{id}", routes.PegandoServicoPeloId)
		r.Post("/servico/campos", routes.AdicionandoCampos)
		r.Get("/servico/{id}/campos", routes.PegandoCampos)

		//Solicitacao
		r.Post("/solicitacao", routes.SalvarSolicitacao)
		r.Get("/solicitacao/{id}", routes.PegandoSolicitacaoPeloId)
		r.Get("/solicitacao/{id}/atualizar-status", routes.AtualizarStatusSolicitacao)
	})
}
