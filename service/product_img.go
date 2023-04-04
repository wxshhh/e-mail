package service

import (
	"context"
	"gin_mall/dao"
	"gin_mall/pkg/e"
	"gin_mall/serializer"
	"strconv"
)

type ProductImgService struct {
}

func (service *ProductImgService) List(ctx context.Context, id string) serializer.Response {
	code := e.Success
	var err error
	productImgDao := dao.NewProductImgDao(ctx)
	pid, err := strconv.Atoi(id)
	if err != nil {
		code = e.InvalidParams
		return serializer.ErrorByCode(code, err)
	}
	productImgs, err := productImgDao.ListProductImg(uint(pid))
	return serializer.BuildListResponse(serializer.BuildProductImgs(productImgs), uint(len(productImgs)))
}
