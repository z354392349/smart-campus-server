package model

import (
	"gin-vue-admin/global"
)

// 考试项

type ExamItem struct {
	global.GVA_MODEL
	ExamID      uint   `json:"examID" form:"examID" gorm:"comment:考试;"`
	CourseName  string `json:"courseName" form:"courseName" `
	CourseID    uint   `json:"courseID" form:"courseID" gorm:"comment:考试科目ID;" `
	StartTime   int    `json:"startTime" form:"startTime" gorm:"comment:考试开始时间;" `
	EndTime     int    `json:"endTime" form:"endTime" gorm:"comment:考试结束时间;" `
	ExamRoomIDs string `json:"examRoomIDs" form:"examRoomIDs" gorm:"comment:分配的考场号用,分割"`
}

// 设置考场分配，年级ID，考试ID，考试项id，考场IDs

// type SetExamRoomItemAllot struct {
// 	ExamID      uint   `json:"examID"  form:"examID" `            // 考试ID
// 	ExamItemID  uint   `json:"examItemID"  form:"examItemID" `    // 考试项ID
// 	GradeID     uint   `json:"gradeID" form:"gradeID" `           // 年级 ID
// 	ExamRoomIDs string `json:"examRoomIDs"  form:"examRoomIDs"  ` // 考场号ID, 用,分割
// }
