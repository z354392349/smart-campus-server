package response

// 车辆通行，人员通行，教师考勤，学生考勤
type GetDashboardCensusNum struct {
	PeopleAccess  int64       `json:"peopleAccess"`
	CarAccess     int64       `json:"carAccess"`
	TeacherCensus interface{} `json:"TeacherCensus"`
	StudentCensus interface{} `json:"studentCensus"`
}
