package response

// 获取学生的每一科目成绩
type StudentCourseResult struct {
	StudentName string  `json:"studentName" ` // 班级名称
	Total       float64 `json:"total"`        // 数量
}

// type StudentTotalResultHistory struct {
// 	ExamName    string `json:"examName" ` // 考试名称
// 	StudentID   uint   `json:"studentID" `
// 	StudentName string `json:"studentName" ` // 学生名称
// 	Result      string `json:"result"`       // 数量
// }
