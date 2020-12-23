package service

import "github.com/artfunder/structs"

// UsersService ...
type UsersService struct{}

// GetAll ...
func (UsersService) GetAll() ([]structs.User, error) {
	return []structs.User{}, nil
}

// GetOne ...
func (UsersService) GetOne(id int) (structs.User, error) {
	return structs.User{}, nil
}

// Create ...
func (UsersService) Create(userInfo structs.User) (structs.User, error) {
	return structs.User{}, nil
}

// Update ...
func (UsersService) Update(id int, userInfo map[string]interface{}) (structs.User, error) {
	return structs.User{}, nil
}

// Delete ...
func (UsersService) Delete(id int) (structs.User, error) {
	return structs.User{}, nil
}
