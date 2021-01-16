package domain


type RoomID int

type Room interface{
	ID()RoomID
	Name()string
}

func NewRoom(id RoomID ,name string)Room{
	return &room{
		id: id,
		name: name,
	}
}

type room struct{
	id RoomID
	name string
}

func(r *room)ID()RoomID{return r.id}
func(r *room)Name()string{return r.name}
