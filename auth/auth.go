package auth

import (
	"api/database"
	"api/models"
	"api/security"
	"api/utils/channels"
	"errors"
	"github.com/jinzhu/gorm"
)

func SignIn(email, password string) (string, error){
	var db *gorm.DB
	var err error
	user := models.User{}
	done := make(chan bool)
	go func(ch chan <- bool) {
		db, err = database.Connect()
		if err != nil {
			ch <- false
			return
		}
		defer db.Close()

		err = db.Debug().Model(models.User{}).Where("email = ?", email).Take(&user).Error
		if err != nil {
			ch <- false
			return
		}

		err = security.VerifyPassword(user.Password, password)
		if err != nil {
			ch <- false
			return
		}
		ch <- true
		return
	}(done)

	if channels.OK(done) {
		return CreateJWTToken(user.ID)
	}
	if gorm.IsRecordNotFoundError(err) {
		return "", errors.New("user credentials is invalid")
	}
	return "", err
}