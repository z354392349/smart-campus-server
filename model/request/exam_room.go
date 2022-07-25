package request

import "gin-vue-admin/model"

type SearchExamRoomParams struct {
	model.Student
	PageInfo
}

// 班级设置班主任

type SetExamRoomTeacher struct {
	TeacherID  uint `json:"teacherID" gorm:"comment:教师ID"`
	ExamRoomID uint `json:"examRoomID" gorm:"comment:考场ID;" `
}
