package model

import (
	"gin-vue-admin/global"
)

type Student struct {
	global.GVA_MODEL
	Name        string `json:"name" form:"name" gorm:"comment:学生名称"`
	Birthday    int    `json:"birthday" form:"birthday" gorm:"comment:出生日期"`
	Sex         int    `json:"sex" form:"sex" gorm:"comment:性别 1表示男，2表示女。"`
	Telephone   string `json:"telephone" form:"telephone" gorm:"comment:家长电话"`
	GradeID     uint   `json:"gradeID" form:"gradeID" gorm:"comment:外键年级ID;" `
	ClassID     uint   `json:"classID" form:"classID" gorm:"comment:外键班级ID;"`
	SysUserID   uint   `json:"sysUserID" form:"sysUserID" gorm:"comment:用户UUID"`
	Description string `json:"description" form:"description" gorm:"comment:描述"`

	ClassName string `json:"className" form:"className" `
	GradeName string `json:"gradeName" form:"gradeName" `
}
