package servico

import (
	"database/sql"
	"errors"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/servico"
)

type repositoryServico struct {
	db *sql.DB
}

func NovoRepositoryServico(database *sql.DB) *repositoryServico {
	return &repositoryServico{db: database}
}

func (r *repositoryServico) Inserir(nome string) (uint, error) {
	stmt, errPrepare := r.db.Prepare("INSERT INTO servico(nome) VALUES (?)")
	if errPrepare != nil {
		panic(errPrepare)
	}

	res, errExec := stmt.Exec(nome)
	if errExec != nil {
		panic(errExec)
	}

	id, errId := res.LastInsertId()
	if errId != nil {
		panic(errId)
	}

	return uint(id), nil
}

func (r *repositoryServico) PegandoPeloId(id uint) (servico.Servico, error) {
	var servicoOutput servico.Servico
	row := r.db.QueryRow("SELECT * FROM servico WHERE id = ?", id)
	errScan := row.Scan(&servicoOutput.Id, &servicoOutput.Nome)
	if errScan != nil {
		return servico.Servico{}, errors.New("Erro na leitura dos dados")
	}

	return servicoOutput, nil
}
