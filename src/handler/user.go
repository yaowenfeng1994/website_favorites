package handler

import (
	"log"
	"time"
	"model"
	"net/http"
	"github.com/gin-gonic/gin"
	"fmt"
)

//type User struct {
//	Account  string  `form:"account" json:"account" binding:"required"`
//	Nickname string  `form:"nickname" json:"nickname" binding:"required"`
//	Mail     string  `form:"mail" json:"mail" binding:"required"`
//	Time     int64   `form:"time" json:"time" binding:"required"`
//}

func CreateUserHandler(c *gin.Context) {
	var user  = map[string]gin.User{}
	var respData map[string]interface{}

	resp := BaseResponse{}
	user = make(map[string]gin.User)
	respData = make(map[string]interface{})

	if err := c.BindJSON(&user); err == nil{
		data := user["data"]
		account := data.Account
		nickname := data.Nickname
		mail := data.Mail
		data.Time = time.Now().Unix()
		t := data.Time
		UserId, err := model.InsertUser(account, nickname, mail, t)
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
	} else {
		log.Print(err)
		c.JSON(http.StatusBadRequest, resp)
	}
}

//func GetUserInfoHandler(c *gin.Context) {
//	account := c.Query("account")
//	session := sessions.Default(c)
//	fmt.Println(session.Options)
//	userAccessToken := session.Get(account)
//	fmt.Printf("[DoSomethine] user access token is %s\n", userAccessToken)
//	c.JSON(http.StatusOK, nil)
//}
//
func LoginHandler(c *gin.Context, w http.ResponseWriter) {

	account := c.PostForm("account")
	nickname := c.PostForm("nickname")
	fmt.Println(account)

	if account == "610733719"{
		var sessionMgr *gin.SessionMgr
		sessionMgr = gin.NewSessionMgr("TestCookieName", 20)

		var sessionID = sessionMgr.StartSession(w)
		var loginUserInfo = gin.User{Account: account, Nickname: nickname}
		var loginUserInfoPointer *gin.User
		loginUserInfoPointer = &loginUserInfo
		var onlineSessionIDList = sessionMgr.GetSessionIDList()

		for _, onlineSessionID := range onlineSessionIDList {
			if userInfo, ok := sessionMgr.GetSessionVal(onlineSessionID, account); ok {
				if value, ok := userInfo.(gin.User); ok {
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

