package controller

type UserController struct{
	InputPort server.UserInputPort
}

func NewUserController()*UserController{
	return &UserController{
		InputPort: interactor.NewUserInteractor(
			presenter.NewHttpPresenter(),
			nil,
		),
	}
}

func (c *UserController)DownloadUser(params int)error{
	return c.InputPort.DownloadUser(
		NewDownloadUserParam(params),
	)
}

func (c *UserController)UploadUser(params string)error{
	return c.InputPort.UploadUser(
		NewUploadUserParam(params),
	)
}

type uploadUserParam struct{
	params map[string]interface{}
}
func (p *uploadUserParam)UserName()string{
	return p.name.(string)
}


func NewUploadUserParam(params map[string]interface{})*uploadUserParam{
	key := "name"
	name,exist :=params[key]
	_,ok := name.(string)
	if !exist || !ok{
		return nil
	}

	return &uploadUserParam{
		params: params,
	}
}



type downloadUserParam struct{
	id domain.UserID
}

func NewDownloadUserParam(userID int)*downloadUserParam{
	return &downloadUserParam{
		id:domain.UserID(userID),
	}
}

func (p *downloadUserParam)UserID()domain.UserID{
	return p.id
}
