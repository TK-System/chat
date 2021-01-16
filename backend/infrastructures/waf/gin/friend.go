package gin
import (
	"chat/backend/interfaces/controller"

	"github.com/gin-gonic/gin"
)


func FriendRoute(r *Router){
	ctrl := controller.NewFriendController()

	r.POST("/friendList/",FriendHandler(ctrl))
}

func FriendHandler (ctrl 	*controller.FriendController)func(c *gin.Context){
	return func(c *gin.Context){
		
	}
}