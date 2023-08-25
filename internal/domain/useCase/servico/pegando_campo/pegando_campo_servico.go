package pegando_campo_servico

import (
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/servico_campo"
)

type PegandoCampoServico struct {
	repositoryServico servico_campo.ServicoCampoInterface
}

func NewPegandoCampoServico(servicoRepository servico_campo.ServicoCampoInterface) *PegandoCampoServico {
	return &PegandoCampoServico{
		repositoryServico: servicoRepository,
	}
}

func (u *PegandoCampoServico) Execute(input PegandoCampoServicoInput) (PegandoCampoServicoOutput, error) {
	campos, errCampos := u.repositoryServico.PegarCamposDoServico(input.ServicoId)
	if errCampos != nil {
		return PegandoCampoServicoOutput{}, errCampos
	}

	var listaCamposReturn []PegandoCampoServicoCampoOutput
	for _, itemCampo := range campos {
		listaCamposReturn = append(listaCamposReturn, PegandoCampoServicoCampoOutput{
			Id:   itemCampo.Id,
			Nome: itemCampo.Nome,
		})
	}

	return PegandoCampoServicoOutput{
		ServicoId: input.ServicoId,
		Campos:    listaCamposReturn,
	}, nil
}
