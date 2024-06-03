package validate

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

type User struct {
	Name        string         `v:"required,alphaunicode"`
	Age         uint8          `v:"gte=10,lte=30"`
	Phone       string         `v:"required,e164"` //+886
	Email       string         `v:"required,email"`
	likeColor1  string         `v:"iscolor"`
	likeColor2  string         `v:"hexcolor|rgb|rgba|hsl|hsla"`
	Address     *Address       `v:"required"`
	ContactUser []*ContactUser `v:"required,gte=1,dive"`                             // dive
	Hobby       []string       `v:"required,gte=2,dive,required,gte=2,alphaunicode"` //dive 深度验证
}

type Address struct {
	Province string `v:"required"`
	City     string `v:"required"`
}

type ContactUser struct {
	Name    string   `v:"required,alphaunicode"`
	Age     uint8    `v:"gte=20,lte=130"`
	Phone   string   `v:"required_without_all=Email Address,omitempty,e164"`  // 如果 Email 和 Address 字段都不为空，则该字段可以为空，否则它必须包含非空值。
	Email   string   `v:"required_without_all=Phone Address,omitempty,email"` //omitempty 字段的值为零值，则将其视为空
	Address *Address `v:"required_without_all=Phone Email"`
}

func StructValidate() {
	v := validate

	address := &Address{
		Province: "台灣",
		// City:     "台北",
	}
	contactUser := &ContactUser{
		Name:    "nick",
		Age:     230,
		Phone:   "+8865645341",
		Email:   "asd@.com",
		Address: address,
	}
	contactUser1 := &ContactUser{
		Name:    "nick",
		Age:     230,
		Phone:   "+8865645341",
		Email:   "asd@.com",
		Address: address,
	}
	user := &User{
		Name:        "ken",
		Age:         218,
		Phone:       "+8865645341",
		Email:       "asd@.com",
		likeColor1:  "#ffff",
		likeColor2:  "rgb(255,255,255)",
		Address:     address,
		ContactUser: []*ContactUser{contactUser, contactUser1},
		Hobby:       []string{"棒球", "足球"},
	}
	err := v.Struct(user)
	if err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok { // errs 是一个 ValidationErrors 类型
			for _, errasd := range errs {
				fmt.Println(errasd)
			}
		}
	}

}
