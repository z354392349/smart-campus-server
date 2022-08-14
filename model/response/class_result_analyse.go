package response

// 获取年级下各 班级平均成绩
type StudentTotalResult struct {
	StudentName string  `json:"studentName" ` // 班级名称
	Total       float64 `json:"total"`        // 数量
}

// 获取年级下各 班级平均成绩
// type GradeAverageResultHistory1 struct {
// 	ExamName  string `json:"examName" `  // 班级名称
// 	ClassName string `json:"className" ` // 班级名称
// 	Result    string `json:"result"`     // 数量
// }
