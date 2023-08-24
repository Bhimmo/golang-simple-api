package todos_campos

import "github.com/Bhimmo/golang-simple-api/internal/domain/entity/campo"

type TodosCampos struct {
	repositoryCampo campo.InterfaceCampoRepository
}

func NewTodosCampos(campoRepository campo.InterfaceCampoRepository) *TodosCampos {
	return &TodosCampos{
		repositoryCampo: campoRepository,
	}
}

func (u *TodosCampos) Execute() ([]TodosCamposOutput, error) {
	campos, errRepository := u.repositoryCampo.BuscarTodos()
	if errRepository != nil {
		return nil, errRepository
	}

	var listaOutput []TodosCamposOutput
	for _, item := range campos {
		listaOutput = append(listaOutput, TodosCamposOutput{
			Id:   item.Id,
			Nome: item.Nome,
		})
	}

	return listaOutput, nil
}
