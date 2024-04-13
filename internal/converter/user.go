package converter

import "github.com/sarastee/avito-test-assignment/internal/model"

// CreateUserToUser function which converts CreateUser to User model struct.
func CreateUserToUser(userID int64, createUser *model.CreateUser) *model.User {
	return &model.User{
		ID:       userID,
		Name:     createUser.Name,
		Role:     createUser.Role,
		Password: createUser.Password,
	}
}
