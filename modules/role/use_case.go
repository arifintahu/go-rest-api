package role

import (
	"github.com/arifintahu/go-rest-api/dto"
	"github.com/arifintahu/go-rest-api/entities"
	"github.com/arifintahu/go-rest-api/modules/role/types"
	"github.com/arifintahu/go-rest-api/repositories"
)

type UseCase struct {
	role repositories.IRoleRepository
}
type IUseCase interface {
	CreateRole(body *dto.RoleInput) (*entities.Role, error)
	GetRoles() (*[]entities.Role, error)
}

var _ IUseCase = (*UseCase)(nil)

func (uc UseCase) CreateRole(body *dto.RoleInput) (*entities.Role, error) {
	existRole, _ := uc.role.GetRoleBySlug(body.Slug)

	if (existRole.ID != 0) {
		return &entities.Role{}, types.ErrRoleSlugExist
	}

	role := entities.Role{
		Name: body.Name,
		Slug: body.Slug,
	}

	return uc.role.CreateRole(&role)
}

func (uc UseCase) GetRoles() (*[]entities.Role, error) {
	return uc.role.GetRoles()
}
