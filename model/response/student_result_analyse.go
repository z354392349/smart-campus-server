package response

// 获取学生的每一科目成绩
type StudentCourseResult struct {
	CourseName string  `json:"courseName" ` // 班级名称
	Result     float64 `json:"result"`      // 数量
}

type StudentExamTotalResultHistory struct {
	ExamName string `json:"examName" ` // 考试名称
	Total    string `json:"total"`     // 数量
}
