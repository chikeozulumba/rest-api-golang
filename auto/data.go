package auto

import (
	"api/models"
)

var users = []models.User{
	models.User{
		Nickname: "Chike Arthur",
		Email: "chike@gmail.com",
		Password: "123456",
	},
}

var posts = []models.Post {
	{
		Title:     "Chike Golang tut",
		Body:   "Lorem ipsum",
	},
}