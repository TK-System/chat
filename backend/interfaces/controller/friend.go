package controller

type FriendController struct{
	InputPort server.FriendInputPort
}

func NewFriendController()*FriendController{
	return &FriendController{
		InputPort: interactor.NewFriendInteractor(
			presenter.NewHttpPresenter(),
			nil,
		),
	}
}