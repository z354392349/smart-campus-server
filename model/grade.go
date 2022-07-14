package model

import (
	"gin-vue-admin/global"
)

type Grade struct {
	global.GVA_MODEL
	Name        string `json:"name" form:"name" gorm:"column:name;comment:字典名（中）"`                  // 字典名（中）
	Description string `json:"description" form:"description" gorm:"column:description;comment:描述"` // 描述
}
