package response

// 车辆通行，人员通行，教师考勤，学生考勤
type GetDashboardCensusNum struct {
	PeopleAccess  int64         `json:"peopleAccess"`
	CarAccess     int64         `json:"carAccess"`
	TeacherCensus TeacherCensus `json:"teacherCensus"`
	StudentCensus StudentCensus `json:"studentCensus"`
}

// 教师考勤 - 日
type TeacherCensus struct {
	Attend int64 `json:"attend" `
	Num    int64 `json:"num"  `
}

// 学生考勤 - 日
type StudentCensus struct {
	Attend int64 `json:"attend" `
	Num    int64 `json:"num"  `
}

// 教师数量考勤
type TeacherNum struct {
	Sex string `json:"sex" `
	Num int64  `json:"num" `
}

// 考试总数
type ExamPassRate struct {
	GradeName string `json:"gradeName" `
	Total     int64  `json:"total" `
	Rate      int64  `json:"rate"`
}

// 教师考勤 - 历史
type TeacherAttendCensus struct {
	Time   string `json:"time" `   // 时间
	Attend int64  `json:"attend" ` // 出席率
	OnTime int64  `json:"onTime"`  // 准时率
}
