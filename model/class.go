package model

import (
	"gin-vue-admin/global"
)

type Class struct {
	global.GVA_MODEL
	Name        string  `json:"name" form:"name" gorm:"column:name;comment:班级名称"`
	GradeID     uint    `json:"gradeID" form:"gradeID" gorm:"comment:外键年级ID;" `
	Grade       Grade   `gorm:"foreignKey:GradeID;"`
	TeacherID   uint    `json:"teacherID" form:"teacherID" gorm:"column:teacherID; comment:班主任"`
	Teacher     Teacher `gorm:"foreignKey:TeacherID;"`
	Description string  `json:"description" form:"description" gorm:"column:description;comment:描述"`
}
