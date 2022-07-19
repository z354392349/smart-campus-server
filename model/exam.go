package model

import (
	"gin-vue-admin/global"
)

// TODO:
// 这个前台 需要增加详情
// TODO: 发布考试，默认全年级学生参加，将本次 成绩设置成为null， 这里还却一个年级字段

type Exam struct {
	global.GVA_MODEL
	Name        string     `json:"name" form:"name" gorm:"comment:考试名称名称"`
	ExamItem    []ExamItem `json:"examItem" form:"examItem" `
	ClassIDs    string     `json:"gradeID" form:"gradeID" gorm:"comment:班级ID，分割"`
	Description string     `json:"description" form:"description" gorm:"comment:描述"`
}
