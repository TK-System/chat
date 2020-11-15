package response

import "github.com/labstack/echo"

type Message struct{

}

func NewMessage()*Message{
	return &Message{}
}

func (m *Message)Get() func(echo.Context)error{
	return nil
}
