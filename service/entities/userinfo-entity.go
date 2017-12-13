package entities

import (
	//"time"
)

// UserInfo .
type UserInfo struct {
	UID        int `orm:"id,auto-inc"` //语义标签
	Username   string
	Password   string
	Email	   string
	Tel		   string
}

// NewUserInfo .
func NewUserInfo(u UserInfo) *UserInfo {
	if len(u.Username) == 0 {
		panic("UserName shold not null!")
	}
	return &u
}
