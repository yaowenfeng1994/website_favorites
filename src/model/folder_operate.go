package model

//func DefaultCreateFolder(account string, t int64) *SQLConnTransaction {
//	Tx, err := Pool.Begin()
//	UserId := AccountGetUserId(account)
//	var DefaultFolderList = []string{"美食", "工具", "游戏", "购物", "健康"}
//	if err != nil {
//		return
//	} else {
//		for _, FolderName := range DefaultFolderList {
//			_, err := Tx.Insert("INSERT INTO folder (`user_id`, `folder_name`, `create_time`)" +
//				" VALUES( ?, ?, ?)", UserId, FolderName, t)
//			if err != nil {
//				Tx.Rollback()
//			}
//		}
//		return Tx
//	}
//}

func QueryFolder(account string) ([]map[string]interface{}, error) {
	var DbData []map[string]interface{}
	rows, err := Pool.Query("SELECT folder.id as folder_id, folder_name From folder LEFT JOIN user " +
		"ON user.id = folder.user_id WHERE user.account = ?", account)
	if err != nil {
		return DbData, err
	} else {
		if len(rows) > 0 {
			return rows, err
		} else {
			return DbData, err
		}
	}
}
