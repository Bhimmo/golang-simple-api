package solicitacao_campo

import (
	"database/sql"
	"errors"

	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/solicitacao_campo"
)

type RepositorySolicitacaoCampo struct {
	db *sql.DB
}

func NewRepositorySolicitacaoCampo(database *sql.DB) *RepositorySolicitacaoCampo {
	return &RepositorySolicitacaoCampo{
		db: database,
	}
}

func (r *RepositorySolicitacaoCampo) SalvarCamposDaSolicitacao(
	campoId uint,
	solicitacaoId uint,
	valor string,
) error {
	stmt, errPrepare := r.db.Prepare("INSERT INTO solicitacao_campo (campoId, solicitacaoId, valor) VALUES (?, ?, ?)")
	if errPrepare != nil {
		return errors.New("erro na preparacao do insert: solicitacao campo")
	}

	_, errExec := stmt.Exec(campoId, solicitacaoId, valor)
	if errExec != nil {
		return errors.New("erro na execucao: save solicitacao campo")
	}
	return nil
}

func (r *RepositorySolicitacaoCampo) BuscarCamposPelaSolicitacao(
	solicitacaoId uint,
) ([]solicitacao_campo.SolicitacaoCampo, error) {
	return nil, nil
}
