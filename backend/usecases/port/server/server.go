package server

import "chat/backend/entity/domain"

// input port
// Interactor で定義され， Controller で使用される

type UserInputPort interface {
	DownloadUser(UserIdentifiers)error
	UploadUser(UserComponents)error
}

// UserComponents こいつの名前迷ってる
// elements, materials,
type UserComponents interface{
	UserName()string
}
type UserIdentifiers interface{
	UserID()domain.UserID
}

// output port
// Presenter で定義され，Interactor で使用される

type UserOutputPort interface{
	DownloadUser([]domain.User)error
	UploadUser()
}
