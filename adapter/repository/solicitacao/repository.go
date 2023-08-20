package solicitacao

import (
	"database/sql"
	"errors"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/solicitacao"
)

type RepositorySolicitacao struct {
	db *sql.DB
}

func NovoRepositorySolicitacao(database *sql.DB) *RepositorySolicitacao {
	return &RepositorySolicitacao{db: database}
}

func (r *RepositorySolicitacao) Salvar(
	servicoId uint,
	statusId uint,
	concluida bool,
	solicitanteId uint,
) (uint, error) {
	stmt, errPrepare := r.db.Prepare("INSERT INTO solicitacao (servicoId, statusId, concluida, solicitanteId) VALUES (?, ?, ?, ?)")
	if errPrepare != nil {
		return 0, errors.New("erro na preparacao: INSERT solicitacao")
	}

	res, errExec := stmt.Exec(servicoId, statusId, concluida, solicitanteId)
	if errExec != nil {
		return 0, errors.New("erro na execucao: INSERT solicitacao")
	}

	id, errId := res.LastInsertId()
	if errId != nil {
		return 0, errors.New("erro nos parametros do retorno exec solicitacao")
	}

	return uint(id), nil
}

func (r *RepositorySolicitacao) BuscarPeloId(id uint) (solicitacao.Solicitacao, error) {
	return solicitacao.Solicitacao{}, nil
}
