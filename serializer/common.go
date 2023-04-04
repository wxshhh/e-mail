package serializer

import (
	"gin_mall/pkg/e"
	"gin_mall/pkg/utils"
)

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Error  string      `json:"error"`
}

type TokenData struct {
	User  interface{} `json:"user"`
	Token string      `json:"token"`
}

type DataList struct {
	Item  interface{} `json:"item"`
	Total uint        `json:"total"`
}

func BuildListResponse(item interface{}, total uint) Response {
	return Response{
		Status: e.Success,
		Data:   DataList{item, total},
		Msg:    "ok",
	}
}

func ErrorByCode(code int, err error) Response {
	utils.LogrusObj.Infoln("err:", err)
	return Response{
		Status: code,
		Msg:    e.GetMsg(code),
		Error:  err.Error(),
	}
}

func Success(data interface{}) Response {
	return Response{
		Status: e.Success,
		Msg:    e.GetMsg(e.Success),
		Data:   data,
	}
}
