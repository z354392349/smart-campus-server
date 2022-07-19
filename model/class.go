package model

import (
	"gin-vue-admin/global"
)

// TODO: 缺个班长, 等学生表完成的

type Class struct {
	global.GVA_MODEL
	Name        string   `json:"name" form:"name" gorm:"comment:班级名称"`
	GradeID     uint     `json:"gradeID" form:"gradeID" gorm:"comment:外键年级ID;" `
	Grade       *Grade   `json:"grade" gorm:"foreignKey:GradeID;"`
	TeacherID   uint     `json:"teacherID" form:"teacherID" gorm:"comment:班主任"`
	Teacher     *Teacher `json:"teacher" gorm:"foreignKey:TeacherID;"`
	Description string   `json:"description" form:"description" gorm:"comment:描述"`
}
