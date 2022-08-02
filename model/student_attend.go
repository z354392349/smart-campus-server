package model

import (
	"gin-vue-admin/global"
)

// 学生考勤记录
type StudentAttend struct {
	global.GVA_MODEL
	StudentID uint   `json:"studentID" form:"studentID"`
	Time      uint   `json:"time" form:"time"`
	Place     string `json:"place" form:"place"`
	Direction uint   `json:"direction" form:"direction"`
}
