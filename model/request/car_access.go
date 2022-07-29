package request

import (
	"gin-vue-admin/global"
)

// 搜索车辆通行记录
type SearchCarAccess struct {
	global.GVA_MODEL
	TeacherName string `json:"teacherID" form:"teacherID" gorm:"comment:教师姓名"`
	CarNum      string `json:"carNum" form:"carNum" gorm:"comment:车牌号;"`
	StartTime   uint   `json:"startTime" form:"startTime" gorm:"comment:时间;" `
	EndTime     uint   `json:"endTime" form:"endTime" gorm:"comment:时间;" `
	PageInfo
}
