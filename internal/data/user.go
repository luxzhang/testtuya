package data

import (
	"database/sql"
	"fmt"
	"log"
)

// User .
type User struct {
	id int16
	username string
}

type UserApi interface {
	GetUser(uid int) (user *User, err error)
}

//查询用户
func (u User) GetUser(uid int) (user *User, err error) {
	db := GlobalConfigSetting.db
	if err != nil && err != sql.ErrNoRows {
		log.Printf("d.QueryRow error(%v)", err)
		return
	}
	querySql := fmt.Sprintf("SELECT u_id, u_username FROM `see_user` WHERE u_id=?;")
	user = new(User)
	err = db.QueryRow(querySql, uid).Scan(&user.id, &user.username)
	if err != nil && err != sql.ErrNoRows {
		log.Printf("QueryRow error(%v)", err)
		return
	}
	return user, nil
}