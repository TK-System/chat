package main

import (
	"chat/backend/infrastructures/waf/echo"
	"fmt"
)

func main() {
	fmt.Println("hello")
	
r := echo.NewRouter()
r.Run()

	return
}
