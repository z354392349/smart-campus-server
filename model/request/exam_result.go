package request

// 年级分页条件查询及排序结构体

type SearchExamResultParams struct {
	Name    string `json:"name" form:"name"`
	GradeID uint   `json:"gradeID" form:"gradeID"`
	ClassID uint   `json:"classID" form:"classID"`
	PageInfo
}
