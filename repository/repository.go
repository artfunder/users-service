package repository

import "github.com/artfunder/structs"

// UsersRepository ...
type UsersRepository interface {
	GetAll() ([]structs.User, error)
	GetOne(id int) (structs.User, error)
	Create(structs.User) (structs.User, error)
}
