package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func CreateFolderApi(c *gin.Context) {
	var sessionID = sessionMgr.CheckCookieValid(c.Writer, c.Request)
	if sessionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"login_success": 0,
		})
	}else {
		c.JSON(http.StatusOK, gin.H{
			"login_success": 1,
		})
	}
}

func GetFolderListApi(c *gin.Context) {
	var sessionID = sessionMgr.CheckCookieValid(c.Writer, c.Request)
	if sessionID == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"login_success": 0,
			"text": "请先登录",
		})
	} else {
		var a= []string{"a", "b", "c", "d", "e"}
		UserInfo, ok := sessionMgr.GetSessionVal(sessionID)
		if ok {
			fmt.Println(UserInfo)
		}

		c.HTML(http.StatusOK,
			"folder.html",
			gin.H{
				"login_success": 1,
				"list": a,
			})
	}
}
