package auto

import (
	"api/database"
	"api/models"
	"api/utils/console"
	"log"
)

func Load() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Debug().DropTableIfExists(&models.Post{}, &models.User{}).Error
	if err != nil {
		log.Fatal(err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Post{}).Error
	if err != nil {
		log.Fatal(err)
	}

	err = db.Debug().Model(&models.Post{}).AddForeignKey("author_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatal(err)
	}

	for i, user := range users {
		err = db.Debug().Model(&models.User{}).Create(&user).Error
		if err != nil {
			log.Fatal(err)
		}
		posts[i].AuthorID = user.ID
		posts[i].Author = user
		err = db.Debug().Model(&models.Post{}).Create(&posts[i]).Error
		if err != nil {
			console.Pretty(&posts[i].Author)
			log.Fatal(err)
		}
	}
}