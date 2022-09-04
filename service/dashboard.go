package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/response"
	"gin-vue-admin/utils"
)

// @Author: 张佳伟
// @Function:GetDashboardCensusNum
// @Description: 获取通行数量，考勤
// @Router:/dashboard/getDashboardCensusNum
// @Date:2022/08/08 21:02:59

func GetDashboardCensusNum() (err error, carAccess int64, peopleAccess int64, teacherCensus response.TeacherCensus, studentCensus response.StudentCensus) {
	var startTime int = 1667318400
	var endTime int = 1667404799

	// 获取当日车辆通行数量
	if err = global.GVA_DB.Model(&model.CarAccess{}).Where("time >=  ? and time <= ? ", startTime, endTime).Count(&carAccess).Error; err != nil {
		err = errors.New("车辆通行数量获取失败")
		return
	}

	// 获取当日人员通行数量
	var teacherAccess int64
	var studentAccess int64
	if err = global.GVA_DB.Model(&model.TeacherAccess{}).Where("time >=  ? and time <= ? ", startTime, endTime).Count(&teacherAccess).Error; err != nil {
		err = errors.New("教师通行数量获取失败")
		return
	}
	if err = global.GVA_DB.Model(&model.StudentAccess{}).Where("time >=  ? and time <= ? ", startTime, endTime).Count(&studentAccess).Error; err != nil {
		err = errors.New("学生通行数量获取失败")
		return
	}
	peopleAccess = teacherAccess + studentAccess

	// 获取教师考勤
	selectSql := "select * from (select teacher_id, time  from teacher_accesses union all select teacher_id, time  from car_accesses) tb where time >=  ? and time <= ?  group by  tb.teacher_id"
	global.GVA_DB.Debug().Raw("select count(*) as total from ("+selectSql+") tb2", startTime, endTime).Scan(&teacherCensus.Attend)

	if err = global.GVA_DB.Model(&model.Teacher{}).Count(&teacherCensus.Num).Error; err != nil {
		err = errors.New("教师数量获取失败")
		return
	}

	// 学生考勤
	if err = global.GVA_DB.Model(&model.StudentAccess{}).Debug().Where("time >=  ? and time <= ? ", startTime, endTime).Group("student_id").Count(&studentCensus.Attend).Error; err != nil {
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
	var teacherNum []response.TeacherNum
	if err = global.GVA_DB.Model(&model.Teacher{}).Select("sex, count(1) as num ").Group("sex").Find(&teacherNum).Error; err != nil {
		err = errors.New("获取教师性别统计数量失败")
		return
	}

	return err, teacherNum
}

// @Author: 张佳伟
// @Function:GetTeacherAttendCensus
// @Description:获取教师考勤率，准点率
// @Router /dashboard/getTeacherNum
// @Date:2022/08/09 17:55:29

func GetTeacherAttendCensus() (err error, list interface{}) {

	//  获取教师总数
	var teacherTotal int64
	if err = global.GVA_DB.Model(&model.Teacher{}).Count(&teacherTotal).Error; err != nil {
		err = errors.New("获取教师总数失败")
		return
	}

	// 获取考勤率
	sql := "select time, count(*) as 	attend from  (select time, teacher_id from (select  left(FROM_UNIXTIME(time),10) as time, teacher_id  from teacher_accesses union all select left(FROM_UNIXTIME(time),10) as time, teacher_id   from  car_accesses) tb2 group by  tb2.time, tb2.teacher_id) tn group by tn.time"
	var teacherAttendCensus []response.AttendCensus
	if err = global.GVA_DB.Raw(sql).Find(&teacherAttendCensus).Error; err != nil {
		err = errors.New("获取教师考勤统计数量失败")
		return
	}

	// 获取考勤率准点率

	sql = `select left (FROM_UNIXTIME(time),10) as time, count(*) as on_time from (select time, teacher_id from ( select time, teacher_id, direction from teacher_accesses union all select time, teacher_id, direction from car_accesses) tb2 where right (FROM_UNIXTIME(tb2.time), 8) <= "09:00:00" and tb2.direction = 1 group by left(FROM_UNIXTIME(tb2.time), 10), tb2.teacher_id) tn group by left(FROM_UNIXTIME(tn.time), 10)`
	var teacherOnTimeCensus []response.AttendCensus
	if err = global.GVA_DB.Raw(sql).Find(&teacherOnTimeCensus).Error; err != nil {
		err = errors.New("获取教师准点统计数量失败")
		return
	}

	for i, _ := range teacherAttendCensus {
		teacherAttendCensus[i].Attend = utils.NumToFixed(teacherAttendCensus[i].Attend, float64(teacherTotal), "4") * 100
		teacherAttendCensus[i].OnTime = utils.NumToFixed(teacherOnTimeCensus[i].OnTime, float64(teacherTotal), "4") * 100
	}

	return err, teacherAttendCensus
}

// @Author: 张佳伟
// @Function: GetExamPassRate
// @Description: 合格率
// @Router: /dashboard/getExamPassRate
// @Date:2022/08/09 10:06:35

func GetExamPassRate() (err error, list interface{}) {

	selectSql := "grades.name as grade_name, count(*) as total"
	leftJoinSql1 := "left join exams on exams.id = exam_results.exam_id"
	leftJoinSql2 := "left join students on students.id = exam_results.student_id"
	leftJoinSql3 := "left join  grades  on  grades.id = students.grade_id"
	whereSql := "exams.grade_id = students.grade_id"
	var examPass []response.ExamPassRate
	if err = global.GVA_DB.Model(&model.ExamResult{}).Select(selectSql).Joins(leftJoinSql1).Joins(leftJoinSql2).Joins(leftJoinSql3).Where(whereSql).Group("students.grade_id").Find(&examPass).Error; err != nil {
		err = errors.New("获取学生参加考试的总数")
		return
	}

	whereSql += " and exam_results.result >= 60 "
	selectSql = "count(*) as rate"
	var examRate []response.ExamPassRate
	if err = global.GVA_DB.Model(&model.ExamResult{}).Select(selectSql).Joins(leftJoinSql1).Joins(leftJoinSql2).Joins(leftJoinSql3).Where(whereSql).Group("students.grade_id").Find(&examRate).Error; err != nil {
		err = errors.New("获取学生参加考试通过的数量")
		return
	}

	for i, _ := range examPass {
		examPass[i].Rate = examRate[i].Rate
	}
	return err, examPass
}

// @Author: 张佳伟
// @Function: GetStudentNum
// @Description: 获取学生数量
// @Router:/dashboard/getStudentNum
// @Date:2022/08/11 17:22:00

func GetStudentNum() (err error, list interface{}) {

	// select count(*) as total, grades.name , classes.name from students left join grades on grades.id = students.grade_id left join classes on classes.id = students.class_id group by sex, class_id
	selectSql := " count(*) as total, grades.name as grade_name , classes.name  class_name, sex"
	leftJoinSql1 := "left join grades on grades.id = students.grade_id"
	leftJoinSql2 := "left join classes on classes.id = students.class_id"
	var studentNum []response.StudentNum
	if err = global.GVA_DB.Model(&model.Student{}).Select(selectSql).Joins(leftJoinSql1).Joins(leftJoinSql2).Group("sex, class_id").Find(&studentNum).Error; err != nil {
		err = errors.New("获取学生人数")
		return
	}

	return err, studentNum
}

// @Author: 张佳伟
// @Function:GetTeacherAttendCensus
// @Description:获取教师考勤率，准点率
// @Router /dashboard/getStudentAttendCensus
// @Date:2022/08/09 17:55:29

func GetStudentAttendCensus() (err error, list interface{}) {

	//  获取学生总数
	var teacherTotal int64
	if err = global.GVA_DB.Model(&model.Student{}).Count(&teacherTotal).Error; err != nil {
		err = errors.New("获取教师总数失败")
		return
	}

	// 获取考勤率
	sql := "select left(FROM_UNIXTIME(time), 10) as time, count(*) as attend from ( select * from student_accesses group by left(FROM_UNIXTIME(time), 10), student_id ) tb group by left(FROM_UNIXTIME(time), 10)"
	var studentAttendCensus []response.AttendCensus
	if err = global.GVA_DB.Raw(sql).Find(&studentAttendCensus).Error; err != nil {
		err = errors.New("获取学生考勤统计数量失败")
		return
	}

	// 获取考勤率准点率
	sql = ` select left(FROM_UNIXTIME(time), 10) as time , count(*) as on_time from ( select * from student_accesses where right (FROM_UNIXTIME(time), 8) <= "09:00:00" and student_accesses.direction = 1 group by left(FROM_UNIXTIME(time), 10), student_id ) tb group by left(FROM_UNIXTIME(time), 10)`
	var teacherOnTimeCensus []response.AttendCensus
	if err = global.GVA_DB.Raw(sql).Find(&teacherOnTimeCensus).Error; err != nil {
		err = errors.New("获取学生准点统计数量失败")
		return
	}

	for i, _ := range studentAttendCensus {
		studentAttendCensus[i].Attend = utils.NumToFixed(studentAttendCensus[i].Attend, float64(teacherTotal), "4") * 100
		studentAttendCensus[i].OnTime = utils.NumToFixed(teacherOnTimeCensus[i].OnTime, float64(teacherTotal), "4") * 100
	}

	return err, studentAttendCensus
}
