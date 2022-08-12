package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/utils"
)

// type classStudentNum struct {
// 	ExamID    uint  `json:"examID" form:"examID" gorm:"comment:考试ID;"`
// 	Result    *int  `json:"result" form:"result"`
// 	CourseID  uint  `json:"courseID" form:"courseID" gorm:"comment:科目ID;" `
// 	StudentID uint  `json:"studentID" form:"studentID" gorm:"comment:学生ID;" `
// 	Num       int64 `json:"num"`
// }

func GetDashboardCensusNum1() (err error, carAccess int64, peopleAccess int64, teacherCensus response.TeacherCensus, studentCensus response.StudentCensus) {

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

func GetTeacherNum1() (err error, list interface{}) {
	var teacherNum []response.TeacherNum
	if err = global.GVA_DB.Model(&model.Teacher{}).Select("sex, count(1) as num ").Group("sex").Find(&teacherNum).Error; err != nil {
		err = errors.New("获取教师性别统计数量失败")
		return
	}

	return err, teacherNum
}

func GetStudentNum1() (err error, list interface{}) {

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

func GetStudentAttendCensus1() (err error, list interface{}) {

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

// @Author: 张佳伟
// @Function:GetGradeAverageResult
// @Description: 获取平均学习成绩
// @Router:/gradeResultAnalyse/getGradeAverageResult
// @Date:2022/08/08 21:02:59

func GetGradeAverageResult(info request.GradeAverageResult) (err error, list interface{}) {

	// 获取最后一次考试
	var lastExam model.Exam
	if err = global.GVA_DB.Model(&model.Exam{}).Where("grade_id = ?", info.GradeID).Last(&lastExam).Error; err != nil {
		err = errors.New("获取最后一条考试信息")
		return
	}

	// 获取每个班级学生数量
	var classNum []float64
	if err = global.GVA_DB.Debug().Model(&model.Student{}).Select("count(*) as total").Where("grade_id = ?", info.GradeID).Group("class_id").Order("class_id").Find(&classNum).Error; err != nil {
		err = errors.New("获取每个班有多少学生")
		return
	}

	// 获取学生成绩汇总
	selectSql := "exam_results.*, classes.name as class_name , sum(`result`) as num"
	leftJoinSql1 := "exam_results left join students on students.id = exam_results.student_id"
	leftJoinSql2 := "left join classes on classes.id = students.class_id"
	var averageResult []response.GradeAverageResult
	db := global.GVA_DB.Model(&model.ExamResult{}).Select(selectSql).Joins(leftJoinSql1).Joins(leftJoinSql2).Where("exam_results.exam_id  = ?", lastExam.ID).Group("students.class_id").Order("students.class_id")

	if info.GradeID != 0 {
		db = db.Where("students.grade_id = ?", info.GradeID)
	}
	if info.CourseID != 0 {
		db = db.Where("exam_results.course_id  = ?", info.GradeID)
	}

	db.Find(&averageResult)

	if err = db.Find(&averageResult).Error; err != nil {
		err = errors.New("获取年级学生成绩失败")
		return
	}

	for i, _ := range averageResult {
		averageResult[i].Num = utils.NumToFixed(averageResult[i].Num, classNum[i], "2")
	}
	return err, averageResult
}

// @Author: 张佳伟
// @Function:GetGradePassPercent
// @Description:获取年级考试通过率
// @Router:/gradeResultAnalyse/GetGradePassPercent
// @Date:2022/08/12 15:40:40

func GetGradePassPercent(info request.GradePassPercent) (err error, percent interface{}) {

	// 获取最后一次考试
	var lastExam model.Exam
	if err = global.GVA_DB.Model(&model.Exam{}).Where("grade_id = ?", info.GradeID).Last(&lastExam).Error; err != nil {
		err = errors.New("获取最后一条考试信息")
		return
	}

	// 参加这次考试的学生数
	var attendStudentNum float64
	if err = global.GVA_DB.Debug().Model(&model.ExamResult{}).Select("count(*) as total").Where("exam_results.exam_id  = ?", lastExam.ID).Find(&attendStudentNum).Error; err != nil {
		err = errors.New("参加这次考试的学生数")
		return
	}

	// 通过这次考试的学生数
	var pasStudentNum float64
	if err = global.GVA_DB.Debug().Model(&model.ExamResult{}).Select("count(*) as total").Where("exam_results.exam_id  = ? AND result >= ?", lastExam.ID, 60).Find(&pasStudentNum).Error; err != nil {
		err = errors.New("参加这次考试的学生数")
		return
	}

	percent = utils.NumToFixed(pasStudentNum, attendStudentNum, "4") * 100

	return err, percent
}
