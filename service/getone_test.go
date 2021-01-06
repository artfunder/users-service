package service_test

import (
	"testing"

	"github.com/artfunder/structs"
	"github.com/artfunder/users-service/repository"
	"github.com/artfunder/users-service/service"
	"github.com/artfunder/users-service/test"
	"github.com/gomagedon/expectate"
)

type GetOneTest struct {
	name         string
	repo         repository.UsersRepository
	id           int
	expectedErr  error
	expectedUser structs.User
}

var getOneTests = []GetOneTest{
	{
		name:         "Succeeds with good repo id 1",
		repo:         new(test.MockRepo),
		id:           1,
		expectedErr:  nil,
		expectedUser: test.SampleUsersWithoutPasswords[0],
	},
	{
		name:         "Succeeds with good repo id 2",
		repo:         new(test.MockRepo),
		id:           2,
		expectedErr:  nil,
		expectedUser: test.SampleUsersWithoutPasswords[1],
	},
	{
		name:         "Succeeds with good repo id 3",
		repo:         new(test.MockRepo),
		id:           3,
		expectedErr:  nil,
		expectedUser: test.SampleUsersWithoutPasswords[2],
	},
	{
		name:         "Fails with good repo id 4",
		repo:         new(test.MockRepo),
		id:           4,
		expectedErr:  repository.ErrNoUserWithID,
		expectedUser: structs.User{},
	},
	{
		name:         "Fails with good repo id 0",
		repo:         new(test.MockRepo),
		id:           0,
		expectedErr:  repository.ErrNoUserWithID,
		expectedUser: structs.User{},
	},
	{
		name:         "Fails with bad repo id 1",
		repo:         new(test.BadRepo),
		id:           1,
		expectedErr:  test.ErrBad,
		expectedUser: structs.User{},
	},
}

func TestGetOne(t *testing.T) {
	for _, tc := range getOneTests {
		t.Run(tc.name, func(t *testing.T) {
			expect := expectate.Expect(t)

			usersService := service.NewUsersService(tc.repo)
			user, err := usersService.GetOne(tc.id)

			expect(err).ToBe(tc.expectedErr)
			expect(user).ToEqual(tc.expectedUser)
		})
	}
}
