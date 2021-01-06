package test

import "github.com/artfunder/structs"

// SampleUsers ...
var SampleUsers = []structs.User{
	{
		ID:        1,
		Firstname: "John",
		Lastname:  "Doe",
		Username:  "johndoe",
		Email:     "johndoe@example.com",
		Password:  "1234",
	},
	{
		ID:        2,
		Firstname: "Jane",
		Lastname:  "Doe",
		Username:  "janedoe",
		Email:     "janedoe@example.com",
		Password:  "4321",
	},
	{
		ID:        3,
		Firstname: "Bobby",
		Lastname:  "Tables",
		Username:  "dropusers",
		Email:     "bobthetablebuilder@sqlinjection.com",
		Password:  "supersecret",
	},
}

// SampleUsersWithoutPasswords ...
var SampleUsersWithoutPasswords = []structs.User{
	{
		ID:        1,
		Firstname: "John",
		Lastname:  "Doe",
		Username:  "johndoe",
		Email:     "johndoe@example.com",
	},
	{
		ID:        2,
		Firstname: "Jane",
		Lastname:  "Doe",
		Username:  "janedoe",
		Email:     "janedoe@example.com",
	},
	{
		ID:        3,
		Firstname: "Bobby",
		Lastname:  "Tables",
		Username:  "dropusers",
		Email:     "bobthetablebuilder@sqlinjection.com",
	},
}
