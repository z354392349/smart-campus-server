package model

import (
	"gin-vue-admin/global"
)

type Grade struct {
	global.GVA_MODEL
	Name        string `json:"name" form:"name" gorm:"comment:年级名称"`
	Description string `json:"description" form:"description" gorm:"comment:描述"`
}
