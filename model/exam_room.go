package model

import (
	"gin-vue-admin/global"
)

type ExamRoom struct {
	global.GVA_MODEL
	Name        string   `json:"name" form:"name" gorm:"comment:考场名称"`
	Address     string   `json:"address" form:"address" gorm:"comment:考场地址"`
	Amount      int      `json:"amount" form:"amount" gorm:"comment:考场容量"`
	Teacher     *Teacher `json:"teacher" form:"teacher" gorm:"foreignKey:TeacherID;"`
	TeacherID   uint     `json:"teacherID" form:"teacherID"`
	Description string   `json:"description" form:"description" gorm:"comment:描述"`
}
