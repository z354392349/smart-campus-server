package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
)

// 教师考勤
type teacherCensus struct {
	Attend int64 `json:"attend" `
	Num    int64 `json:"num"  `
}

// 学生考勤
type studentCensus struct {
	Attend int64 `json:"attend" `
	Num    int64 `json:"num"  `
}

// 学生考勤
type TeacherNum struct {
	Sex string `json:"sex" `
	Num int64  `json:"num" `
}

// @Author: 张佳伟
// @Function:GetDashboardCensusNum
// @Description: 获取通行数量，考勤
// @Router:/dashboard/getDashboardCensusNum
// @Date:2022/08/08 21:02:59

func GetDashboardCensusNum() (err error, carAccess int64, peopleAccess int64, teacherCensus teacherCensus, studentCensus studentCensus) {

	// 获取当日车辆通行数量
	if err = global.GVA_DB.Model(&model.CarAccess{}).Where("time >=  ? and time <= ? ", 1667318400, 1667404799).Count(&carAccess).Error; err != nil {
		err = errors.New("车辆通行数量获取失败")
		return
	}

	// 获取当日人员通行数量
	var teacherAccess int64
	var studentAccess int64
	if err = global.GVA_DB.Model(&model.TeacherAccess{}).Where("time >=  ? and time <= ? ", 1667318400, 1667404799).Count(&teacherAccess).Error; err != nil {
		err = errors.New("教师通行数量获取失败")
		return
	}
	if err = global.GVA_DB.Model(&model.StudentAccess{}).Where("time >=  ? and time <= ? ", 1667318400, 1667404799).Count(&studentAccess).Error; err != nil {
		err = errors.New("学生通行数量获取失败")
		return
	}
	peopleAccess = teacherAccess + studentAccess

	// 获取教师考勤
	selectSql := "select * from (select teacher_id, time  from teacher_accesses union all select teacher_id, time  from car_accesses) tb where time >=  ? and time <= ?  group by  tb.teacher_id"
	global.GVA_DB.Debug().Raw("select count(*) as total from ("+selectSql+") tb2", 1667318400, 1667404799).Scan(&teacherCensus.Attend)

	if err = global.GVA_DB.Model(&model.Teacher{}).Count(&teacherCensus.Num).Error; err != nil {
		err = errors.New("教师数量获取失败")
		return
	}

	// 学生考勤
	if err = global.GVA_DB.Model(&model.StudentAccess{}).Debug().Where("time >=  ? and time <= ? ", 1667318400, 1667404799).Group("student_id").Count(&studentCensus.Attend).Error; err != nil {
		err = errors.New("教师考勤数量获取失败")
		return
	}

	if err = global.GVA_DB.Model(&model.Student{}).Count(&studentCensus.Num).Error; err != nil {
		err = errors.New("学生数量获取失败")
		return
	}

	return err, carAccess, peopleAccess, teacherCensus, studentCensus
}

// @Author: 张佳伟
// @Function:GetTeacherNum
// @Description:获取教师数量区分男女
// @Router:/dashboard/getTeacherNum
// @Date:2022/08/08 21:20:46

func GetTeacherNum() (err error, list interface{}) {
	var teacherNum []TeacherNum
	if err = global.GVA_DB.Model(&model.Teacher{}).Select("sex, count(1) as num ").Group("sex").Find(&teacherNum).Error; err != nil {
		err = errors.New("获取教师性别统计数量失败")
		return
	}

	return err, teacherNum
}
