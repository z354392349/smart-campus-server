package request

import "gin-vue-admin/model"

type SearchExamRoomParams1 struct {
	model.Student
	PageInfo
}

// 设置考场分配，年级ID，考试ID，考试项id，考场IDs

type SetExamRoomItemAllot struct {
	ExamID      uint   `json:"examID"  form:"examID" `            // 考试ID
	ExamItemID  uint   `json:"examItemID"  form:"examItemID" `    // 考试项ID
	GradeID     uint   `json:"gradeID" form:"gradeID" `           // 年级 ID
	ExamRoomIDs string `json:"examRoomIDs"  form:"examRoomIDs"  ` // 考场号ID, 用,分割
}
