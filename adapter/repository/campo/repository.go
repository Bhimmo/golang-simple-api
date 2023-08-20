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

func (r *RepositoryCampo) Salvar(campo campo.Campo) error {
	stmt, errPrepare := r.db.Prepare("INSERT INTO campo (valor, solicitacaoId) VALUES (?, ?)")
	if errPrepare != nil {
		return errors.New("erro na preparacao: solicitacao campo")
	}

	_, errExec := stmt.Exec(campo.Valor, campo.SolicitacaoId)
	if errExec != nil {
		return errors.New("erro na execucao: solicitacao campo")
	}
	return nil
}
func (r *RepositoryCampo) BuscarCampoPeloSolicitanteId(solicitanteId uint) ([]campo.Campo, error) {
	return nil, nil
}
