package request

// 获取班级成绩 成绩分析通用
type ClassResultAnalyse struct {
	ClassID  uint `json:"classID" form:"classID"`
	GradeID  uint `json:"gradeID" form:"gradeID"`
	CourseID uint `json:"courseID" form:"courseID"`
}
