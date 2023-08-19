package campo

import "github.com/Bhimmo/golang-simple-api/internal/domain/entity/campo"

type InMemoryCampoRepository struct {
	Campo []any
}

func (r *InMemoryCampoRepository) Salvar(campo campo.Campo) error {
	r.Campo = append(r.Campo, campo)
	return nil
}
