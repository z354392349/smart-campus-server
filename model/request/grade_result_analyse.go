package request

// 获取年级成绩 成绩分析通用
type GradeResultAnalyse struct {
	GradeID  uint `json:"gradeID" form:"gradeID"`
	CourseID uint `json:"courseID" form:"courseID"`
}
