package domain

import "context"

type User struct {
	Id        int64 `gorm:"primary_key;size:36;<-:create"`
	FirstName string
	LastName  string
	Username  string
}

type UserRepository interface {
	GetById(context.Context, int64) (User, error)
	Store(context.Context, *User) error
	Update(context.Context, *User) error
}

type UserUseCase interface {
	GetById(context.Context, int64) (User, error)
	Store(context.Context, *User) error
}

func NewUser(id int64, firstName string, lastName string, username string) User {
	u := User{
		Id:        id,
		FirstName: firstName,
		LastName:  lastName,
		Username:  username,
	}

	return u
}
