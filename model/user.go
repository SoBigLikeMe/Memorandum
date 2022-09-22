package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserName       string `gorm:"unique" json:"userName,omitempty"`
	PasswordDigest string //密文
}
