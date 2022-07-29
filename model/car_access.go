package model

import (
	"gin-vue-admin/global"
)

// 车辆通行记录 教师姓名, 车牌号，时间，地点， 方向
type CarAccess struct {
	global.GVA_MODEL
	TeacherID uint   `json:"teacherID" form:"teacherID" gorm:"comment:教师ID"`
	CarNum    string `json:"carNum" form:"carNum" gorm:"comment:车牌号;"`
	Time      uint   `json:"time" form:"time" gorm:"comment:时间;" `
	Place     string `json:"place" form:"place" gorm:"comment:地点;"`
	Direction string `json:"direction" form:"direction" gorm:"comment:方向1是进2是出;"`
}
