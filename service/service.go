package services

import "github.com/Api-Micro-Service/user"

var users = []models.User{
	{ID: 1, Name: "Noa", Email: "NoaTricerri@epitech.eu"},
	{ID: 2, Name: "Azad", Email: "AzadMakezian@epitech.eu"},
}

func GetUsers() []models.User {
	return users
}

func GetUserByID(id int) *models.User {
	for _, user := range users {
		if user.ID == id {
			return &user
		}
	}
	return nil
}