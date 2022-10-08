package repositories

import (
	"github.com/arifintahu/go-rest-api/entities"
	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

type IRoleRepository interface {
	CreateRole(role *entities.Role) (*entities.Role, error)
	GetRoles() (*[]entities.Role, error)
	GetRoleBySlug(slug string) (*entities.Role, error)
}

func NewRoleRepository(db *gorm.DB) IRoleRepository {
	return &RoleRepository{db}
}

func (repo *RoleRepository) CreateRole(role *entities.Role) (*entities.Role, error) {
	err := repo.db.Create(role).Take(&role).Error
	return role, err
}

func (repo *RoleRepository) GetRoles() (*[]entities.Role, error) {
	roles := []entities.Role{}
	err := repo.db.
			Model(&entities.Role{}).
			Limit(100).
			Find(&roles).
			Error
	return &roles, err
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
