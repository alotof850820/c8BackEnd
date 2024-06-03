package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserName       string `gorm:"unique"`
	Email          string
	PasswordDigest string // 以加密的形式儲存密碼
	NickName       string
	Status         string
	Avatar         string
	Money          string
}

const (
	PasswordCost        = 12       // 密碼加密強度
	Active       string = "active" //激活用戶
)

// 加密密碼
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PasswordCost) // 生成密碼
	if err != nil {
		return err
	}
	user.PasswordDigest = string(bytes)
	return err
}

func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password)) // 驗證密碼
	return err == nil
}
