package request

import (
	"gin-vue-admin/global"
)

// TODO: 没用上，先废弃

type AllotExamRoom struct {
	global.GVA_MODEL
	StudentsID  string `json:"studentsID" form:"studentsID" gorm:"comment:学生ID"`
	ExamItemID  uint   `json:"examItemID" form:"examItemID" gorm:"comment:考试项ID;"`
	ExamID      uint   `json:"examID" form:"examID" gorm:"comment:考试;"`
	ExamRoomID  uint   `json:"examRoomID" form:"examRoomID" gorm:"comment:考试;"`
	Description string `json:"description" form:"description" gorm:"comment:描述"`
}

// 年级分页条件查询及排序结构体
// TODO: 需要改

type SearchAllotExamRoomParams struct {
	PageInfo
}
