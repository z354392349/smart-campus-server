package response

import (
	"gin-vue-admin/global"
)

// 车辆通行记录 教师姓名, 车牌号，时间，地点， 方向
type CarAccess struct {
	global.GVA_MODEL
	TeacherName string `json:"teacherName" form:"teacherName"`
	CarNum      string `json:"carNum" form:"carNum"`
	Time        uint   `json:"startTime" form:"startTime"`
	Place       string `json:"place" form:"place"`
	Direction   string `json:"direction" form:"direction"`
}
