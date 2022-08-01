package domain

import (
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
