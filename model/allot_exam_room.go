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
}
