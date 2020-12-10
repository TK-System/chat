package database

type DBHandler interface{

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
