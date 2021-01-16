package gin

func FriendRouter(r *Router){
	ctrl := controller.NewFriendListController()

	r.POST("/friendList/",FriendListHandler)
}

func FriendHundler (ctrl 	*controller.FriendListHundler)func(c *gin.Context){
	return func(c *gin.Context){
		
	}
}