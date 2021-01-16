package presenter

import (
	"chat/backend/entity/domain"
	"chat/backend/usecases/port/server"
)

type HttpPresenter struct{
	ResponseServer
}

func NewHttpPresenter(resServer ResponseServer)*HttpPresenter{
	return &HttpPresenter{
		ResponseServer: resServer,
	}
}

type ResponseServer interface{
	Response(content interface{})error
}

