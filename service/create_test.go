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
}

func TestCreate(t *testing.T) {
	for _, tc := range createTests {
		t.Run(tc.name, func(t *testing.T) {
			testCreate(t, tc)
		})
	}
}

func testCreate(t *testing.T, tc CreateTest) {
	expect := expectate.Expect(t)

	usersService := service.NewUsersService(tc.repo)
	user, err := usersService.Create(tc.inputUser)

	expect(err).ToBe(tc.expectedErr)
	expect(user).ToEqual(tc.expectedUser)
}
