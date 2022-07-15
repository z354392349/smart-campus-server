package model

import (
	"gin-vue-admin/global"
)

// TODO:[]Teacher 不知道是否要加指针, 2切片怎么存储没有掌握

//	TeacherID   []uint     `json:"teacherID" form:"teacherID" gorm:"column:teacherID; comment:监考老师ID"`
//	Teacher     *[]Teacher `json:"teacher" gorm:"foreignKey:TeacherID;"`
type ReleaseExam struct {
	global.GVA_MODEL
	Name        string   `json:"name" form:"name" gorm:"column:name;comment:考试名称名称"`
	Description string   `json:"description" form:"description" gorm:"column:description;comment:描述"`
	GradeID     []uint   `json:"gradeID" form:"gradeID" gorm:"comment:参加考试年级ID;" `
	Grade       *[]Grade `json:"grade" gorm:"foreignKey:GradeID;"`
	StartTime   int      `json:"startTime" form:"startTime" gorm:"column:startTime;comment:开始时间"`
	EndTime     int      `json:"endTime" form:"endTime" gorm:"column:endTime;comment:结束时间"`
}
