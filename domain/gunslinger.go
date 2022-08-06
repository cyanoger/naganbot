package domain

import (
	"context"
	"github.com/google/uuid"
	"time"
)

type Gunslinger struct {
	Id          uuid.UUID `gorm:"primary_key;size:36;<-:create"`
	Game        Game
	Player      User
	JoinedAt    time.Time
	ShotHimself bool
}

type GunslingerRepository interface {
	GetById(context.Context, uuid.UUID) (Gunslinger, error)
	Store(context.Context, *Gunslinger) error
	Update(context.Context, *Gunslinger) error
}

func NewGunslinger(game Game, player User) Gunslinger {
	g := Gunslinger{
		Id:          uuid.New(),
		Game:        game,
		Player:      player,
		JoinedAt:    time.Now(),
		ShotHimself: false,
	}

	return g
}
