package handler

import (

	"time"

	"net/http"
	"github.com/gin-gonic/gin"
	"libs"
	"model"
	"log"
)

//type User struct {
//	Account  string  `form:"account" json:"account" binding:"required"`
//	Nickname string  `form:"nickname" json:"nickname" binding:"required"`
//	Mail     string  `form:"mail" json:"mail" binding:"required"`
//	Time     int64   `form:"time" json:"time" binding:"required"`
//}

var sessionMgr *libs.SessionMgr = nil

func init() {
	sessionMgr = libs.NewSessionMgr("TestCookieName", 60)
}

func CreateUserApi(c *gin.Context) {
	//var user  = map[string]libs.User{}
	account := c.PostForm("account")
	nickname := c.PostForm("nickname")
	mail := c.PostForm("mail")
	password := c.PostForm("password")
	t := time.Now().Unix()
	var respData map[string]interface{}

	resp := BaseResponse{}
	//user = make(map[string]libs.User)
	respData = make(map[string]interface{})

	UserId, err := model.InsertUser(account, nickname, mail, password, t)
	respData["user_id"] = UserId
	if err != nil {
		resp.InitBaseResponse(0x0003, respData)
		log.Println(err.Error())
		c.JSON(http.StatusBadRequest, resp)
	} else {
		resp.InitBaseResponse(0x0000, respData)
		log.Printf("create account success(user_id: %d)!", UserId)
		c.JSON(http.StatusOK, resp)
	}
}

func GetUserInfoApi(c *gin.Context) {
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

func LoginApi(c *gin.Context) {

	account := c.PostForm("account")
	password := c.PostForm("password")

	if account == "610733719"{
		var sessionID = sessionMgr.StartSession(c.Writer, c.Request)
		var loginUserInfo = libs.User{Account: account, Password: password}
		var loginUserInfoPointer *libs.User
		loginUserInfoPointer = &loginUserInfo
		var onlineSessionIDList = sessionMgr.GetSessionIDList()

		for _, onlineSessionID := range onlineSessionIDList {
			if userInfo, ok := sessionMgr.GetSessionVal(onlineSessionID, account); ok {
				if value, ok := userInfo.(libs.User); ok {
					if value.Account == account {
						sessionMgr.EndSessionBy(onlineSessionID)
						}
					}
				}
			}
		sessionMgr.SetSessionVal(sessionID, account, loginUserInfoPointer)

		c.JSON(http.StatusOK, gin.H{
			"login_success": 1,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"login_success": 0,
		})
	}
}

func LogoutApi(c *gin.Context) {
	sessionMgr.EndSession(c.Writer, c.Request) //用户退出时删除对应session
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
	return
}
