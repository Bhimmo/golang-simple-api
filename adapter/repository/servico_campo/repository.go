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
	rows, errRows := r.db.Query(`
		select 
			c.* 
		from servico_campo sc 
			join campo c on c.id = sc.campoId 
		where 
			sc.servicoId = ?`,
		servicoId)

	if errRows != nil {
		return nil, errors.New("campos do servico: servico invalido")
	}

	var listaCamposRetorno []campo.Campo
	for rows.Next() {
		var campoLista campo.Campo
		errScan := rows.Scan(&campoLista.Id, &campoLista.Nome)
		if errScan != nil {
			return nil, errScan
		}
		listaCamposRetorno = append(listaCamposRetorno, campoLista)
	}

	return listaCamposRetorno, nil
}

func (r *RepostioryServicoCampo) PegarCampoExistenteByCampoIdAndServicoId(
	campoId uint,
	servicoId uint,
) bool {
	row := r.db.QueryRow("SELECT * FROM servico_campo where campoId = ? AND servicoId = ?", campoId, servicoId)

	var idServicoCampo uint
	var ServicoCampoServicoId uint
	var ServicoCampoCampoId uint
	err := row.Scan(&idServicoCampo, &ServicoCampoServicoId, &ServicoCampoCampoId)
	if err != nil {
		return false
	}
	return true
}
