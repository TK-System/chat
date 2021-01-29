package database

import (
	"chat/backend/entity/domain"
)

type UserDBHandler interface{
	GetUser(int)(User,error)
	GetUsers([]int)([]User,error)
	CreateUser([]int)error
}

type User interface{
	ID()int
	Name()string
	Sex()int
}


type UserRepositoryAdapter struct{
	//config  
	UserDBHandler
}

func NewUserRepositoryAdapter(dbHandler UserDBHandler)*UserRepositoryAdapter{
	return &UserRepositoryAdapter{
		UserDBHandler: dbHandler,
	}
}

func (adapter *UserRepositoryAdapter)Find(userIDs []domain.UserID)([]domain.User,error){
	ids := make([]int,len(userIDs))
	for i,v := range userIDs{
		ids[i]=int(v)
	}
	users,err := adapter.GetUsers(ids)
	if err !=nil{
		return nil,err
	}

	retUsers := make([]domain.User,len(users))
	for i,v := range users{
		retUsers[i]=domain.NewUser(domain.UserID(v.ID()),v.Name())
	}

return retUsers,nil

}

func (adapter *UserRepositoryAdapter)Add([]domain.User)error{
	err := adapter.CreateUser([]int{1,1,1})
	return err
}
