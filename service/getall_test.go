package service_test

import (
	"testing"

	"github.com/artfunder/structs"
	"github.com/artfunder/users-service/service"
	"github.com/artfunder/users-service/test"
	"github.com/gomagedon/expectate"
)

func TestGetAll_SucceedsWithMockRepo(t *testing.T) {
	expect := expectate.Expect(t)

	usersService := service.NewUsersService(new(test.MockRepo))
	users, err := usersService.GetAll()

	expect(err).ToBe(nil)
	expect(users).ToEqual(test.SampleUsersWithoutPasswords)
}

func TestGetAll_FailsWithBadRepo(t *testing.T) {
	expect := expectate.Expect(t)

	usersService := service.NewUsersService(new(test.BadRepo))
	users, err := usersService.GetAll()

	expect(err).ToBe(test.ErrBad)
	expect(users).ToEqual([]structs.User{})
}
