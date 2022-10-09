package user

import (
	"github.com/arifintahu/go-rest-api/dto"
	"github.com/arifintahu/go-rest-api/entities"
	"github.com/arifintahu/go-rest-api/modules/user/types"
	"github.com/arifintahu/go-rest-api/repositories"
	"github.com/arifintahu/go-rest-api/utils/bcrypt"
)

type UseCase struct {
	user repositories.IUserRepository
}

type IUseCase interface {
	CreateUser(params *dto.UserInput) (*entities.User, error)
	GetUsers() (*[]entities.User, error)
	GetUserDetail(id uint64) (*entities.User, error)
	UpdateUser(id uint64, params *dto.UserUpdate) (*entities.User, error)
	DeleteUser(id uint64) (error)
}

var _ IUseCase = (*UseCase) (nil)

func (uc UseCase) CreateUser(params *dto.UserInput) (*entities.User, error) {
	existUser, _ := uc.user.GetUserByEmail(params.Email)

	if (existUser.ID != 0) {
		return &entities.User{}, types.ErrUserEmailExist
	}

	hashedPassword, err := bcrypt.HashPassword(params.Password)

	if err != nil {
		return &entities.User{}, err
	}

	user := entities.User{
		RoleID: params.RoleID,
		FirstName: params.FirstName,
		LastName: params.LastName,
		Email: params.Email,
		Password: hashedPassword,
	}

	return uc.user.CreateUser(&user)
}

func (uc UseCase) GetUsers() (*[]entities.User, error) {
	users, err := uc.user.GetUsers()

	if err != nil {
		return users, err
	}

	var mappedUser []entities.User
	for _, row := range *users {
		r := MappingUserOutput(row)
		mappedUser = append(mappedUser, r)
	}

	return &mappedUser, nil
}

func (uc UseCase) GetUserDetail(id uint64) (*entities.User, error) {
	user, err := uc.user.GetUserDetail(id)

	if err != nil {
		return user, err
	}

	mappedUser := MappingUserOutput(*user)
	return &mappedUser, nil
}

func (uc UseCase) UpdateUser(id uint64, params *dto.UserUpdate) (*entities.User, error) {
	_, err := uc.user.GetUserDetail(id)
	if err != nil {
		return &entities.User{}, types.ErrUserNotFound
	}

	user := entities.User{
		RoleID: params.RoleID,
		FirstName: params.FirstName,
		LastName: params.LastName,
	}

	return uc.user.UpdateUser(id, &user)
}

func (uc UseCase) DeleteUser(id uint64) (error) {
	_, err := uc.user.GetUserDetail(id)
	if err != nil {
		return types.ErrUserNotFound
	}

	return uc.user.DeleteUser(id)
}
