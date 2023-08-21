package solicitacao

import (
	"database/sql"
	"errors"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/servico"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/solicitacao"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/status"
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
	rowSolicitacao := r.db.QueryRow("SELECT * FROM solicitacao WHERE id = ?", id)

	var idSolicitacao uint
	var servicoId uint
	var statusId uint
	var concluida bool
	var solicitanteId uint
	errScan := rowSolicitacao.Scan(&idSolicitacao, &servicoId, &statusId, &concluida, &solicitanteId)
	if errScan != nil {
		return solicitacao.Solicitacao{}, errors.New("Erro para busca informacoes solicitacao")
	}

	//Servico
	var nomeServico string
	rowServico := r.db.QueryRow("SELECT * FROM servico WHERE id = ?", servicoId)
	errScanServico := rowServico.Scan(&servicoId, &nomeServico)
	if errScanServico != nil {
		return solicitacao.Solicitacao{}, errors.New("Erro para buscar informacoes do servico")
	}
	se := servico.Servico{Id: servicoId, Nome: nomeServico}

	//Status
	st := status.NovoStatus()
	st.TendoStatusDesejado(statusId)

	//Solicitacao retorno
	ss := solicitacao.NovaSolicitacao(se, st, concluida, solicitanteId)
	ss.SetandoId(idSolicitacao)

	return *ss, nil
}
