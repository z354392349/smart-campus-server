package request

// 考场分配结果分页条件查询及排序结构体

type SearchAllotExamRoomParams struct {
	ClassID uint   `json:"classID" form:"classID"`
	GradeID uint   `json:"gradeID" form:"gradeID"`
	Name    string `json:"name" form:"name"`
	PageInfo
}
