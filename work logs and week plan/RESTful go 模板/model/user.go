package model

type User struct{
	Id int64 `form:"id" json:"id"`
	Name string `xorm:"name" form:"name" json:"name"`
	Birthday int `xorm:"birthday" form:"birthday" json:"birthday"`
	CreateTime int64 `xorm:"create_time" form:"createTime" json:"createTime"`
}