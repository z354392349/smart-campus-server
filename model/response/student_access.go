package response

import (
	"gin-vue-admin/global"
)

// 学生通行记录
type StudentAccess struct {
	global.GVA_MODEL
	StudentName string `json:"studentName" form:"studentName"`
	GradeName   string `json:"gradeName" form:"gradeName"`
	ClassName   string `json:"className" form:"className"`
	Time        uint   `json:"time" form:"time"`
	Place       string `json:"place" form:"place"`
	Direction   uint   `json:"direction" form:"direction"`
}
