package request

import "gin-vue-admin/model"

// 年级分页条件查询及排序结构体

type SearchStudentParams struct {
	model.Student
	PageInfo
}

// 批量修改学生年级和班级

type SetStudentsGradeAndClass struct {
	StudentsID []int `json:"studentsID" form:"studentsID"` // 学生id
	GradeID    uint  `json:"gradeID" form:"gradeID"`       // 年级id
	ClassID    uint  `json:"classID" form:"classID"`       // 班级id
}
