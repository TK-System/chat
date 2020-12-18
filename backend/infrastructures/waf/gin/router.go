package gin

import (
	"net/http"
	"io/ioutil"
	"fmt"

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
	r.GET("/login/",LoginHandler)
	r.POST("/friendList/",FriendListHandler)

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

func LoginHandler(c *gin.Context){
	res := fileRead("/home/ampamman/go/src/chat/frontend/login.html")
	fmt.Println("\n",c.PostForm("mail"),"\n",c.PostForm("password"))
	c.Writer.WriteString(
		res,
	)
}

func FriendListHandler(c *gin.Context){
	res := fileRead("/home/ampamman/go/src/chat/frontend/friendList.html")
	fmt.Println("\n",c.PostForm("mail"),"\n",c.PostForm("password"))
	fmt.Println(c)
	c.Writer.WriteString(
		res,
	)
}

func fileRead(path string)string{
	bytes, err := ioutil.ReadFile(path)
    if err != nil {
        panic(err)
    }

	return string(bytes)
}
