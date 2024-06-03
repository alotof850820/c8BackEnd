package e

const (
	Success       = 200
	Error         = 500
	InvalidParams = 400

	// user 錯誤
	ErrorExistUser             = 30001 //用戶已存在
	ErrorFailEncryption        = 30002 //密碼加密失敗
	ErrorExistUserNotFound     = 30003 //用戶不存在
	ErrorNotCompare            = 30004 //密碼錯誤
	ErrorAuthToken             = 30005 //token生成失敗
	ErrorAuthCheckTokenTimeout = 30006 //token過期
	ErrorUploadFail            = 30007 //上傳失敗

	// notice 錯誤
	ErrorSendEmail = 30008 //發送郵件失敗

	// product 錯誤
	ErrorProductImgUpload = 40001 //上傳商品失敗

	// 收藏錯誤
	ErrorFavoriteExist = 50001 //收藏已存在
)
