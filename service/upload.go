package service

import (
	"gin_mall/conf"
	"io/ioutil"
	"mime/multipart"
	"os"
	"strconv"
)

// UploadAvatarToLocal 上传头像
func UploadAvatarToLocal(file multipart.File, uid uint, username string) (filePath string, err error) {
	// 强制类型转换，方便后续路径拼接
	bid := strconv.Itoa(int(uid))
	basePath := "." + conf.AvatarPath + "user" + bid + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	avatarPath := basePath + username + ".jpg"
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(avatarPath, content, os.ModePerm)
	if err != nil {
		return "", err
	}
	return "user" + bid + "/" + username + ".jpg", nil
}

// UploadProductToLocal 上传头像
func UploadProductToLocal(file multipart.File, uid uint, productName string) (filePath string, err error) {
	// 强制类型转换，方便后续路径拼接
	bid := strconv.Itoa(int(uid))
	basePath := "." + conf.ProductPath + "boss" + bid + "/"
	if !DirExistOrNot(basePath) {
		CreateDir(basePath)
	}
	productPath := basePath + productName + ".jpg"
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	err = ioutil.WriteFile(productPath, content, os.ModePerm)
	if err != nil {
		return "", err
	}
	return "boss" + bid + "/" + productName + ".jpg", nil
}

// DirExistOrNot 判断目录是否存在
func DirExistOrNot(filePath string) bool {
	s, err := os.Stat(filePath)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// CreateDir 创建目录
func CreateDir(dirPath string) bool {
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		return false
	}
	return true
}
