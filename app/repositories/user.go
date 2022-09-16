package repositories

import (
	"time"

	"github.com/arifintahu/go-rest-api/app/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type IUserRepository interface {
	CreateUser(user *models.User) (error)
	GetUsers() (*[]models.User, error)
	GetUserDetail(id uint64) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
	UpdateUser(id uint64, user *models.User) (*models.User, error)
	DeleteUser(id uint64) (error)
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db}
}

func (repo *UserRepository) CreateUser(user *models.User) (error) {
	return repo.db.Create(user).Error
}

func (repo *UserRepository) GetUsers() (*[]models.User, error) {
	users := []models.User{}
	err := repo.db.
			Model(&models.User{}).
			Limit(100).
			Find(&users).
			Error
	return &users, err
}

func (repo *UserRepository) GetUserDetail(id uint64) (*models.User, error) {
	user := models.User{}
	err := repo.db.
			Model(&models.Role{}).
			Where("id = ?", id).
			Take(&user).
			Error
	return &user, err
}

func (repo *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	user := models.User{}
	err := repo.db.
			Model(&models.Role{}).
			Where("email = ?", email).
			Take(&user).
			Error
	return &user, err
}

func (repo *UserRepository) UpdateUser(id uint64, userUpdate *models.User) (*models.User, error) {
	user := models.User{}
	err := repo.db.
			Model(&models.User{}).
			Where("id = ?", id).
			Take(&user).
			UpdateColumns(
				map[string]interface{}{
					"role_id": userUpdate.RoleID,
					"first_name": userUpdate.FirstName,
					"last_name": userUpdate.LastName,
					"updated_at": time.Now(),
				},
			).
			Error

	return &user, err
}

func (repo *UserRepository) DeleteUser(id uint64) (error) {
	user := models.User{}
	err := repo.db.
			Model(&models.User{}).
			Where("id = ?", id).
			Take(&user).
			Delete(&user).
			Error
	return err
}
