package model

import (
	"gin-vue-admin/global"
)

type Teacher struct {
	global.GVA_MODEL
	Name        string `json:"name" form:"name" gorm:"column:name;comment:字典名（中）"`
	Birthday    int    `json:"birthday" form:"birthday" gorm:"column:birthday;comment:出生日期"`
	Sex         int    `json:"sex" form:"sex" gorm:"column:sex;comment:性别 1表示男，2表示女。"`
	Telephone   string `json:"telephone" form:"telephone" gorm:"column:telephone;comment:电话"`
	Description string `json:"description" form:"description" gorm:"column:description;comment:描述"`
}
