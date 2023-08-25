package servico_campo

import (
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/campo"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/servico"
)

type ServicoCampo struct {
	Id      uint
	Servico servico.Servico
	Campos  []campo.Campo
}

func NewServicoCampo(servicoInput servico.Servico, campos []campo.Campo) *ServicoCampo {
	return &ServicoCampo{
		Servico: servicoInput,
		Campos:  campos,
	}
}
