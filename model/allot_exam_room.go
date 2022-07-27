package model

import (
	"gin-vue-admin/global"
)

// 考场分配表 结果表
type AllotExamRoom struct {
	global.GVA_MODEL
	StudentID  uint `json:"studentID" form:"studentID" gorm:"comment:学生ID"`
	ExamID     uint `json:"examID" form:"examID" gorm:"comment:考试ID;"`
	CourseID   uint `json:"courseID" form:"courseID" gorm:"comment:考试科目ID;" `
	ExamRoomID uint `json:"examRoomID" form:"examRoomID" gorm:"comment:考场ID;"`

	StudentName  string `json:"studentName" form:"studentName"`   // 学生姓名
	ExamName     string `json:"examName" form:"examName"`         // 考试名称
	CourseName   string `json:"courseName" form:"courseName"`     // 科目名称
	ExamRoomName string `json:"examRoomName" form:"examRoomName"` // 考场房间名称
	GradeName    string `json:"gradeName" form:"gradeName"`       // 考场房间名称
	ClassName    string `json:"className" form:"className"`       // 考场房间名称

}
