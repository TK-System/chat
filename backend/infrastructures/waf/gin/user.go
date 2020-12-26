package gin

import (
	"chat/backend/interfaces/controller"

	"github.com/gin-gonic/gin"
)



func UserHandler(ctrl 	*controller.UserController)func(c *gin.Context){
	return func(c *gin.Context){
		
	}
}
