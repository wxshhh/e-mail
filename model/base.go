package model

type BasePage struct {
	PageNum  int `form:"page_num"`  // 第几页
	PageSize int `form:"page_size"` // 一页有几个
}
