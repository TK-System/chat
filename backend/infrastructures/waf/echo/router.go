package echo

import (
	"chat/backend/infrastructures/waf/echo/request"

	"github.com/labstack/echo"
)

type Router struct{
	r *echo.Echo
}

func NewRouter()*Router{
	return &Router{
		r :echo.New(),
	}
}

func (r *Router)setRoute(){ 
	msg := request.NewMessage()
	r.r.GET("/",msg.Get())
}

func (r *Router)runServer()error{
	return r.r.Start(":12345")
}

func (r *Router)Run()error{
	r.setRoute()
	return r.runServer()
}
