package domain

type User struct {
	Id        int64 `gorm:"primary_key;size:36;<-:create"`
	FirstName string
	LastName  string
	Username  string
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
