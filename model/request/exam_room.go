package request

import "gin-vue-admin/model"

type SearchExamRoomParams struct {
	model.Student
	PageInfo
}
