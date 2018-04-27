package main

import (
	"fmt"
	"time"
	"handler"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
)

func main(){
	fmt.Println("hello,this is my first golang project!")
	//handler.Insert()
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	store := sessions.NewCookieStore([]byte("secret"))
	store.Options(sessions.Options{
		MaxAge: int(10 * time.Second), //1min
		Path:   "/",
	})
	fmt.Println(store)
	r.Use(sessions.Sessions("mysession", store))

	r.POST("/user/create", handler.CreateUserHandler)
	r.POST("/user/login", handler.LoginHandler)
	r.GET("/user/get", handler.GetUserInfoHandler)
	r.Run(":9999")
}
