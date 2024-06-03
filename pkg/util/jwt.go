package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("yijiansanlian") // 隨機字串作為密鑰(私鑰)

// 宣告
type Claims struct {
	ID        uint   `json:"id"`
	UserName  string `json:"userName"`
	Authority int    `json:"authority"`
	jwt.RegisteredClaims
}

// 給予token
func GenerateToken(id uint, userName string, authority int) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour) // 有效期為24小時
	claims := Claims{
		id,
		userName,
		authority,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime), // 有效期
			Issuer:    "FanOne-Mall",                  // 签发人
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 生成token
	token, err := tokenClaims.SignedString(jwtSecret)                // 签名
	return token, err
}

// 驗證token
func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

type EmailClaims struct {
	UserID        uint   `json:"user_id"`
	Email         string `json:"email"`
	Password      string `json:"password"`
	OperationType uint   `json:"operation_type"`
	jwt.RegisteredClaims
}

// 給予信箱token
func GenerateEmailToken(userId, operationType uint, Email, Password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(24 * time.Hour) // 有效期為24小時
	claims := EmailClaims{
		userId,
		Email,
		Password,
		operationType,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime), // 有效期
			Issuer:    "FanOne-Mall",                  // 签发人
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims) // 生成token
	token, err := tokenClaims.SignedString(jwtSecret)                // 签名
	return token, err
}

// 驗證信箱token
func ParseEmailToken(token string) (*EmailClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &EmailClaims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*EmailClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
