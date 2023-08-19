package campo

import "github.com/Bhimmo/golang-simple-api/internal/domain/entity/campo"

type InMemoryCampoRepository struct {
	Campo []any
}

func (r *InMemoryCampoRepository) Salvar(campo campo.Campo) error {
	r.Campo = append(r.Campo, campo)
	return nil
}

func (r *InMemoryCampoRepository) BuscarCampoPeloSolicitanteId(
	solicitanteId uint,
) ([]campo.Campo, error) {
	var campos []campo.Campo
	campos = append(campos, *campo.NovoCampo(1, "TESTE", 2))
	return campos, nil
}
