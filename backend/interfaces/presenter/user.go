package presenter

import (
	"chat/backend/entity/domain"
	// "chat/backend/usecases/port/server"
)

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

	err := p.Response(body)
	return err
}

func (p *HttpPresenter)UploadUser(){
	return 
}