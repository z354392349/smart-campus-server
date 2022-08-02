package response

import (
	"gin-vue-admin/global"
)

// 教师通行记录
type TeacherAttend struct {
	global.GVA_MODEL
	TeacherName string `json:"teacherName" form:"teacherName"`
	Time        uint   `json:"time" form:"time"`
	Direction   uint   `json:"direction" form:"direction"`
}
