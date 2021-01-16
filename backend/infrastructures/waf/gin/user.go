package gin

import (
	"chat/backend/interfaces/controller"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *Router){
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
		params := getParams(keys)
		

		err := ctrl.DownloadUser()
		if err !=nil {
			// log error
		}
		return
	}
}

func getParams(c *gin.Context,keys []string)(map[string]string){
	values := make(map[string]string,len(key))
	for _,k :=range keys{
		values[k]= c.Param(k)
	}
	return values
}