package repository

import "chat/backend/entity/domain"

type UserRepository interface{
	Find([]domain.UserID)([]domain.User,error)
	Add([]domain.User)error
}

type GroupRepository interface{
	Find([]domain.GroupID)([]domain.Group,error)
	Add([]domain.Group)error
}

type MessageRepository interface{
	Find([]domain.MessageID)([]domain.Message,error)
	Add([]domain.Message)error
}

