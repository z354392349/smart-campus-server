package response

import (
	"gin-vue-admin/global"
)

// 教师通行记录
type TeacherAccess struct {
	global.GVA_MODEL
	TeacherName string `json:"teacherName" form:"teacherName"`
	Time        uint   `json:"time" form:"time"`
	Place       string `json:"place" form:"place"`
	Direction   uint   `json:"direction" form:"direction"`
}
