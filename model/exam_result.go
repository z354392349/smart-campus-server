package model

import (
	"gin-vue-admin/global"
)

// 学生成绩表，以本表为主表，查学生信息，考试信息，

type ExamResult struct {
	global.GVA_MODEL
	ExamID    uint     `json:"examID" form:"examID" gorm:"comment:考试ID;"`
	Course    *Course  `json:"course" form:"course" gorm:"foreignKey:CourseID;comment:考试科目;" `
	CourseID  uint     `json:"courseID" form:"courseID" gorm:"comment:考试科目ID;" `
	Student   *Student `json:"student" form:"course" gorm:"foreignKey:StudentID;comment:学生;" `
	StudentID uint     `json:"studentID" form:"studentID" gorm:"comment:学生ID;" `
	Result    int      `json:"result" form:"result" gorm:"comment:成绩;" `
}
