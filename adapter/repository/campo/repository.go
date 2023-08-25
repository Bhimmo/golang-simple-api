package campo

import (
	"database/sql"
	"errors"

	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/campo"
)

type RepositoryCampo struct {
	db *sql.DB
}

func NovoRepositoryCampo(database *sql.DB) *RepositoryCampo {
	return &RepositoryCampo{db: database}
}

func (r *RepositoryCampo) Salvar(campo campo.Campo) (uint, error) {
	stmt, errPrepare := r.db.Prepare("INSERT INTO campo (nome) VALUES (?)")
	if errPrepare != nil {
		return 0, errors.New("erro na preparacao: campo")
	}

	res, errExec := stmt.Exec(campo.Nome)
	if errExec != nil {
		return 0, errors.New("erro na execucao: campo")
	}

	id, errId := res.LastInsertId()
	if errId != nil {
		return 0, errors.New("erro na retirada do Id: campo")
	}

	return uint(id), nil
}

func (r *RepositoryCampo) BuscarPeloId(id uint) (campo.Campo, error) {
	var campoRetorno campo.Campo
	row := r.db.QueryRow("SELECT * FROM campo WHERE id = ?", id)

	errScan := row.Scan(&campoRetorno.Id, &campoRetorno.Nome)
	if errScan != nil {
		return campo.Campo{}, errors.New("campo nao encontrado")
	}

	return campoRetorno, nil
}

func (r *RepositoryCampo) BuscarTodos() ([]campo.Campo, error) {
	rows, errQuery := r.db.Query("SELECT * FROM campo")
	if errQuery != nil {
		return nil, errQuery
	}

	var campoRetorno []campo.Campo
	for rows.Next() {
		var campoLista campo.Campo
		errScan := rows.Scan(&campoLista.Id, &campoLista.Nome)
		if errScan != nil {
			return nil, errScan
		}
		campoRetorno = append(campoRetorno, campoLista)
	}
	return campoRetorno, nil
}
