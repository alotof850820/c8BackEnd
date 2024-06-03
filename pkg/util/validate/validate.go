package validate

import (
	"log"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New() // 創建驗證器
	validate.SetTagName("v")   // 設置驗證器的標籤名
}

func outRes(tag string, err *error) {
	log.Println("---------------start" + tag + "---------------")
	log.Println(*err)
	log.Println("---------------end" + tag + "---------------")
	err = nil
}
