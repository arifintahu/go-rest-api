package account

import (
	"github.com/arifintahu/go-rest-api/dto"
	"github.com/arifintahu/go-rest-api/modules/account/types"
	"github.com/arifintahu/go-rest-api/repositories"
	"github.com/arifintahu/go-rest-api/utils/bcrypt"
	"github.com/arifintahu/go-rest-api/utils/jwt"
)

type UseCase struct {
	user repositories.IUserRepository
}

type IUseCase interface {
	Login(body *dto.AccountLogin) (dto.AccountLoginResponse, error)
}

var _ IUseCase = (*UseCase)(nil)

func (uc UseCase) Login(body *dto.AccountLogin) (dto.AccountLoginResponse, error) {
	user, _ := uc.user.GetUserByEmail(body.Email)

	if (user.ID == 0) {
		return dto.AccountLoginResponse{}, types.ErrAccountEmailNotFound
	}

	isValid := bcrypt.CheckPasswordHash(body.Password, user.Password)
	if !isValid {
		return dto.AccountLoginResponse{}, types.ErrAccountPasswordInvalid
	}

	token, err := jwt.GenerateJWT(user)
	if err != nil {
		return dto.AccountLoginResponse{}, err
	}

	return dto.AccountLoginResponse{
		Token: token,
	}, nil
}
