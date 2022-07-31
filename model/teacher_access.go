package model

import (
	"gin-vue-admin/global"
)

// 教师通行记录

type TeacherAccess struct {
	global.GVA_MODEL
	TeacherID uint   `json:"teacherID" form:"teacherID" gorm:"comment:教师ID;"`
	Time      uint   `json:"time" form:"time" gorm:"comment:时间;" `
	Place     string `json:"place" form:"place" gorm:"comment:地点;"`
	Direction uint   `json:"direction" form:"direction" gorm:"comment:方向1是进2是出;"`
}
