package interactor

import (
	"chat/backend/entity/domain"
	"chat/backend/usecases/port/repository"
	"chat/backend/usecases/port/server"

	"golang.org/x/xerrors"
)


type UserInteractor struct{
	//Config 
	OutputPort server.UserOutputPort
	Repository repository.UserRepository
}

func NewUserInteractor(
	//conf
	out server.UserOutputPort,
	repo repository.UserRepository,
) *UserInteractor{
	return &UserInteractor{
		//Config: conf,
		OutputPort: out,
		Repository: repo,
	}
}

func(i *UserInteractor)DownloadUser(userIdentifir server.UserIdentifiers)error{
	return i.DownloadUsers(userIdentifir)
}

func(i *UserInteractor)DownloadUsers(identifirs ...server.UserIdentifiers)error{
	userIDs := make([]domain.UserID,0,len(identifirs))
	for _,identifir := range identifirs{
		userIDs = append(userIDs, identifir.UserID())
	}

	users ,err :=i.Repository.Find(userIDs)
	if err !=nil {
		return xerrors.Errorf("interactor.go DownloadUsers() : %w",err)
	}

	return i.OutputPort.DownloadUser(users)
}



func(i *UserInteractor)UploadUser(server.UserComponents)error{

	return nil
}
