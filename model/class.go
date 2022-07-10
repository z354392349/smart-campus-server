package model

import (
	"gin-vue-admin/global"
)

// 如果含有time.Time 请自行import time包
type Class struct {
	global.GVA_MODEL
	Name        string `json:"name" form:"name" gorm:"column:name;comment:字典名（中）"`                  // 字典名（中）
	Description string `json:"description" form:"description" gorm:"column:description;comment:描述"` // 描述
	Master      string `json:"Master" form:"Master" gorm:"column:Master;comment:班主任"`               // 描述
	Grade       Grade  `gorm:"comment:外键年级ID" `                                                     // 描述
}
