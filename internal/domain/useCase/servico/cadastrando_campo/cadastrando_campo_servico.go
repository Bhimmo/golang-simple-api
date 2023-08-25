package cadastrando_campo_servico

import (
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/campo"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/servico"
	"github.com/Bhimmo/golang-simple-api/internal/domain/entity/servico_campo"
)

type CadastrandoCampoServico struct {
	repositoryServicoCampo servico_campo.ServicoCampoInterface
	repositoryServico      servico.InterfaceServicoRepository
	repositoryCampo        campo.InterfaceCampoRepository
}

func NewCadastrandoCampoServico(
	servicoCampoRepository servico_campo.ServicoCampoInterface,
	servicoRepository servico.InterfaceServicoRepository,
	campoRepository campo.InterfaceCampoRepository,
) *CadastrandoCampoServico {
	return &CadastrandoCampoServico{
		repositoryServicoCampo: servicoCampoRepository,
		repositoryServico:      servicoRepository,
		repositoryCampo:        campoRepository,
	}
}

func (u *CadastrandoCampoServico) Execute(input CadastrandoCampoServicoInput) error {
	servicoBusca, errServico := u.repositoryServico.PegandoPeloId(input.ServicoId)
	if errServico != nil {
		return errServico
	}

	for _, idCampo := range input.Campos {
		EntityCampoBusca, errCampo := u.repositoryCampo.BuscarPeloId(idCampo)
		if errCampo != nil {
			return errCampo
		}

		//Verificando se o campo ja esta cadastrado
		campoExiste := u.repositoryServicoCampo.PegarCampoExistenteByCampoIdAndServicoId(
			EntityCampoBusca.Id, servicoBusca.Id,
		)
		if !campoExiste {
			errSalvar := u.repositoryServicoCampo.SalvarCampoNoServico(servicoBusca.Id, EntityCampoBusca)
			if errSalvar != nil {
				return errSalvar
			}
		}
	}
	return nil
}
