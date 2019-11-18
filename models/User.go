package models

import (
	"api/security"
	"errors"
	"github.com/badoux/checkmail"
	"html"
	"strings"
	"time"
)

type User struct {
	ID uint32 `gorm:"primary_key;auto_increment" json:"id"`
	Nickname string `gorm:"size:20;not null;unique" json:"nickname"`
	Email string `gorm:"size:50;not null;unique" json:"email"`
	Password string `gorm:"size:60;not null" json:"password"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`

	Posts []Post `gorm:"foreignkey:AuthorID" json:"posts,omitempty"`
}

func (u *User) BeforeSave() error {
	hashedPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.Nickname = html.EscapeString(strings.TrimSpace(u.Nickname))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Nickname == "" {
			return errors.New("nickname is required")
		}
		if u.Email == "" {
			return errors.New("email is required")
		}

		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return err
		}
		return nil
	default:
		if u.Nickname == "" {
			return errors.New("nickname is required")
		}
		if u.Email == "" {
			return errors.New("email is required")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid email format")
		}
		return nil
	}
}