package service_test

import (
	"testing"

	"github.com/artfunder/structs"
	"github.com/artfunder/users-service/repository"
	"github.com/artfunder/users-service/service"
	"github.com/artfunder/users-service/test"
	"github.com/gomagedon/expectate"
)

type CreateTest struct {
	name         string
	repo         repository.UsersRepository
	inputUser    structs.User
	expectedErr  error
	expectedUser structs.User
}

var createTests = []CreateTest{
	{
		name: "Succeeds with good repo and unique user",
		repo: new(test.MockRepo).SetNextID(4),
		inputUser: structs.User{
			Firstname: "Philip",
			Lastname:  "Fry",
			Username:  "fryguy",
			Email:     "philip.j.fry@planetexpress.net",
			Password:  "ihatemyjob",
		},
		expectedErr: nil,
		expectedUser: structs.User{
			ID:        4,
			Firstname: "Philip",
			Lastname:  "Fry",
			Username:  "fryguy",
			Email:     "philip.j.fry@planetexpress.net",
		},
	},
	{
		name: "Succeeds with good repo and different ID",
		repo: new(test.MockRepo).SetNextID(11),
		inputUser: structs.User{
			Firstname: "Philip",
			Lastname:  "Fry",
			Username:  "fryguy",
			Email:     "philip.j.fry@planetexpress.net",
			Password:  "ihatemyjob",
		},
		expectedErr: nil,
		expectedUser: structs.User{
			ID:        11,
			Firstname: "Philip",
			Lastname:  "Fry",
			Username:  "fryguy",
			Email:     "philip.j.fry@planetexpress.net",
		},
	},
	{
		name: "Succeeds without Firstname",
		repo: new(test.MockRepo).SetNextID(4),
		inputUser: structs.User{
			Lastname: "Doe",
			Username: "blank_doe",
			Email:    "mr.doe@nobody.com",
			Password: "1234",
		},
		expectedErr: nil,
		expectedUser: structs.User{
			ID:       4,
			Lastname: "Doe",
			Username: "blank_doe",
			Email:    "mr.doe@nobody.com",
		},
	},
	{
		name: "Succeeds without Lastname",
		repo: new(test.MockRepo).SetNextID(4),
		inputUser: structs.User{
			Firstname: "John",
			Username:  "johnny123",
			Email:     "johnny123@email.com",
			Password:  "1234",
		},
		expectedErr: nil,
		expectedUser: structs.User{
			ID:        4,
			Firstname: "John",
			Username:  "johnny123",
			Email:     "johnny123@email.com",
		},
	},
	{
		name: "Fails without Username",
		repo: new(test.MockRepo).SetNextID(4),
		inputUser: structs.User{
			Firstname: "John",
			Lastname:  "Doe",
			Email:     "john_doe@example.com",
			Password:  "1234",
		},
		expectedErr:  service.ErrMustHaveUsername,
		expectedUser: structs.User{},
	},
}

func TestCreate(t *testing.T) {
	for _, tc := range createTests {
		t.Run(tc.name, func(t *testing.T) {
			runCreateTest(t, tc)
		})
	}
}

func runCreateTest(t *testing.T, tc CreateTest) {
	expect := expectate.Expect(t)

	usersService := service.NewUsersService(tc.repo)
	user, err := usersService.Create(tc.inputUser)

	expect(err).ToBe(tc.expectedErr)
	expect(user).ToEqual(tc.expectedUser)
}
