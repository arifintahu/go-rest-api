package user

import (
	"github.com/arifintahu/go-rest-api/modules/user/types"

	"github.com/arifintahu/go-rest-api/entities"
)

func MappingUserOutput(user entities.User) types.UserResponse {
	return types.UserResponse{
		ID: user.ID,
		RoleID: user.RoleID,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
	}
}
