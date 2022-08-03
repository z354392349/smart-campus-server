package model

import (
	"gin-vue-admin/global"
)

// 学生成绩表，以本表为主表，查学生信息，考试信息，

type ExamResult struct {
	global.GVA_MODEL
	ExamID      uint   `json:"examID" form:"examID" gorm:"comment:考试ID;"`
	Result      *int   `json:"result" form:"result"`
	CourseID    uint   `json:"courseID" form:"courseID" gorm:"comment:科目ID;" `
	StudentID   uint   `json:"studentID" form:"studentID" gorm:"comment:学生ID;" `
	Description string `json:"description" form:"description" gorm:"comment:描述"`
}
