package repositories

import (
	"github.com/arifintahu/go-rest-api/app/models"
	"gorm.io/gorm"
)

type RoleRepository struct {
	db *gorm.DB
}

type IRoleRepository interface {
	CreateRole(role *models.Role) (*models.Role, error)
	GetRoles() (*[]models.Role, error)
	GetRoleBySlug(slug string) (*models.Role, error)
}

func NewRoleRepository(db *gorm.DB) IRoleRepository {
	return &RoleRepository{db}
}

func (repo *RoleRepository) CreateRole(role *models.Role) (*models.Role, error) {
	err := repo.db.Create(role).Take(&role).Error
	return role, err
}

func (repo *RoleRepository) GetRoles() (*[]models.Role, error) {
	roles := []models.Role{}
	err := repo.db.
			Model(&models.Role{}).
			Limit(100).
			Find(&roles).
			Error
	return &roles, err
}

func (repo *RoleRepository) GetRoleBySlug(slug string) (*models.Role, error) {
	role := models.Role{}
	err := repo.db.
			Model(&models.Role{}).
			Where("slug = ?", slug).
			Take(&role).
			Error
	return &role, err
}
