package domain

type FriendID int
type Friend interface{
	ID() FriendID
	Like() bool

}

type friend struct{
	id FriendID
	like bool
}

func NewFriend(id FriendID,like bool)Friend{
	return &friend{
		id: id,
		like:like,
	}
}

func(f *friend)ID()FriendID{return f.id}
func(f *friend)Like()bool{return f.like}
