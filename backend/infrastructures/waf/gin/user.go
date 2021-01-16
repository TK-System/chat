package gin

import (
	"strconv"
	"chat/backend/interfaces/controller"

	"github.com/gin-gonic/gin"
)

func UserRoute(r *Router){
	ctrl := controller.NewUserController()
	r.GET("/user",downloadUser(ctrl))
}

func downloadUser(ctrl 	*controller.UserController)func(c *gin.Context){
	return func(c *gin.Context){
		key := "user_id"
		params := getParams(c,[]string{key})
		id ,err := strconv.Atoi(params[key])
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