package service

import (
	"context"
	"gin_mall/dao"
	"gin_mall/model"
	"gin_mall/pkg/e"
	"gin_mall/serializer"
	"mime/multipart"
	"strconv"
	"sync"
)

type ProductService struct {
	ID            uint   `json:"id" form:"id"`
	Name          string `json:"name" form:"name"`
	CategoryID    uint   `json:"category_id" form:"category_id"`
	Title         string `json:"title" form:"title"`
	Info          string `json:"info" form:"info"`
	ImgPath       string `json:"img_path" form:"img_path"`
	Price         string `json:"price" form:"price"`
	DiscountPrice string `json:"discount_price" form:"discount_price"`
	OnSale        bool   `json:"on_sale" form:"on_sale"`
	Num           int    `json:"num" form:"num"`
	model.BasePage
}

func (service *ProductService) Create(ctx context.Context, uid uint, files []*multipart.FileHeader) serializer.Response {
	var boss *model.User
	var err error
	code := e.Success
	userDao := dao.NewUserDao(ctx)
	boss, _ = userDao.GetUserByID(uid)
	// 以第一张作为封面图
	tmp, _ := files[0].Open()
	// 将封面写入到本地，返回本地路径
	path, err := UploadProductToLocal(tmp, uid, service.Name)
	if err != nil {
		code = e.ErrorProductImgUpload
		return serializer.ErrorByCode(code, err)
	}
	// 构建商品信息
	product := &model.Product{
		Name:          service.Name,
		CategoryID:    service.CategoryID,
		Title:         service.Title,
		Info:          service.Info,
		ImgPath:       path,
		Price:         service.Price,
		DiscountPrice: service.DiscountPrice,
		OnSale:        true,
		Num:           service.Num,
		BossID:        boss.ID,
		BossName:      boss.Username,
		BossAvatar:    boss.Avatar,
	}
	// 写入数据库
	productDao := dao.NewProductDao(ctx)
	err = productDao.CreateProduct(product)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}
	// 并发创建图片 TODO
	wg := new(sync.WaitGroup)
	wg.Add(len(files))
	for i, file := range files {
		num := strconv.Itoa(i)
		productImgDao := dao.NewProductImgDaoByDB(productDao.DB)
		tmp, _ = file.Open()
		path, err = UploadProductToLocal(tmp, uid, service.Name+"_"+num)
		if err != nil {
			code = e.ErrorProductImgUpload
			return serializer.ErrorByCode(code, err)
		}
		productImg := &model.ProductImg{
			ProductID: product.ID,
			ImgPath:   path,
		}
		err = productImgDao.CreateProductImg(productImg)
		if err != nil {
			code = e.Error
			return serializer.ErrorByCode(code, err)
		}
		wg.Done()
	}
	wg.Wait()
	return serializer.Success(serializer.BuildProduct(product))
}

func (service *ProductService) List(ctx context.Context) serializer.Response {
	var products []*model.Product
	var err error
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	condition := make(map[string]interface{})
	if service.CategoryID != 0 {
		condition["category_id"] = service.CategoryID
	}

	// 查询符合条件的商品总数
	productDao := dao.NewProductDao(ctx)
	total, err := productDao.CountProductByCondition(condition)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}

	// 异步查询符合条件的商品信息
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		productDao = dao.NewProductDaoByDB(productDao.DB)
		products, _ = productDao.ListProductByCondition(condition, service.BasePage)
		wg.Done()
	}()
	wg.Wait()

	// 序列化
	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(total))
}

func (service *ProductService) Search(ctx context.Context) serializer.Response {
	var products []*model.Product
	var err error
	code := e.Success
	if service.PageSize == 0 {
		service.PageSize = 15
	}
	productDao := dao.NewProductDao(ctx)
	products, err = productDao.SearchProduct(service.Info, service.BasePage)
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}
	return serializer.BuildListResponse(serializer.BuildProducts(products), uint(len(products)))
}

func (service *ProductService) Show(ctx context.Context, id string) serializer.Response {
	var product *model.Product
	var err error
	code := e.Success
	productDao := dao.NewProductDao(ctx)
	pid, err := strconv.Atoi(id)
	if err != nil {
		code = e.InvalidParams
		return serializer.ErrorByCode(code, err)
	}
	product, err = productDao.GetProductByID(uint(pid))
	if err != nil {
		code = e.Error
		return serializer.ErrorByCode(code, err)
	}
	return serializer.Success(serializer.BuildProduct(product))
}
