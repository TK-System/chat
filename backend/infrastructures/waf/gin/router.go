package gin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type Router struct{
	*gin.Engine
	port string

}

func NewRouter(port string)*Router{
	return &Router{
		Engine: gin.Default(),
		port:port,
	}
}

func (r *Router)setRouter(){
	r.GET("/",MessageHandler)
	r.GET("/user")
}

func (r *Router)RunServer()error{
	r.setRouter()
	return r.Run(r.port)
}


func MessageHandler(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"message":"hello world",
	})
}
