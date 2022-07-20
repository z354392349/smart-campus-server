package model

import (
	"gin-vue-admin/global"
)

// 分配考场

type AllotExamRoom struct {
	global.GVA_MODEL
	Student     Student   `json:"students" form:"students" gorm:"foreignKey:StudentID;" `
	StudentID   uint      `json:"studentID" form:"studentID" gorm:"comment:学生ID"`
	Course      *Course   `json:"course" form:"course" gorm:"foreignKey:CourseID;" `
	CourseID    uint      `json:"courseID" form:"courseID" gorm:"comment:考试科目ID;" `
	Exam        *Exam     `json:"exam" form:"exam" gorm:"foreignKey:ExamID;"`
	ExamID      uint      `json:"examID" form:"examID" gorm:"comment:考试;"`
	ExamRoom    *ExamRoom `json:"examRoom" form:"examRoom" gorm:"foreignKey:ExamID;"`
	ExamRoomID  uint      `json:"examRoomID" form:"examRoomID" gorm:"comment:考试;"`
	Description string    `json:"description" form:"description" gorm:"comment:描述"`
}
