package request

import "gin-vue-admin/model"

// 教师分页条件查询及排序结构体

type SearchTeacherParams struct {
	model.Teacher
	PageInfo
}
