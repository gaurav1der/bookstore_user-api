package services

import (
	"github.com/gaurav1der/bookstore_user-api/domain/users"
	"github.com/gaurav1der/bookstore_user-api/utils/errors"
)

func GetUser(userID int64) (*users.User, *errors.RestErr) {
	// if userID <= 0 {
	// 	return nil, errors.NewBadRequestError("Invalid user id")
	// }

	result := &users.User{ID: userID}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}
func CreateUser(user users.User) (*users.User, *errors.RestErr) {

	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, nil
	}

	// user.Email = strings.TrimSpace(strings.ToLower(user.Email))
	// if user.Email == "" {
	// 	return nil, errors.NewBadRequestError("Invalid email address")
	// }
	return &user, nil
}
