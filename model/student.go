package model

import (
	"gin-vue-admin/global"
)

// TODO:学生表要加一个 状态， 记录毕业，还是离校，还是在上学
type Student struct {
	global.GVA_MODEL
	Name      string `json:"name" form:"name" gorm:"column:name;comment:学生名称"`
	Birthday  int    `json:"birthday" form:"birthday" gorm:"column:birthday;comment:出生日期"`
	Sex       int    `json:"sex" form:"sex" gorm:"column:sex;comment:性别 1表示男，2表示女。"`
	Telephone string `json:"telephone" form:"telephone" gorm:"column:telephone;comment:家长电话"`
	Nation    string `json:"nation" form:"nation" gorm:"column:nation;comment:民族"`
	GradeID   uint   `json:"gradeID" form:"gradeID" gorm:"comment:外键年级ID;" `
	Grade     *Grade `gorm:"foreignKey:GradeID;"`
	ClassID   uint   `json:"classID" form:"classID" gorm:"comment:外键班级ID;" `
	Class     *Class `gorm:"foreignKey:ClassID;"`

	Description string `json:"description" form:"description" gorm:"column:description;comment:描述"`
}
