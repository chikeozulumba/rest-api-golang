package models

import (
	"errors"
	"html"
	"strings"
	"time"
)

type Post struct {
	ID uint64 `gorm:"primary_key;auto_increment" json:"id"`
	Title string `gorm:"size:30;not null;unique" json:"title"`
	Body string `gorm:"size:10000;not null;unique" json:"body"`
	Author User `gorm:"-" json:"author"`
	AuthorID uint32 `gorm:"not null" json:"author_id"`
	CreatedAt time.Time `gorm:"type:timestamp" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp" json:"updated_at"`
}

func (p *Post) Prepare() {
	p.Title = html.EscapeString(strings.TrimSpace(p.Title))
	p.Body = html.EscapeString(strings.TrimSpace(p.Body))
	p.Author = User{}
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Post) Validate() error {
	if p.Title == "" {
		return errors.New("title is required")
	}
	if p.Body == "" {
		return errors.New("post body is required")
	}
	if p.AuthorID < 1 {
		return errors.New("no author associated")
	}
	return nil
}