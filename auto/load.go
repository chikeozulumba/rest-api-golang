package auto

import (
	"api/database"
	"api/models"
	"api/utils/console"
	"log"
)

var users = []models.User{
	models.User{
		Nickname: "Chike Arthur",
		Email: "chike@gmail.com",
		Password: "123456",
	},
}

func Load() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Debug().DropTableIfExists(&models.User{}).Error
	if err != nil {
		log.Fatal(err)
	}
	err = db.Debug().AutoMigrate(&models.User{}).Error
	if err != nil {
		log.Fatal(err)
	}

	for _, user := range users {
		err = db.Debug().Model(&models.User{}).Create(&user).Error
		if err != nil {
			log.Fatal(err)
		}
		console.Pretty(user)
	}
}