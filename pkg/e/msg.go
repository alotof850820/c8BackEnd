package e

var MsgFlags = map[int]string{
	Success:       "ok",     //200
	Error:         "fail",   //500
	InvalidParams: "请求参数错误", //400

	ErrorExistUser:             "用戶已存在",     //30001
	ErrorFailEncryption:        "密碼加密失敗",    //30002
	ErrorExistUserNotFound:     "用戶不存在",     //30003
	ErrorNotCompare:            "密碼錯誤",      //30004
	ErrorAuthToken:             "token生成失敗", //30005
	ErrorAuthCheckTokenTimeout: "token過期",   //30006
	ErrorUploadFail:            "上傳失敗",      //30007

	ErrorSendEmail: "發送郵件失敗", //30008

	ErrorProductImgUpload: "上傳商品失敗", //40001

	ErrorFavoriteExist: "收藏已存在", //50001
}

func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if ok {
		return msg
	} else {
		return MsgFlags[Error]
	}
}
