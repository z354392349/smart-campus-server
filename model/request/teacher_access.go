package request

import (
	"gin-vue-admin/global"
)

// 搜索教师通行记录
type SearchTeacherAccess struct {
	global.GVA_MODEL
	TeacherName string `json:"teacherName" form:"teacherName" gorm:"comment:教师姓名"`
	StartTime   uint   `json:"startTime" form:"startTime" gorm:"comment:时间;" `
	EndTime     uint   `json:"endTime" form:"endTime" gorm:"comment:时间;" `
	PageInfo
}
