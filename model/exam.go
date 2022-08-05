package model

import (
	"gin-vue-admin/global"
)

type Exam struct {
	global.GVA_MODEL
	Name     string     `json:"name" form:"name" gorm:"comment:考试名称名称"`
	ExamItem []ExamItem `json:"examItem" form:"examItem"`
	GradeID  uint       `json:"gradeID" form:"gradeID" gorm:"comment:年级ID"`
	Grade    *Grade     `json:"grade" form:"grade" gorm:"foreignKey:GradeID;"`

	Description string `json:"description" form:"description" gorm:"comment:描述"`
}
