package role

import (
	"github.com/arifintahu/go-rest-api/dto"
	"github.com/arifintahu/go-rest-api/entities"
	"github.com/arifintahu/go-rest-api/modules/role/types"
	"github.com/arifintahu/go-rest-api/repositories"
	"github.com/arifintahu/go-rest-api/utils/pagination"
)

type UseCase struct {
	role repositories.IRoleRepository
}
type IUseCase interface {
	CreateRole(body *dto.RoleInput) (*entities.Role, error)
	GetRoles(query *dto.RoleListQuery) (*[]entities.Role, int64, error)
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

func (uc UseCase) GetRoles(query *dto.RoleListQuery) (*[]entities.Role, int64, error) {
	offset, limit := pagination.OffsetAndLimit(query.Page, query.Limit)
	params :=  dto.RoleListParams{
		Offset: offset,
		Limit: limit,
	}

	roles, err := uc.role.GetRoles(&params)
	if err != nil {
		return &[]entities.Role{}, 0, err
	}

	total, err := uc.role.GetRolesTotal()
	if err != nil {
		return &[]entities.Role{}, 0, err
	}
	
	return roles, total, nil
}
