package presenter

import (
	"chat/backend/entity/domain"
	// "chat/backend/usecases/port/server"
)

func (p *HttpPresenter)	DownloadFriend(Friends []domain.Friend)error{
	type GetFriendResponse struct{
		FriendID int `json:"FriendId"`
		// FriendName string `json:"FriendName"`
	}
	
	body := make([]GetFriendResponse,len(Friends))
	for i,Friend :=range Friends{
		body[i]=GetFriendResponse{
			FriendID: int(Friend.ID()),
			// FriendName: Friend.Name(),
		}
	}

	err := p.Response(body)
	return err
}

func (p *HttpPresenter)UploadFriend(){
	return 
}