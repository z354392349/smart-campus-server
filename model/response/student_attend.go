package response

import (
	"gin-vue-admin/global"
)

// 学生考勤记录
type StudentAttend struct {
	global.GVA_MODEL
	StudentName string `json:"studentName" form:"studentName"`
	GradeName   string `json:"gradeName" form:"gradeName"`
	ClassName   string `json:"className" form:"className"`
	Time        uint   `json:"time" form:"time"`
	Direction   uint   `json:"direction" form:"direction"`
}
