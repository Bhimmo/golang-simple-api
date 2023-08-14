package pegando_pelo_id

import (
	"errors"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/servico"
)

type pegandoPeloIdUseCase struct {
	repository servico.InterfaceServicoRepository
}

func NovoPegandoPeloId(repositoryInput servico.InterfaceServicoRepository) *pegandoPeloIdUseCase {
	return &pegandoPeloIdUseCase{repository: repositoryInput}
}

func (p *pegandoPeloIdUseCase) Execute(input PegandoPeloIdInput) (servico.Servico, error) {
	s, err := p.repository.PegandoPeloId(input.Id)
	if err != nil {
		return servico.Servico{}, errors.New("Servico nao encontrado")
	}

	return s, nil
}
