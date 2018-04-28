package main

import (
	"fmt"
	"handler"
	"github.com/gin-gonic/gin"
)


func main(){
	fmt.Println("hello,this is my first golang project!")
	//handler.Insert()
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.POST("/user/create", handler.CreateUserHandler)
	r.GET("/user/login", handler.LoginHandler)
	r.GET("/user/logout", handler.LogoutHandler)
	r.GET("/user/get", handler.GetUserInfoHandler)
	r.Run(":9999")
}
