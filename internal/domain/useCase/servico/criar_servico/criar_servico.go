package criar_servico

import "github.com/Bhimmo/golang-simple-api/internal/domain/entity/servico"

type criarServicoUseCase struct {
	repository servico.InterfaceServicoRepository
}

func NovoCriarServicoUseCase(repositoryInput servico.InterfaceServicoRepository) *criarServicoUseCase {
	return &criarServicoUseCase{repository: repositoryInput}
}

func (c *criarServicoUseCase) Execute(input CriarServicoInput) (*servico.Servico, error) {
	s := servico.NovoServico()
	s.Nome = input.Nome

	idServico, err := c.repository.Inserir(s.Nome)
	if err != nil {
		panic("Erro na hora de salvar servico")
	}
	s.Id = idServico

	return s, nil
}
