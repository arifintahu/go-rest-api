package repositories

import (
	"github.com/arifintahu/go-rest-api/dto"
	"github.com/arifintahu/go-rest-api/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

type IUserRepository interface {
	CreateUser(user *entities.User) (*entities.User, error)
	GetUsers(params *dto.UserListParams) (*[]entities.User, error)
	GetUsersTotal() (int64, error)
	GetUserDetail(id uint64) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
	UpdateUser(id uint64, user *entities.User) (*entities.User, error)
	DeleteUser(id uint64) (error)
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{db}
}

func (repo *UserRepository) CreateUser(user *entities.User) (*entities.User, error) {
	err := repo.db.
			Create(user).
			Take(&user).
			Error
	return user, err
}

func (repo *UserRepository) GetUsers(params *dto.UserListParams) (*[]entities.User, error) {
	users := []entities.User{}
	err := repo.db.
			Offset(params.Offset).
			Limit(params.Limit).
			Find(&users).
			Error
	return &users, err
}

func (repo *UserRepository) GetUsersTotal() (int64, error) {
	user := entities.User{}
	var total int64
	err := repo.db.
			Find(&user).
			Count(&total).
			Error

	return total, err
}

func (repo *UserRepository) GetUserDetail(id uint64) (*entities.User, error) {
	user := entities.User{}
	err := repo.db.
			Joins("Role").
			Where("users.id = ?", id).
			Omit("users.password").
			Take(&user).
			Error
	return &user, err
}

func (repo *UserRepository) GetUserByEmail(email string) (*entities.User, error) {
	user := entities.User{}
	err := repo.db.
			Joins("Role").
			Where("users.email = ?", email).
			Take(&user).
			Error
	return &user, err
}

func (repo *UserRepository) UpdateUser(id uint64, userUpdate *entities.User) (*entities.User, error) {
	user := entities.User{}
	err := repo.db.
			Where("id = ?", id).
			UpdateColumns(userUpdate).
			Omit("password").
			Take(&user).
			Error

	return &user, err
}

func (repo *UserRepository) DeleteUser(id uint64) (error) {
	user := entities.User{}
	err := repo.db.
			Where("id = ?", id).
			Delete(&user).
			Error
	return err
}
