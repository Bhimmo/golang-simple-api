package pegando_campo_pelo_id

import "github.com/Bhimmo/golang-simple-api/internal/domain/entity/campo"

type PegandoCampoPeloId struct {
	repositoryCampo campo.InterfaceCampoRepository
}

func NewPegandoCampoPeloId(
	campoRepository campo.InterfaceCampoRepository,
) *PegandoCampoPeloId {
	return &PegandoCampoPeloId{
		repositoryCampo: campoRepository,
	}
}

func (u *PegandoCampoPeloId) Execute(
	input PegandoCampoPeloIdInput,
) (PegandoCampoPeloIdOutput, error) {
	campoBusca, errRepository := u.repositoryCampo.BuscarPeloId(input.Id)
	if errRepository != nil {
		return PegandoCampoPeloIdOutput{}, errRepository
	}

	return PegandoCampoPeloIdOutput{
		Id:   campoBusca.Id,
		Nome: campoBusca.Nome,
	}, nil
}
