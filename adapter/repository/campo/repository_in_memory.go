package campo

import (
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/campo"
)

type InMemoryCampoRepository struct {
	Campo []campo.Campo
}

func (r *InMemoryCampoRepository) Salvar(campo campo.Campo) (uint, error) {
	r.Campo = append(r.Campo, campo)
	return uint(len(r.Campo) - 1), nil
}

func (r *InMemoryCampoRepository) BuscarPeloId(id uint) (campo.Campo, error) {
	return r.Campo[id-1], nil
}
