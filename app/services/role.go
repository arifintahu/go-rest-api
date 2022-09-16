package services

import (
	"errors"

	"github.com/arifintahu/go-rest-api/app/dto"
	"github.com/arifintahu/go-rest-api/app/models"
	"github.com/arifintahu/go-rest-api/app/repositories"
	"gorm.io/gorm"
)

type RoleService struct {
	db *gorm.DB
}

type IRoleService interface {
	CreateRole(params *dto.RoleInput) (*models.Role, error)
	GetRoles() (*[]models.Role, error)
}

func NewRoleService(db *gorm.DB) IRoleService {
	return &RoleService{db}
}

func (service *RoleService) CreateRole(params *dto.RoleInput) (*models.Role, error) {
	roleRepository := repositories.NewRoleRepository(service.db)
	existRole, _ := roleRepository.GetRoleBySlug(params.Slug)

	if (existRole.ID != 0) {
		return &models.Role{}, errors.New("Role slug is exist")
	}

	role := models.Role{
		Name: params.Name,
		Slug: params.Slug,
	}

	return roleRepository.CreateRole(&role)
}

func (service *RoleService) GetRoles() (*[]models.Role, error) {
	roleRepository := repositories.NewRoleRepository(service.db)
	return roleRepository.GetRoles()
}
