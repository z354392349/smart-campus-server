package request

import (
	"gin-vue-admin/global"
)

// 分配考场新增，修改

type AllotExamRoom struct {
	global.GVA_MODEL
	StudentsID  string `json:"studentsID" form:"studentsID" gorm:"comment:学生ID"`
	CourseID    uint   `json:"courseID" form:"courseID" gorm:"comment:考试科目ID;" `
	ExamID      uint   `json:"examID" form:"examID" gorm:"comment:考试;"`
	ExamRoomID  uint   `json:"examRoomID" form:"examRoomID" gorm:"comment:考试;"`
	Description string `json:"description" form:"description" gorm:"comment:描述"`
}

// 年级分页条件查询及排序结构体

type SearchAllotExamRoomParams struct {
	PageInfo
}
