package model

import (
	"libs"
	"github.com/gin-gonic/gin"
)

var (
	Pool *libs.SQLConnPool
)

func init() {
	Pool = libs.InitMySQLPool("127.0.0.1", "website_favorites", "root", "1q2w3e", "utf8", 200, 100)
}

//func InsertUser(account string, nickname string, mail string, password string, t int64) (int64, error) {
//	LastUserId, err := Pool.Insert("INSERT INTO user (`account`, `nickname`, `mail`, `create_time`)" +
//		" VALUES( ?, ?, ?, ?)", account, nickname, mail, t)
//	if LastUserId != 0 {
//		fmt.Println(password, LastUserId)
//		_, err1 := Pool.Insert("INSERT INTO password (`password`, `user_id`, `create_time`)" +
//			" VALUES( ?, ?, ?)", password, LastUserId, t)
//		return LastUserId, err1
//	} else {
//		return LastUserId, err
//	}
//}

func InsertUser(account string, nickname string, mail string, password string, t int64) (int64, error) {
	Tx, err := Pool.Begin()
	var DefaultFolderList = []string{"美食", "工具", "游戏", "购物", "健康"}
	if err != nil {
		return 0, err
	} else {
		LastUserId, err := Tx.Insert("INSERT INTO user (`account`, `nickname`, `mail`, `create_time`)" +
			" VALUES( ?, ?, ?, ?)", account, nickname, mail, t)
		if err != nil {
			Tx.Rollback()
			return 0, err
		} else {
			_, err1 := Tx.Insert("INSERT INTO password (`password`, `user_id`, `create_time`)" +
				" VALUES( ?, ?, ?)", password, LastUserId, t)
			if err1 != nil {
				Tx.Rollback()
				return 0, err1
			} else {
				for _, FolderName := range DefaultFolderList {
					_, err2 := Tx.Insert("INSERT INTO folder (`user_id`, `folder_name`, `create_time`)" +
						" VALUES( ?, ?, ?)", LastUserId, FolderName, t)
					if err2 != nil {
						Tx.Rollback()
						return 0, err2
					}
				}
				err3 := Tx.Commit()
				return LastUserId, err3
			}
		}
	}
}

func QueryUserPassword(account string) ([]gin.H, error) {
	var DbData []gin.H
	rows, err := Pool.Query("SELECT password From password LEFT JOIN user " +
		"ON user.id = password.user_id WHERE user.account = ?", account)
	if err != nil {
		return DbData, err
	} else {
		if rows != nil {
			for _, row := range rows {
				DbData = append(DbData, gin.H(row))
			}
		}
		return DbData, err
	}
}

func AccountGetUserId(account string) interface{} {
	var UserList  []gin.H
	rows, err := Pool.Query("SELECT user_id From user WHERE user.account = ?", account)
	if err != nil {
		return 0
	} else {
		if rows != nil {
			for _, row := range rows {
				UserList = append(UserList, gin.H(row))
			}
			UserId := UserList[0]["user_id"]
			return UserId
		} else {
			return 0
		}
	}
}
