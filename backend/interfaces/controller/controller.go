package controller

import (
	"chat/backend/entity/domain"
	"chat/backend/interfaces/presenter"
	"chat/backend/usecases/interactor"
	"chat/backend/usecases/port/server"
)

type UserController struct{
	InputPort server.UserInputPort
}

func NewUserController()*UserController{
	return &UserController{
		InputPort: interactor.NewUserInteractor(
			presenter.NewHttpPresenter(),
			nil,
		),
	}
}

func (c *UserController)DownloadUser(params int)error{
	return c.InputPort.DownloadUser(
		NewDownloadUserParam(params),
	)
}

func (c *UserController)UploadUser(params string)error{
	return c.InputPort.UploadUser(
		NewUploadUserParam(params),
	)
}

type uploadUserParam struct{
	name string
}

func NewUploadUserParam(name string)*uploadUserParam{
	return &uploadUserParam{
		name: name,
	}
}

func (p *uploadUserParam)UserName()string{
	return p.name
}


type downloadUserParam struct{
	id domain.UserID
}

func NewDownloadUserParam(userID int)*downloadUserParam{
	return &downloadUserParam{
		id:domain.UserID(userID),
	}
}

func (p *downloadUserParam)UserID()domain.UserID{
	return p.id
}
