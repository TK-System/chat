package cmd

import (
	"chat/backend/infrastructures/logger"
	"chat/backend/infrastructures/waf/gin"
	"fmt"

	"github.com/spf13/cobra"
)



var ChatCmd = &cobra.Command{
	Use: "chat",
	Short: "kyota's first app",
	//Args:cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		//cmd.Flags().GetString()
		
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

var ChatRun = &cobra.Command{
	Use: "run",
	Short: "kyota's first app",
	Args:cobra.MinimumNArgs(1),
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

//func init(){
//	ChatCmd.AddCommand(ChatRun)
//	cmdArgument()
//}


func cmdArgument(){
	ChatCmd.PersistentFlags().StringP("settings","s","","setting file path")
	ChatCmd.PersistentFlags().IntP("debug-level","d",0,"debug level 0: fatal, error, 1: info, 2: debug")
}
