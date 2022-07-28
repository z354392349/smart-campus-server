package response

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model"
)

// 考试项
type Exam struct {
	global.GVA_MODEL
	Name        string       `json:"name" form:"name" gorm:"comment:考试名称名称"`
	ExamItem    []ExamItem   `json:"examItem" form:"examItem"`
	GradeID     uint         `json:"gradeID" form:"gradeID" gorm:"comment:年级ID"`
	Grade       *model.Grade `json:"grade" form:"grade" gorm:"foreignKey:GradeID;"`
	Description string       `json:"description" form:"description" gorm:"comment:描述"`
}
