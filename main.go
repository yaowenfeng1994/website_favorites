package main

import (
	"fmt"
	"handler"
	"github.com/gin-gonic/gin"
	"page"
)


func main(){
	fmt.Println("hello,this is my first golang project!")
	//handler.Insert()
	gin.SetMode(gin.DebugMode)

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/user/create", page.RegisterPage)
	r.GET("/user/create_success", page.RegisterSuccessPage)
	r.POST("/user/create", handler.CreateUserApi)

	r.GET("/user/login", page.LoginPage)
	r.POST("/user/login", handler.LoginApi)

	r.GET("/folder", handler.GetFolderListApi)
	r.GET("/user/logout", handler.LogoutApi)
	//r.GET("/user/get", handler.GetUserInfoApi)
	r.Run(":9999")
}
