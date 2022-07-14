package model

import (
	"gin-vue-admin/global"
)

type Course struct {
	global.GVA_MODEL
	Name        string `json:"name" form:"name" gorm:"column:name;comment:课程名称"`
	Description string `json:"description" form:"description" gorm:"column:description;comment:描述"`
}
