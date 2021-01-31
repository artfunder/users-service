package repository

import "github.com/artfunder/structs"

// SQLiteRepo ...
type SQLiteRepo struct{}

// GetAll ...
func (SQLiteRepo) GetAll() ([]structs.User, error) {
	return nil, nil
}

// GetOne ...
func (SQLiteRepo) GetOne(id int) (structs.User, error) {
	return structs.User{}, ErrNoUserWithID
}

// Create ...
func (SQLiteRepo) Create(userInfo structs.User) (structs.User, error) {
	return structs.User{}, nil
}
