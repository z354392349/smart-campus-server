package model

import (
	"gin-vue-admin/global"
)

// 学生通行记录

type StudentAccess struct {
	global.GVA_MODEL
	StudentID uint   `json:"studentID" form:"studentID" gorm:"comment:学生ID;"`
	Time      uint   `json:"time" form:"time" gorm:"comment:时间;" `
	Place     string `json:"place" form:"place" gorm:"comment:地点;"`
	Direction uint   `json:"direction" form:"direction" gorm:"comment:方向1是进2是出;"`
}
