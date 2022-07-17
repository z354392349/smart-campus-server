package model

import (
	"gin-vue-admin/global"
)

// TODO:
// 这个前台 需要增加详情

type ExamItem struct {
	global.GVA_MODEL
	ExamID    uint    `json:"examID" form:"examID" gorm:"comment:考试;"`
	Course    *Course `json:"course" form:"course" gorm:"foreignKey:CourseID;comment:考试科目;" `
	CourseID  uint    `json:"courseID" form:"courseID" gorm:"comment:考试科目ID;" `
	StartTime int     `json:"startTime" form:"startTime" gorm:"comment:考试开始时间;" `
	EndTime   int     `json:"endTime" form:"endTime" gorm:"comment:考试结束时间;" `
}

// TODO: 发布考试，默认全年级学生参加，将本次 成绩设置成为null， 这里还却一个年级字段
type Exam struct {
	global.GVA_MODEL
	Name        string     `json:"name" form:"name" gorm:"column:name;comment:考试名称名称"`
	ExamItem    []ExamItem `json:"examItem" form:"examItem" `
	Description string     `json:"description" form:"description" gorm:"column:description;comment:描述"`
}
