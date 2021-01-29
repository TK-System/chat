package gin

import (
	"chat/backend/interfaces/controller"
	"chat/backend/interfaces/gateway/database"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *Router,db database.UserDBHandler){
	ctrl := controller.NewUserController(db)
	r.GET("/user/:id",downloadUser(ctrl))
}

func downloadUser(ctrl 	*controller.UserController)func(c *gin.Context){
	return func(c *gin.Context){
		//rw := NewResponseWriter(c)


		userID := c.Param("id")
		id ,err := strconv.Atoi(userID)
		if err !=nil {
			// log error
		}
		err = ctrl.DownloadUser(id)
		if err !=nil {
			// log error
		}
		return
	}
}



func uploadUser(ctrl 	*controller.UserController)func(c *gin.Context){
	return func(c *gin.Context){


		keys := []string{
			"name",
			"mail",
			"sex",
		}
		params := getParams(c,keys)


		err := ctrl.UploadUser(params["name"])
		if err !=nil {
			// log error
		}
		return
	}
}

func getParams(c *gin.Context,keys []string)(map[string]string){
	values := make(map[string]string,len(keys))
	for _,k :=range keys{
		values[k]= c.Param(k)
	}
	return values
}
