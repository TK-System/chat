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


func (p *HttpPresenter)	DownloadUser(users []domain.User)error{
	type GetUserResponse struct{
		UserID int `json:"userId"`
		UserName string `json:"userName"`
	}
	
	body := make([]GetUserResponse,len(users))
	for i,user :=range users{
		body[i]=GetUserResponse{
			UserID: int(user.ID()),
			UserName: user.Name(),
		}
	}

	return nil
}
