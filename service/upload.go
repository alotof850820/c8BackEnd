package service

import (
	"gin_mall_tmp/conf"
	"io"
	"mime/multipart"
	"os"
	"strconv"
)

// 上傳頭貼圖片
func UploadAvatarToLocalStatic(file multipart.File, userId uint, userName string) (filePath string, err error) {
	bId := strconv.Itoa(int(userId))                       //路徑拼接
	basePath := "." + conf.AvatarPath + "user" + bId + "/" //./static/imgs/avatar/bId/
	if !DirExisOrNot(basePath) {
		CreateDir(basePath)
	}
	avatarPath := basePath + userName + ".jpg" //把file後綴提取出來
	content, err := io.ReadAll(file)           // io.Reader 中读取所有可用的数据，并将其以 []byte 的形式返回。
	if err != nil {
		return "", err
	}
	err = os.WriteFile(avatarPath, content, 0666)
	if err != nil {
		return "", err
	}
	return "user" + bId + "/" + userName + ".jpg", nil
}

// 上傳商品封面圖片
func UploadProductToLocalStatic(file multipart.File, userId uint, productName string) (filePath string, err error) {
	bId := strconv.Itoa(int(userId))                        //路徑拼接
	basePath := "." + conf.ProductPath + "boss" + bId + "/" //./static/imgs/product/bId/
	if !DirExisOrNot(basePath) {
		CreateDir(basePath)
	}
	productPath := basePath + productName + ".jpg" //把file後綴提取出來
	content, err := io.ReadAll(file)               // io.Reader 中读取所有可用的数据，并将其以 []byte 的形式返回。
	if err != nil {
		return "", err
	}
	err = os.WriteFile(productPath, content, 0666)
	if err != nil {
		return "", err
	}
	return "boss" + bId + "/" + productName + ".jpg", nil
}

// 判斷文件路徑是否存在
func DirExisOrNot(fileAddr string) bool {
	s, err := os.Stat(fileAddr)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 創建文件夾
func CreateDir(dirName string) bool {
	//755 7户对目录拥有读、写、执行权限。 5组成员对目录拥有读和执行权限，但没有写权限。5其他用户对目录拥有读和执行权限，但没有写权限。
	err := os.MkdirAll(dirName, 0755)
	return err == nil
}
