package domain

import (
	"context"
	"database/sql"
	"github.com/google/uuid"
	"time"
)

type Game struct {
	Id          uuid.UUID `gorm:"primary_key;size:36;<-:create"`
	Chat        Chat
	Owner       User
	Gunslingers []Gunslinger
	CreatedAt   time.Time
	PlayedAt    sql.NullTime
}

type GameRepository interface {
	GetById(context.Context, uuid.UUID) (Game, error)
	Store(context.Context, *Game) error
	Update(context.Context, *Game) error
}

func NewGame(chat Chat, owner User) Game {
	g := Game{
		Id:          uuid.New(),
		Chat:        chat,
		Owner:       owner,
		Gunslingers: []Gunslinger{},
		CreatedAt:   time.Now(),
		PlayedAt:    sql.NullTime{Time: time.Now()},
	}

	return g
}
