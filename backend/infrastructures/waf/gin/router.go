package gin

import (
	"chat/backend/infrastructures/strage/mysql"
	"fmt"
	"io/ioutil"
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
db ,err:= mysql.NewDB("tksystem","covit19","localhost","3306","chatdb")
if err !=nil{
return 
}

	r.GET("/healthcheck",HealthCheck)
	r.GET("/",HealthCheck)
	
	UserRoute(r,db.NewUserTransfer())
	FriendRoute(r)

	r.GET("/login/",LoginHandler)
}

func (r *Router)RunServer()error{
	r.setRouter()
	return r.Run(r.port)
}


func HealthCheck(c *gin.Context){
	c.JSON(http.StatusOK,gin.H{
		"message":"hello world",
	})
}




func LoginHandler(c *gin.Context){
	res := fileRead("/home/ubuntu/app/page/login.html")
	fmt.Println("\n",c.PostForm("mail"),"\n",c.PostForm("password"))
	c.Writer.WriteString(
		res,
	)
}

func FriendListHandler(c *gin.Context){
	res := fileRead("/home/ubuntu/app/page/friendList.html")
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
