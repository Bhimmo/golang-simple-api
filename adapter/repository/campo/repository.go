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
func (r *RepositoryCampo) BuscarCampoPeloSolicitacaoId(solicitacaoId uint) ([]campo.Campo, error) {
	rows, errQuery := r.db.Query("SELECT * FROM campo WHERE solicitacaoId = ?", solicitacaoId)
	defer rows.Close()
	if errQuery != nil {
		return []campo.Campo{}, errors.New("Erro na consulta solcitacao campo")
	}

	var listaCampoRetornar []campo.Campo
	for rows.Next() {
		var itemListaCampo campo.Campo
		errScan := rows.Scan(&itemListaCampo.Id, &itemListaCampo.Valor, &itemListaCampo.SolicitacaoId)
		if errScan != nil {
			return []campo.Campo{}, errors.New("erro no tranformar solicitacao campo")
		}
		listaCampoRetornar = append(listaCampoRetornar, itemListaCampo)
	}
	return listaCampoRetornar, nil
}
