package cadastrar_campo

import (
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/campo"
)

type CadastrarCampo struct {
	repositoryCampo campo.InterfaceCampoRepository
}

func NovoCadastrarCampo(
	campoRepository campo.InterfaceCampoRepository,
) *CadastrarCampo {
	return &CadastrarCampo{
		repositoryCampo: campoRepository,
	}
}

func (c *CadastrarCampo) Execute(input CadastrarCampoInput) (CadastrarCampoOutput, error) {
	EntityCampo := campo.NovoCampo()
	EntityCampo.Nome = input.Nome

	//Salvar banco
	returnIdCampo, errSalvar := c.repositoryCampo.Salvar(*EntityCampo)
	if errSalvar != nil {
		return CadastrarCampoOutput{}, errSalvar
	}
	EntityCampo.Id = returnIdCampo

	//retornar campo com id
	return CadastrarCampoOutput{
		Id:   EntityCampo.Id,
		Nome: EntityCampo.Nome,
	}, nil
}
