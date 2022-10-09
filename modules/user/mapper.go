package user

import "github.com/arifintahu/go-rest-api/entities"

func MappingUserOutput(user entities.User) entities.User {
	return entities.User{
		ID: user.ID,
		RoleID: user.RoleID,
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
