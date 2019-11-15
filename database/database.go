package database

import (
	"api/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Connect() (*gorm.DB , error){
	dbCredentials := config.LoadDB()
	db, err := gorm.Open(dbCredentials[0], dbCredentials[1])
	if err != nil {
		return nil, err
	}
	return db, nil
}

