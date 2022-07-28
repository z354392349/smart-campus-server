package response

import "gin-vue-admin/global"

// 考试项

type ExamItem struct {
	global.GVA_MODEL
	ExamID      uint   `json:"examID" form:"examID"`
	CourseID    uint   `json:"courseID" form:"courseID"`
	StartTime   int    `json:"startTime" form:"startTime"`
	EndTime     int    `json:"endTime" form:"endTime"`
	ExamRoomIDs string `json:"examRoomIDs" form:"examRoomIDs"`
	CourseName  string `json:"courseName" form:"courseName"`
}
