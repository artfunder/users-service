package test

import (
	"errors"

	"github.com/artfunder/structs"
	"github.com/artfunder/users-service/repository"
)

// MockRepo ...
type MockRepo struct {
	nextID int
}

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

// Create ...
func (r MockRepo) Create(user structs.User) (structs.User, error) {
	user.ID = r.nextID
	return user, nil
}

// SetNextID ...
func (r *MockRepo) SetNextID(id int) *MockRepo {
	r.nextID = id
	return r
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

// Create ...
func (BadRepo) Create(user structs.User) (structs.User, error) {
	return structs.User{}, ErrBad
}

// ErrBad ...
var ErrBad = errors.New("Big ol' error")
