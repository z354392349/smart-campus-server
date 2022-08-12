package request

// 获取年级成绩 总览平均数，科目平均数
type GradeAverageResult struct {
	GradeID  uint `json:"gradeID" form:"gradeID"`
	CourseID uint `json:"courseID" form:"courseID"`
}

// 获取年级成绩 百分比
type GradePassPercent struct {
	GradeID uint `json:"gradeID" form:"gradeID"`
}
