package cmd

import (
	"chat/backend/infrastructures/logger"
	"chat/backend/infrastructures/waf/gin"
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "chat",
	Short: "kyota's first app",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello")
		l := logger.NewLogger(true,true,"")
		defer l.Close()
	
		r := gin.NewRouter(":1234")
		err :=r.RunServer()
		if err !=nil{
			fmt.Printf("%#v\n",err)
		}
	},
}
