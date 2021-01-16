package server

import "chat/backend/entity/domain"


type FriendInputPort interface {
	DownloadFriend(FriendIdentifiers)error
	UploadFriend(FriendComponents)error
}

// UserComponents こいつの名前迷ってる
// elements, materials,
type FriendComponents interface{
	FriendName()string
}
type FriendIdentifiers interface{
	FriendID()domain.FriendID
}

// output port
// Presenter で定義され，Interactor で使用される

type FriendOutputPort interface{
	DownloadFriend([]domain.Friend)error
	UploadFriend()
}
