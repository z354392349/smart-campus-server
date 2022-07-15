package request

import "gin-vue-admin/model"

// 年级分页条件查询及排序结构体

type SearchStudentParams struct {
	model.Student
	PageInfo
}
