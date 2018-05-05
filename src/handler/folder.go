package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"model"
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
	var respData map[string]interface{}
	resp := BaseResponse{}
	respData = make(map[string]interface{})
	if sessionID == "" {
		resp.InitBaseResponse(0x0004, respData)
		c.JSON(http.StatusBadRequest, resp)
	} else {
		//var a= []string{"a", "b", "c", "d", "ee"}
		account, ok := sessionMgr.GetSessionVal(sessionID)
		if ok {
			fmt.Println(account)
			switch account := account.(type) {
			case string:
				DbData, err := model.QueryFolder(account)
				if err != nil {
					resp.InitBaseResponse(0x0002, respData)
					c.JSON(http.StatusBadRequest, resp)
				} else {
					fmt.Println(DbData)
					c.HTML(http.StatusOK,
						"folder.html",
						gin.H{
							"list": DbData,
						})
				}
			}
		} else {
			resp.InitBaseResponse(0x0001, respData)
			c.JSON(http.StatusBadRequest, resp)
		}
	}
}
