package model

import (
	"gin-vue-admin/global"
)

// 学生成绩表，以本表为主表，查学生信息，考试信息，

type ExamResult struct {
	global.GVA_MODEL
	ExamID     uint `json:"examID" form:"examID" gorm:"comment:考试ID;"`
	ExamItemID uint `json:"examItemID" form:"examItemID" gorm:"comment:考试项ID;" `
	StudentID  uint `json:"studentID" form:"studentID" gorm:"comment:学生ID;" `
	Result     int  `json:"result" form:"result" gorm:"comment:成绩;" `

	GradeID     string `json:"gradeID" form:"gradeID" gorm:"-"`         // 年级名称 根据StudentID
	ClassName   string `json:"className" form:"className" gorm:"-"`     // 班级名称 根据StudentID
	StudentName string `json:"studentName" form:"studentName" gorm:"-"` // 学生姓名 根据StudentID
	CourseName  string `json:"courseName" form:"courseName" gorm:"-"`   // 科目名称 根据ExamItemID

}
