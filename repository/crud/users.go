package crud

import (
	"api/models"
	"api/utils/channels"
	"errors"
	"github.com/jinzhu/gorm"
	"time"
)

type repositoryUsersCRUD struct {
	db *gorm.DB
}

func NewRepositoryUsersCRUD(db *gorm.DB) *repositoryUsersCRUD {
	return &repositoryUsersCRUD{db}
}

func (r *repositoryUsersCRUD) Save(user models.User) (models.User, error) {
	var err error
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.User{}).Create(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
		return
	}(done)
	if channels.OK(done) {
		return user, nil
	}
	return models.User{}, err
}

func (r *repositoryUsersCRUD) FindAll() ([]models.User, error) {
	var err error
	var users []models.User
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.User{}).Limit(1000).Find(&users).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
		return
	}(done)
	if channels.OK(done) {
		return users, nil
	}
	return nil, err
}

func (r *repositoryUsersCRUD) FindById(uid uint32) (models.User, error) {
	var err error
	var user models.User
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		err = r.db.Debug().Model(&models.User{}).Where("id = ?", uid).Find(&user).Error
		if err != nil {
			ch <- false
			return
		}
		ch <- true
		return
	}(done)
	if channels.OK(done) {
		return user, nil
	}
	if gorm.IsRecordNotFoundError(err) {
		return models.User{}, errors.New("user record is not found")
	}
	return user, err
}

func (r *repositoryUsersCRUD) Update(uid uint32, user models.User) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.User{}).Where("id = ?", uid).Take(&models.User{}).UpdateColumns(
			map[string]interface{}{
				"nickname": user.Nickname,
				"email": user.Email,
				"updated_at": time.Now(),
			},
			)
		ch <- true
		return
	}(done)
	if channels.OK(done) {
		if rs.Error != nil {
			return 0, rs.Error
		}
		return rs.RowsAffected, nil
	}
	return 0, rs.Error
}

func (r *repositoryUsersCRUD) Delete(uid uint32) (int64, error) {
	var rs *gorm.DB
	done := make(chan bool)
	go func(ch chan<- bool) {
		defer close(ch)
		rs = r.db.Debug().Model(&models.User{}).Where("id = ?", uid).Take(&models.User{}).Delete(&models.User{})
		ch <- true
		return
	}(done)
	if channels.OK(done) {
		if rs.Error != nil {
			return 0, rs.Error
		}
		return rs.RowsAffected, nil
	}
	return 0, rs.Error
}