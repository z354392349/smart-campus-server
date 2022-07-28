package response

import (
	"gin-vue-admin/global"
)

// 考场分配表 结果表
type AllotExamRoom struct {
	global.GVA_MODEL
	StudentName  string `json:"studentName" form:"studentName"`   // 学生姓名
	ExamName     string `json:"examName" form:"examName"`         // 考试名称
	CourseName   string `json:"courseName" form:"courseName"`     // 科目名称
	ExamRoomName string `json:"examRoomName" form:"examRoomName"` // 考场房间名称
	GradeName    string `json:"gradeName" form:"gradeName"`       // 考场房间名称
	ClassName    string `json:"className" form:"className"`       // 考场房间名称
}
