package request

import (
	"gin-vue-admin/global"
)

type SearchStudentAttend struct {
	global.GVA_MODEL
	GradeID     uint   `json:"gradeID" form:"gradeID"`
	ClassID     uint   `json:"classID" form:"classID"`
	StudentName string `json:"studentName" form:"studentName"`
	StartTime   uint   `json:"startTime" form:"startTime"`
	EndTime     uint   `json:"endTime" form:"endTime"`
	PageInfo
}
