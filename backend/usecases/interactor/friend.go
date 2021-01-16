package interactor

import (
	// "chat/backend/entity/domain"
	"chat/backend/usecases/port/repository"
	"chat/backend/usecases/port/server"

	// "golang.org/x/xerrors"
)


type FriendInteractor struct{
	//Config 
	OutputPort server.FriendOutputPort
	Repository repository.FriendRepository
}

func NewFriendInteractor(
	//conf
	out server.FriendOutputPort,
	repo repository.FriendRepository,
) *FriendInteractor{
	return &FriendInteractor{
		//Config: conf,
		OutputPort: out,
		Repository: repo,
	}
}

func (i *FriendInteractor)DownloadFriend(server.FriendIdentifiers) error   {
return nil
}


func (i *FriendInteractor)UploadFriend(server.FriendComponents) error     {
	return nil
}