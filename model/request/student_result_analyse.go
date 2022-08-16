package request

// 获取学生成绩 成绩分析通用
type StudentResultAnalyse struct {
	ClassID   uint `json:"classID" form:"classID"`
	GradeID   uint `json:"gradeID" form:"gradeID"`
	CourseID  uint `json:"courseID" form:"courseID"`
	StudentID uint `json:"studentID" form:"studentID"`
}
