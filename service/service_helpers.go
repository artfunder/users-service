package service

import "github.com/artfunder/structs"

func checkUserHasRequiredFields(userInfo structs.User) error {
	if userInfo.Username == "" {
		return ErrMustHaveUsername
	}

	if userInfo.Email == "" {
		return ErrMustHaveEmail
	}

	if userInfo.Password == "" {
		return ErrInvalidPassword
	}

	return nil
}

func withoutPassword(user structs.User) structs.User {
	user.Password = ""
	return user
}

func removePasswords(users []structs.User) {
	for i := range users {
		users[i].Password = ""
	}
}
