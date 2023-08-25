package servico

import (
	"database/sql"
	"errors"

	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/servico"
)

type RepositoryServico struct {
	db *sql.DB
}

func NovoRepositoryServico(database *sql.DB) *RepositoryServico {
	return &RepositoryServico{db: database}
}

func (r *RepositoryServico) Inserir(nome string) (uint, error) {
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

func (r *RepositoryServico) PegandoPeloId(id uint) (servico.Servico, error) {
	var servicoOutput servico.Servico
	row := r.db.QueryRow("SELECT * FROM servico WHERE id = ?", id)
	errScan := row.Scan(&servicoOutput.Id, &servicoOutput.Nome)
	if errScan != nil {
		return servico.Servico{}, errors.New("servico nao encontrado")
	}

	return servicoOutput, nil
}
