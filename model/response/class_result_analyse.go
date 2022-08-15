package response

// 获取年级下各 班级平均成绩
type StudentTotalResult struct {
	StudentName string  `json:"studentName" ` // 班级名称
	Total       float64 `json:"total"`        // 数量
}

// 获取每个学生的成绩
type StudentTotalResultHistory struct {
	ExamName    string `json:"examName" ` // 考试名称
	StudentID   uint   `json:"studentID" `
	StudentName string `json:"studentName" ` // 学生名称
	Result      string `json:"result"`       // 数量
}
