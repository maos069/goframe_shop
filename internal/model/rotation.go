package model

//RotationCreateUpdateBase 创建/修改Rotation基类
type RotationCreateUpdateBase struct {
	PicUrl    string      `json:"picUrl"    description:"轮播图片"`
	Link      string      `json:"link"      description:"跳转链接"`
	Sort      int         `json:"sort"      description:"排序字段"`
}

type RotationCreateInput struct{
	RotationCreateUpdateBase
}

type RotationCreateOutput struct{
	RotationId uint
}