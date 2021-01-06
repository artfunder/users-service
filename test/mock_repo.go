package test

import (
	"errors"

	"github.com/artfunder/structs"
	"github.com/artfunder/users-service/repository"
)

// MockRepo ...
type MockRepo struct{}

// GetAll ...
func (MockRepo) GetAll() ([]structs.User, error) {
	return SampleUsers, nil
}

// GetOne ...
func (MockRepo) GetOne(id int) (structs.User, error) {
	if id > len(SampleUsers) || id <= 0 {
		return structs.User{}, repository.ErrNoUserWithID
	}
	return SampleUsers[id-1], nil
}

// BadRepo ...
type BadRepo struct{}

// GetAll ...
func (BadRepo) GetAll() ([]structs.User, error) {
	return nil, ErrBad
}

// GetOne ...
func (BadRepo) GetOne(id int) (structs.User, error) {
	return SampleUsers[2], ErrBad
}

// ErrBad ...
var ErrBad = errors.New("Big ol' error")
