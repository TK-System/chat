package domain

type UserID int

type User interface{
	ID()UserID
	Name()string
}

func NewUser(id UserID ,name string)User{
	return &user{
		id: id,
		name: name,
	}
}

type user struct{
	id UserID
	name string
}

func(u *user)ID()UserID{return u.id}
func(u *user)Name()string{return u.name}
