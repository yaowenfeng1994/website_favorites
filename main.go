package main

import (
	"fmt"
	"handler"
	"gopkg.in/gin-gonic/gin.v1"
)

//func main(){
//
//	router := gin.Default()
//
//	router.GET("/get", func(c *gin.Context) {
//		firstname := c.DefaultQuery("firstname", "Guest")
//		lastname := c.Query("lastname")
//		c.String(http.StatusOK,"Hello %s %s", firstname, lastname)
//	})
//	router.Run(":8000")
//}

func main(){
	fmt.Println("hello,this is my first golang project!")
	//handler.Insert()
	gin.SetMode(gin.DebugMode)
	r := gin.Default()

	r.POST("/user", handler.UserHandler)
	r.Run(":9999")

}
