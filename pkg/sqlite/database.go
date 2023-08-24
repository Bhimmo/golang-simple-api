package sqlite

import (
	"database/sql"
	// Driver for SQLite3 Database
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func Init() {
	database, err := sql.Open("sqlite3", "database.db")
	errorDb(err)

	Db = database
	generateBaseDados()
}

func generateBaseDados() {
	//Generate tables
	_, errServico := Db.Exec("CREATE TABLE IF NOT EXISTS servico (id INTEGER PRIMARY KEY AUTOINCREMENT, nome TEXT)")
	_, errSolicitacao := Db.Exec("CREATE TABLE IF NOT EXISTS solicitacao (id INTEGER PRIMARY KEY AUTOINCREMENT, servicoId INTEGER, statusId INTEGER, concluida INTEGER, solicitanteId INTEGER, FOREIGN KEY (servicoId) REFERENCES servico(id))")
	_, errCampo := Db.Exec("CREATE TABLE IF NOT EXISTS campo (id INTEGER PRIMARY KEY AUTOINCREMENT, nome TEXT)")
	_, errCampoSolicitacao := Db.Exec("CREATE TABLE IF NOT EXISTS solicitacao_campo (id INTEGER PRIMARY KEY AUTOINCREMENT, campoId INTEGER, valor TEXT, solicitacaoId INTEGER, FOREIGN KEY (campoId) REFERENCES campo(id), FOREIGN KEY (solicitacaoId) REFERENCES solicitacao(id))")

	errorDb(errServico)
	errorDb(errSolicitacao)
	errorDb(errCampoSolicitacao)
	errorDb(errCampo)
}

func errorDb(err error) {
	if err != nil {
		panic(err)
	}
}
