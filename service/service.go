package service

import (
	"github.com/artfunder/structs"
	"github.com/artfunder/users-service/repository"
)

// NewUsersService ...
func NewUsersService(repo repository.UsersRepository) *UsersService {
	service := new(UsersService)
	service.repo = repo
	return service
}

// UsersService ...
type UsersService struct {
	repo repository.UsersRepository
}

// GetAll ...
func (service UsersService) GetAll() ([]structs.User, error) {
	users, err := service.repo.GetAll()
	if err != nil {
		return []structs.User{}, err
	}
	for i := range users {
		users[i].Password = ""
	}
	return users, nil
}

// GetOne ...
func (service UsersService) GetOne(id int) (structs.User, error) {
	user, err := service.repo.GetOne(id)
	if err != nil {
		return structs.User{}, err
	}
	user.Password = ""
	return user, nil
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
