package service

import "github.com/artfunder/structs"

// UsersService ...
type UsersService interface {
	GetAll() ([]structs.User, error)
	GetOne(id int) (structs.User, error)
	Create(userInfo structs.User) (structs.User, error)
	Update(id int, userInfo map[string]interface{}) (structs.User, error)
	Delete(id int) (structs.User, error)
}
