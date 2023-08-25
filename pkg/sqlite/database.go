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
	_, errSolicitacao := Db.Exec("CREATE TABLE IF NOT EXISTS solicitacao (id INTEGER PRIMARY KEY AUTOINCREMENT, servicoId INTEGER NOT NULL, statusId INTEGER NOT NULL, concluida INTEGER NOT NULL, solicitanteId INTEGER NOT NULL, createdAt DATETIME NOT NULL, FOREIGN KEY (servicoId) REFERENCES servico(id))")
	_, errCampo := Db.Exec("CREATE TABLE IF NOT EXISTS campo (id INTEGER PRIMARY KEY AUTOINCREMENT, nome TEXT)")
	_, errCampoSolicitacao := Db.Exec("CREATE TABLE IF NOT EXISTS solicitacao_campo (id INTEGER PRIMARY KEY AUTOINCREMENT, campoId INTEGER, valor TEXT, solicitacaoId INTEGER, FOREIGN KEY (campoId) REFERENCES campo(id), FOREIGN KEY (solicitacaoId) REFERENCES solicitacao(id))")
	_, errServicoCampo := Db.Exec("CREATE TABLE IF NOT EXISTS servico_campo (id INTEGER PRIMARY KEY AUTOINCREMENT, servicoId INTEGER, campoId INTEGER, FOREIGN KEY (campoId) REFERENCES campo(id), FOREIGN KEY (servicoId) REFERENCES servico(id))")

	errorDb(errServico)
	errorDb(errSolicitacao)
	errorDb(errCampoSolicitacao)
	errorDb(errCampo)
	errorDb(errServicoCampo)
}

func errorDb(err error) {
	if err != nil {
		panic(err)
	}
}
