package controller

import (
// 	"chat/backend/entity/domain"
	"chat/backend/interfaces/presenter"
	"chat/backend/usecases/interactor"
	"chat/backend/usecases/port/server"
)

type FriendController struct{
	InputPort server.FriendInputPort
}

func NewFriendController()*FriendController{
	return &FriendController{
		InputPort: interactor.NewFriendInteractor(
			presenter.NewHttpPresenter(nil),
			nil,
		),
	}
}