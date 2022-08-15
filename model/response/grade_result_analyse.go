package response

// 获取年级下各 班级平均成绩
type GradeAverageResult struct {
	ClassName string  `json:"className" ` // 班级名称
	Num       float64 `json:"Num"`        // 数量
}

// 获取年级下各 班级平均成绩
type GradeAverageResultHistory struct {
	ExamName  string `json:"examName" `  // 考试名称
	ClassName string `json:"className" ` // 班级名称
	Result    string `json:"result"`     // 数量
}
