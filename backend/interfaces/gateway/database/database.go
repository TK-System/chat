package database

type UserDBHandler interface{
	GetUser(int)
	GetUsers([]int)()
	CreateUser()
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
	adapter.GetUsers(ids)
}

func (adapter *UserRepositoryAdapter)Add([]domain.User)error{
	err := adapter.CreateUser()
	return err
}