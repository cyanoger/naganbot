package domain

import (
	"context"
	"github.com/google/uuid"
)

type Chat struct {
	Id        int64 `gorm:"primary_key;size:36;<-:create"`
	Title     string
	FirstName string
	LastName  string
}

type ChatRepository interface {
	GetById(context.Context, uuid.UUID) (Chat, error)
	Store(context.Context, *Chat) error
	Update(context.Context, *Chat) error
}

func NewChat(id int64, title string, firstName string, lastName string) Chat {
	c := Chat{
		Id:        id,
		Title:     title,
		FirstName: firstName,
		LastName:  lastName,
	}

	return c
}
