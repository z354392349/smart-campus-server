package model

import (
	"gin-vue-admin/global"
)

// TODO:学生表要加一个 状态， 记录毕业，还是离校，还是在上学

type Student struct {
	global.GVA_MODEL
	Name        string `json:"name" form:"name" gorm:"comment:学生名称"`
	Birthday    int    `json:"birthday" form:"birthday" gorm:"comment:出生日期"`
	Sex         int    `json:"sex" form:"sex" gorm:"comment:性别 1表示男，2表示女。"`
	Telephone   string `json:"telephone" form:"telephone" gorm:"comment:家长电话"`
	GradeID     uint   `json:"gradeID" form:"gradeID" gorm:"comment:外键年级ID;" `
	Grade       *Grade `json:"grade" form:"grade" gorm:"foreignKey:GradeID;"`
	ClassID     uint   `json:"classID" form:"classID" gorm:"comment:外键班级ID;" `
	Class       *Class `json:"class" form:"class"  gorm:"foreignKey:ClassID;"`
	SysUserID   uint   `json:"sysUserID" form:"sysUserID" gorm:"comment:用户UUID"`
	Description string `json:"description" form:"description" gorm:"comment:描述"`
}
