package presenter

import (
	"chat/backend/entity/domain"
	"chat/backend/usecases/port/server"
)

type HttpPresenter struct{
	server.UserOutputPort
}

func NewHttpPresenter()*HttpPresenter{
	return &HttpPresenter{

	}
}

func (p *HttpPresenter)	DownloadUser([]domain.User)error{
	

	return nil
}
