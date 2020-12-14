package main

import (
	"chat/backend/infrastructures/waf/gin"
	"fmt"
)

func main() {
	fmt.Println("hello")

	r := gin.NewRouter(":1234")
	err :=r.RunServer()
	if err !=nil{
		fmt.Printf("%#v\n",err)
	}
	return
}
