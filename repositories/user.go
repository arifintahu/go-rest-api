package repositories

import (
	"time"

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
			Model(&entities.User{}).
			Offset(params.Offset).
			Limit(params.Limit).
			Find(&users).
			Error
	return &users, err
}

func (repo *UserRepository) GetUsersTotal() (int64, error) {
	var total int64
	err := repo.db.
			Model(&entities.User{}).
			Count(&total).
			Error

	return total, err
}

func (repo *UserRepository) GetUserDetail(id uint64) (*entities.User, error) {
	user := entities.User{}
	err := repo.db.
			Model(&entities.User{}).
			Where("id = ?", id).
			Take(&user).
			Error
	return &user, err
}

func (repo *UserRepository) GetUserByEmail(email string) (*entities.User, error) {
	user := entities.User{}
	err := repo.db.
			Model(&entities.User{}).
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
			UpdateColumns(
				map[string]interface{}{
					"role_id": userUpdate.RoleID,
					"first_name": userUpdate.FirstName,
					"last_name": userUpdate.LastName,
					"updated_at": time.Now(),
				},
			).
			Take(&user).
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
