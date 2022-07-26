package model

import (
	"gin-vue-admin/global"
)

// 考场分配结果
type AllotExamRoom struct {
	global.GVA_MODEL
	StudentID  uint `json:"studentID" form:"studentID" gorm:"comment:学生ID"`
	ExamID     uint `json:"examID" form:"examID" gorm:"comment:考试;"`
	ExamItemID uint `json:"examItemID" form:"examItemID" gorm:"comment:考试项目ID;" `
	ExamRoomID uint `json:"examRoomID" form:"examRoomID" gorm:"comment:考试;"`

	Exam     *Exam     `json:"exam" form:"exam" `         // gorm:"foreignKey:ExamID;"
	ExamItem *ExamItem `json:"examItem" form:"examItem" ` // gorm:"foreignKey:ExamItemID;"
	ExamRoom *ExamRoom `json:"examRoom" form:"examRoom"`  //  gorm:"foreignKey:ExamID;"
	Student  *Student  `json:"students" form:"students" ` // gorm:"foreignKey:StudentID;"
}
