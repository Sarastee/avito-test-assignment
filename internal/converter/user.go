package converter

import "github.com/sarastee/avito-test-assignment/internal/model"

func CreateUserToUser(userID int64, createUser *model.CreateUser) *model.User {
	return &model.User{
		ID:       userID,
		Name:     createUser.Name,
		Role:     createUser.Role,
		Password: createUser.Password,
	}
}
