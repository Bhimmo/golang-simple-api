package servico_campo

import (
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/campo"
)

type ServicoCampoInterface interface {
	SalvarCampoNoServico(servicoId uint, entityCampo campo.Campo) error
	PegarCamposDoServico(servicoId uint) ([]campo.Campo, error)
	PegarCampoExistenteByCampoIdAndServicoId(campoId uint, servicoId uint) bool
}
