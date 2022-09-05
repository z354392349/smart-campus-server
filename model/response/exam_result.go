package response

import "gin-vue-admin/global"

type ExamResult struct {
	global.GVA_MODEL
	ExamName    string `json:"examName" form:"examName"`       // 考试名称
	GradeName   string `json:"gradeName" form:"gradeName"`     // 年级名称
	ClassName   string `json:"className" form:"className"`     // 班级名称
	StudentName string `json:"studentName" form:"studentName"` // 学生姓名
	CourseName  string `json:"courseName" form:"courseName"`   // 科目名称
	Result      *int   `json:"result" form:"result"`           // 成绩
	Description string `json:"description" form:"description"` // 描述
}
