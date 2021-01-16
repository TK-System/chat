package domain

type GroupID int
type GroupName string
type Group interface{
	ID() GroupID
	Name() GroupName
}

type group struct{
	id GroupID
	name GroupName
}

func NewGroup(id GroupID,name GroupName)Group{
	return &group{
		id: id,
		name: name,
	}
}

func(g *group)ID()GroupID{return g.id}
func(g *group)Name()GroupName{return g.name}
