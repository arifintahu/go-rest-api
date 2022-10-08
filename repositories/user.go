package repositories

import (
	"time"

	"github.com/arifintahu/go-rest-api/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type IUserRepository interface {
	CreateUser(user *entities.User) (error)
	GetUsers() (*[]entities.User, error)
	GetUserDetail(id uint64) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	UpdateUser(id uint64, user *entities.User) (*entities.User, error)
	DeleteUser(id uint64) (error)
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db}
}

func (repo *UserRepository) CreateUser(user *entities.User) (error) {
	return repo.db.Create(user).Error
}

func (repo *UserRepository) GetUsers() (*[]entities.User, error) {
	users := []entities.User{}
	err := repo.db.
			Model(&entities.User{}).
			Limit(100).
			Find(&users).
			Error
	return &users, err
}

func (repo *UserRepository) GetUserDetail(id uint64) (*entities.User, error) {
	user := entities.User{}
	err := repo.db.
			Model(&entities.Role{}).
			Where("id = ?", id).
			Take(&user).
			Error
	return &user, err
}

func (repo *UserRepository) GetUserByEmail(email string) (*entities.User, error) {
	user := entities.User{}
	err := repo.db.
			Model(&entities.Role{}).
			Where("email = ?", email).
			Take(&user).
			Error
	return &user, err
}

func (repo *UserRepository) UpdateUser(id uint64, userUpdate *entities.User) (*entities.User, error) {
	user := entities.User{}
	err := repo.db.
			Model(&entities.User{}).
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
	user := entities.User{}
	err := repo.db.
			Model(&entities.User{}).
			Where("id = ?", id).
			Take(&user).
			Delete(&user).
			Error
	return err
}
