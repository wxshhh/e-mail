package v1

import (
	"encoding/json"
	"gin_mall/pkg/e"
	"gin_mall/serializer"
)

func ErrorResponse(err error) serializer.Response {
	if _, ok := err.(*json.UnmarshalTypeError); ok {
		return serializer.Response{
			Status: 400,
			Msg:    "JSON类型不匹配",
			Error:  err.Error(),
		}
	}
	return serializer.Response{
		Status: e.InvalidParams,
		Msg:    e.GetMsg(e.InvalidParams),
		Error:  err.Error(),
	}
}
