package servico_campo

import (
	"database/sql"
	"errors"

	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/campo"
)

type RepostioryServicoCampo struct {
	db *sql.DB
}

func NewRepositoryServicoCampo(database *sql.DB) *RepostioryServicoCampo {
	return &RepostioryServicoCampo{
		db: database,
	}
}

func (r *RepostioryServicoCampo) SalvarCampoNoServico(servicoId uint, entityCampo campo.Campo) error {
	stmt, errPrepare := r.db.Prepare("INSERT INTO servico_campo (servicoId, campoId) VALUES (?, ?)")
	if errPrepare != nil {
		return errPrepare
	}

	_, errExec := stmt.Exec(servicoId, entityCampo.Id)
	if errExec != nil {
		return errors.New("erro sem salvar campo no servico")
	}

	return nil
}

func (r *RepostioryServicoCampo) PegarCamposDoServico(servicoId uint) ([]campo.Campo, error) {
	return nil, nil
}
