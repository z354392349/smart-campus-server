package model

import (
	"gin-vue-admin/global"
)

// 考试项

type ExamItem struct {
	global.GVA_MODEL
	ExamID      uint   `json:"examID" form:"examID" gorm:"comment:考试;"`
	CourseID    uint   `json:"courseID" form:"courseID" gorm:"comment:考试科目ID;" `
	StartTime   int    `json:"startTime" form:"startTime" gorm:"comment:考试开始时间;" `
	EndTime     int    `json:"endTime" form:"endTime" gorm:"comment:考试结束时间;" `
	ExamRoomIDs string `json:"examRoomIDs" form:"examRoomIDs" gorm:"comment:分配的考场号用,分割"`
}
