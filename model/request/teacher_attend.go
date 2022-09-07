package request

import (
	"gin-vue-admin/global"
)

// 搜索教师考勤记录
type SearchTeacherAttend struct {
	global.GVA_MODEL
	TeacherID   uint   `json:"teacherID" form:"teacherID" gorm:"-"`
	TeacherName string `json:"teacherName" form:"teacherName" gorm:"comment:教师姓名"`
	StartTime   uint   `json:"startTime" form:"startTime" gorm:"comment:时间;" `
	EndTime     uint   `json:"endTime" form:"endTime" gorm:"comment:时间;" `
	PageInfo
}
