package page

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterPage(c *gin.Context) {
	c.HTML(http.StatusOK, "register.html", gin.H{"user": "User"})
	return
}


func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", gin.H{"user": "User"})
	return
}
