package e

var MsgFlags = map[int]string{
	Success:       "ok",
	Error:         "fail",
	InvalidParams: "请求参数错误",

	ErrorExistUser:             "用户名已存在",
	ErrorFailEncryption:        "密码加密失败",
	ErrorUserNotFound:          "用户不存在",
	ErrorNotCompare:            "账号密码不匹配",
	ErrorAuthToken:             "token认证失败",
	ErrorAuthCheckTokenTimeout: "token已过期",
	ErrorUploadFail:            "图片上传失败",
	ErrorSendEmail:             "发送邮件失败",
	ErrorUserKey:               "密钥长度不足",

	ErrorProductImgUpload: "图片上传失败",

	ErrorFavoriteExist: "已收藏",

	ErrorAddressNotExist: "地址不存在",

	ErrorInsufficientFund: "用户金额不足",
	ErrorOutOfStock:       "商品库存不足",
}

// GetMsg 获取状态码对应的信息
func GetMsg(code int) string {
	msg, ok := MsgFlags[code]
	if !ok {
		return MsgFlags[Error]
	}
	return msg
}
