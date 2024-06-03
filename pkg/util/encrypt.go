package util

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"errors"
)

var Encrypt *Encryption

// AES 對稱加密
type Encryption struct {
	Key string
}

func init() {
	Encrypt = NewEncryption()
}

func NewEncryption() *Encryption {
	return &Encryption{}
}

// 对 srcByte 进行填充，使其长度满足 blockSize 的倍数。
func PadPwd(srcByte []byte, blockSize int) []byte {
	padNum := blockSize - len(srcByte)%blockSize      //计算需要填充的字节数。
	ret := bytes.Repeat([]byte{byte(padNum)}, padNum) // 得到了填充的字节序列。
	srcByte = append(srcByte, ret...)                 //将字节逐个追加到 srcByte 中
	return srcByte
}

// AesEncoding 加密
func (k *Encryption) AesEncoding(src string) string {
	srcByte := []byte(src)                     //转换为字节数组
	block, err := aes.NewCipher([]byte(k.Key)) // 创建AES 加密块，使用 k.Key 作为密钥。
	if err != nil {
		return src
	} else {
		// 密碼填充
		NewSrcByte := PadPwd(srcByte, block.BlockSize())
		dst := make([]byte, len(NewSrcByte))
		block.Encrypt(dst, NewSrcByte) //對稱加密，并将加密结果存储到 dst 中。
		// base64編碼
		pwd := base64.StdEncoding.EncodeToString(dst)
		return pwd

	}
}

// UnPadPed 去掉填充的部分
func UnPadPwd(dst []byte) ([]byte, error) {
	if len(dst) <= 0 {
		return dst, errors.New("長度有誤")
	} else {
		// 去掉的長度
		unpadNum := int(dst[len(dst)-1]) // 计算需要去除的填充长度。
		strErr := "error"
		op := []byte(strErr)
		if len(dst) < unpadNum {
			return op, nil
		} else {
			str := dst[:(len(dst) - unpadNum)] //进行解填充
			return str, nil
		}
	}
}

// AesDecoding 解密
func (k *Encryption) AesDecoding(pwd string) string {
	pwdByte, err := base64.StdEncoding.DecodeString(pwd)
	if err != nil {
		return pwd // 返回原始密码字符串，因为无法解码
	}
	block, err := aes.NewCipher([]byte(k.Key)) // 创建AES 解密块，使用 k.Key 作为密钥。
	if err != nil {
		return pwd // 返回原始密码字符串，因为无法创建解密块
	}
	dst := make([]byte, len(pwdByte))
	block.Decrypt(dst, pwdByte) // 解密
	dst, err = UnPadPwd(dst)    // 去掉填充
	if err != nil {
		return "0" // 返回错误码 "0"，表示解密失败
	}
	return string(dst) // 返回解密后的明文字符串
}

func (k *Encryption) SetKey(key string) {
	k.Key = key
}
