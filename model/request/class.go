package request

import "gin-vue-admin/model"

// 年级分页条件查询及排序结构体

type SearchClassParams struct {
	model.Class
	PageInfo
}

// 班级设置班长

type SetClassMonitor struct {
	StudentID uint `json:"studentID" gorm:"comment:学生ID"`
	ClassID   uint `json:"classID" gorm:"comment:班级ID;" `
}

// 班级设置班主任

type SetClassTeacher struct {
	TeacherID uint `json:"teacherID" gorm:"comment:学生ID"`
	ClassID   uint `json:"classID" gorm:"comment:班级ID;" `
}
