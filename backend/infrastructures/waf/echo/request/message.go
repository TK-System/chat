package request

import (
	"net/http"

	"github.com/labstack/echo"
)

type Message struct {
}

func NewMessage()*Message{
	return &Message{}
}

func (m *Message)Get()echo.HandlerFunc{
	return func(c echo.Context)error{
		return c.JSON(http.StatusOK,&TestJSON{"HelloWorld"})
	}
}

type TestJSON struct{
	Message string `json:"message"`
}
