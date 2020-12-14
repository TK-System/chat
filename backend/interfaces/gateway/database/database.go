package database

import "chat/backend/entity/domain"

type DBHandler interface{
	User()
}

type UserRepositoryAdapter struct{
	//config 
	DBHandler
}

func NewUserRepositoryAdapter(dbHandler DBHandler)*UserRepositoryAdapter{
	return &UserRepositoryAdapter{
		DBHandler: dbHandler,
	}
}

func (adapter *UserRepositoryAdapter)	Find([]domain.UserID)([]domain.User,error){
	adapter
}
