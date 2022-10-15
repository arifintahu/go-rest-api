package user

import (
	"github.com/arifintahu/go-rest-api/dto"
	"github.com/arifintahu/go-rest-api/entities"
	"github.com/arifintahu/go-rest-api/modules/user/types"
	"github.com/arifintahu/go-rest-api/repositories"
	"github.com/arifintahu/go-rest-api/utils/bcrypt"
	"github.com/arifintahu/go-rest-api/utils/pagination"
)

type UseCase struct {
	user repositories.IUserRepository
}

type IUseCase interface {
	CreateUser(body *dto.UserInput) (*entities.User, error)
	GetUsers(query *dto.UserListQuery) (*[]types.UserResponse, int64, error)
	GetUserDetail(id uint64) (*entities.User, error)
	UpdateUser(id uint64, body *dto.UserUpdate) (*entities.User, error)
	DeleteUser(id uint64) (error)
}

var _ IUseCase = (*UseCase) (nil)

func (uc UseCase) CreateUser(body *dto.UserInput) (*entities.User, error) {
	existUser, _ := uc.user.GetUserByEmail(body.Email)

	if (existUser.ID != 0) {
		return &entities.User{}, types.ErrUserEmailExist
	}

	hashedPassword, err := bcrypt.HashPassword(body.Password)

	if err != nil {
		return &entities.User{}, err
	}

	user := entities.User{
		RoleID: body.RoleID,
		FirstName: body.FirstName,
		LastName: body.LastName,
		Email: body.Email,
		Password: hashedPassword,
	}

	return uc.user.CreateUser(&user)
}

func (uc UseCase) GetUsers(query *dto.UserListQuery) (*[]types.UserResponse, int64, error) {
	offset, limit := pagination.OffsetAndLimit(query.Page, query.Limit)
	params :=  dto.UserListParams{
		Offset: offset,
		Limit: limit,
	}
	
	users, err := uc.user.GetUsers(&params)
	if err != nil {
		return &[]types.UserResponse{}, 0, err
	}

	total, err := uc.user.GetUsersTotal()
	if err != nil {
		return &[]types.UserResponse{}, 0, err
	}

	var userResponses []types.UserResponse
	for _, row := range *users {
		r := MappingUserOutput(row)
		userResponses = append(userResponses, r)
	}

	return &userResponses, total, nil
}

func (uc UseCase) GetUserDetail(id uint64) (*entities.User, error) {
	user, err := uc.user.GetUserDetail(id)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (uc UseCase) UpdateUser(id uint64, body *dto.UserUpdate) (*entities.User, error) {
	_, err := uc.user.GetUserDetail(id)
	if err != nil {
		return &entities.User{}, types.ErrUserNotFound
	}

	user := entities.User{
		RoleID: body.RoleID,
		FirstName: body.FirstName,
		LastName: body.LastName,
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
