package request

// 成绩分页条件查询及排序结构体
type SearchExamResultParams struct {
	Name    string `json:"name" form:"name"`
	GradeID uint   `json:"gradeID" form:"gradeID"`
	ClassID uint   `json:"classID" form:"classID"`
	ExamID  uint   `json:"examID" form:"examID" gorm:"comment:考试ID;"`
	PageInfo
}

//设置成绩备注
type SetExamResultParams struct {
	ID          uint   `json:"ID" form:"ID"`
	Result      int    `json:"result" form:"result" `
	Description string `json:"description" form:"description"`
}
