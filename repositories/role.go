package repositories

import (
	"github.com/arifintahu/go-rest-api/dto"
	"github.com/arifintahu/go-rest-api/entities"
	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

type IRoleRepository interface {
	CreateRole(role *entities.Role) (*entities.Role, error)
	GetRoles(params *dto.RoleListParams) (*[]entities.Role, error)
	GetRolesTotal() (int64, error)
	GetRoleBySlug(slug string) (*entities.Role, error)
}

func NewRoleRepository(db *gorm.DB) IRoleRepository {
	return &RoleRepository{db}
}

func (repo *RoleRepository) CreateRole(role *entities.Role) (*entities.Role, error) {
	err := repo.db.Create(role).Take(&role).Error
	return role, err
}

func (repo *RoleRepository) GetRoles(params *dto.RoleListParams) (*[]entities.Role, error) {
	roles := []entities.Role{}
	err := repo.db.
			Model(&entities.Role{}).
			Offset(params.Offset).
			Limit(params.Limit).
			Find(&roles).
			Error
	return &roles, err
}

func (repo *RoleRepository) GetRolesTotal() (int64, error) {
	var total int64
	err := repo.db.
			Model(&entities.Role{}).
			Count(&total).
			Error

	return total, err
}

func (repo *RoleRepository) GetRoleBySlug(slug string) (*entities.Role, error) {
	role := entities.Role{}
	err := repo.db.
			Model(&entities.Role{}).
			Where("slug = ?", slug).
			Take(&role).
			Error
	return &role, err
}
