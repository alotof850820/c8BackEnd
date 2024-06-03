package serializer

import (
	"gin_mall_tmp/model"
	"gin_mall_tmp/pkg/util"
)

// 前端money 結構
type Money struct {
	UserId    uint   `json:"user_id" form:"user_id"`
	UserName  string `json:"user_name" form:"user_name"`
	UserMoney string `json:"user_money" form:"user_money"`
}

func BuildMoney(user *model.User, key string) *Money {
	// 設置key 以便解密
	util.Encrypt.SetKey(key)
	return &Money{
		UserId:    user.ID,
		UserName:  user.UserName,
		UserMoney: util.Encrypt.AesDecoding(user.Money), // 金錢解密
	}
}
