package sqlite

import (
	"database/sql"
	// Driver for SQLite3 Database
	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB

func Init() {
	database, err := sql.Open("sqlite3", "../../database.db")
	errorDb(err)

	Db = database
	generateBaseDados()
}

func generateBaseDados() {
	//Generate tables
	_, errServico := Db.Exec("CREATE TABLE IF NOT EXISTS servico (id INTEGER PRIMARY KEY AUTOINCREMENT, nome TEXT)")
	_, errSolicitacao := Db.Exec("CREATE TABLE IF NOT EXISTS solicitacao (id INTEGER PRIMARY KEY AUTOINCREMENT, servicoId INTEGER, statusId INTEGER, concluida INTEGER, solicitanteId INTEGER)")
	_, errCampo := Db.Exec("CREATE TABLE IF NOT EXISTS campo (id INTEGER PRIMARY KEY AUTOINCREMENT, valor TEXT, solicitacaoId INTEGER)")

	errorDb(errServico)
	errorDb(errSolicitacao)
	errorDb(errCampo)
}

func errorDb(err error) {
	if err != nil {
		panic(err)
	}
}
