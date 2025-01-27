package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"web_app/models"
)

// 把每一步数据库操作封装成函数
// 待logic层根据业务需求调用
const secret = "liwenzhou.com"

func Login(user *models.User) (err error) {
	oPassword := user.Password //用户登录的密码
	sqlStr := `select user_id,username,password from user where username=?`
	err = db.Get(user, sqlStr, user.Username)
	if err == sql.ErrNoRows {
		return ErrorUserNotExist
	}
	//查询数据库失败
	if err != nil {
		return err
	}

	//判断密码是否正确
	password := encryptPassword(oPassword)
	if password != user.Password {
		return ErrorInvalidPassword
	}
	return
}

// InsertUser 向数据库中插入一条新的用户记录
func InsertUser(user *models.User) (err error) {
	//对密码进行加密
	user.Password = encryptPassword(user.Password)
	//执行SQL语句入库
	sqlStr := `insert into user(user_id,username,password) values (?,?,?)`
	_, err = db.Exec(sqlStr, user.UserID, user.Username, user.Password)
	return err
}

// func QueryUserByUsername() {
//
// }
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(username string) (err error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	//	fmt.Println("daozhel")
	if err = db.Get(&count, sqlStr, username); err != nil {
		return err
	}
	//	fmt.Println("111")
	if count > 0 {
		return ErrorUserExist
	}
	//	fmt.Println("caol")
	return
}
