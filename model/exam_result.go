package model

import (
	"gin-vue-admin/global"
)

// 学生成绩表，以本表为主表，查学生信息，考试信息，

type ExamResult struct {
	global.GVA_MODEL
	ExamID    uint `json:"examID" form:"examID" gorm:"comment:考试ID;"`
	Result    int  `json:"result" form:"result" gorm:"comment:成绩;" `
	CourseID  uint `json:"courseID" form:"courseID" gorm:"comment:科目ID;" `
	StudentID uint `json:"studentID" form:"studentID" gorm:"comment:学生ID;" `

	ExamName    string `json:"examName" form:"examName"`        // 考试名称 1
	CourseName  string `json:"courseName" form:"courseName"`    // 科目名称1
	StudentName string `json:"studentName" form:"studentName" ` // 学生姓名 1
	GradeName   string `json:"gradeName" form:"gradeName" `     // 年级名称
	ClassName   string `json:"className" form:"className" `     // 班级名称

}
