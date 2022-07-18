package model

import (
	"gin-vue-admin/global"
)

type ExamItem struct {
	global.GVA_MODEL
	ExamID    uint    `json:"examID" form:"examID" gorm:"comment:考试;"`
	Course    *Course `json:"course" form:"course" gorm:"foreignKey:CourseID;comment:考试科目;" `
	CourseID  uint    `json:"courseID" form:"courseID" gorm:"comment:考试科目ID;" `
	StartTime int     `json:"startTime" form:"startTime" gorm:"comment:考试开始时间;" `
	EndTime   int     `json:"endTime" form:"endTime" gorm:"comment:考试结束时间;" `
}
