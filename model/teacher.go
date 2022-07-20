package model

import (
	"gin-vue-admin/global"
)

type Teacher struct {
	global.GVA_MODEL
	Name        string  `json:"name" form:"name" gorm:"comment:字典名（中）"`
	Birthday    int     `json:"birthday" form:"birthday" gorm:"comment:出生日期"`
	Sex         int     `json:"sex" form:"sex" gorm:";comment:性别 1表示男，2表示女。"`
	SubjectID   uint    `json:"subjectID" form:"subjectID" gorm:"comment:科目id"`
	Subject     *Course `json:"subject" gorm:"foreignKey:SubjectID;"`
	Telephone   string  `json:"telephone" form:"telephone" gorm:"comment:电话"`
	Description string  `json:"description" form:"description" gorm:"comment:描述"`
	SysUserID   uint    `json:"sysUserID" form:"sysUserID" gorm:"comment:用户UUID"`
}
