package model

import (
	"gin-vue-admin/global"
)

// TODO: 缺个班长, 等学生表完成的
type Class struct {
	global.GVA_MODEL
	Name        string   `json:"name" form:"name" gorm:"column:name;comment:班级名称"`
	GradeID     uint     `json:"gradeID" form:"gradeID" gorm:"comment:外键年级ID;" `
	Grade       *Grade   `json:"grade" gorm:"foreignKey:GradeID;"`
	TeacherID   uint     `json:"teacherID" form:"teacherID" gorm:"column:teacherID; comment:班主任"`
	Teacher     *Teacher `json:"teacher" gorm:"foreignKey:TeacherID;"`
	Description string   `json:"description" form:"description" gorm:"column:description;comment:描述"`
}
