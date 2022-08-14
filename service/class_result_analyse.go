package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
	"gin-vue-admin/utils"
)

// @Author: 张佳伟
// @Function:GetStudentTotalResult
// @Description: 获取班级下每个学生总成绩
// @Router:/studentResultAnalyse/getStudentTotalResult
// @Date:2022/08/14 18:05:44

func GetStudentTotalResult(info request.ClassResultAnalyse) (err error, list interface{}) {

	// 获取最后一次考试
	var lastExam model.Exam
	if err = global.GVA_DB.Model(&model.Exam{}).Where("grade_id = ?", info.GradeID).Last(&lastExam).Error; err != nil {
		err = errors.New("获取最后一条考试信息")
		return
	}

	// 获取学生成绩汇总
	selectSql := "students.name as student_name, sum(exam_results.result) as total "
	leftJoinSql1 := "left join students on students.id = exam_results.student_id"
	whereSql1 := "students.class_id = ? and exam_id = ? "
	var studentTotalResult []response.StudentTotalResult
	db := global.GVA_DB.Model(&model.ExamResult{}).Select(selectSql).Joins(leftJoinSql1).Where(whereSql1, info.ClassID, lastExam.ID).Group("students.id").Order("total desc")

	if info.CourseID != 0 {
		db = db.Where("exam_results.course_id  = ?", info.CourseID)
	}

	db.Find(&studentTotalResult)

	if err = db.Find(&studentTotalResult).Error; err != nil {
		err = errors.New("获取年级学生成绩失败")
		return
	}

	return err, studentTotalResult
}

func GetGradePassPercent1(info request.GradeResultAnalyse) (err error, percent interface{}) {

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

func GetGradeAverageResultHistory1(info request.GradeResultAnalyse) (err error, list interface{}) {

	// 参加这次考试的学生数
	selectSql := " exams.name as  exam_name, classes.name as class_name, FORMAT(sum(`result`) / COUNT(distinct students.id), 2) as result"
	leftJoinSql1 := "left join students on students.id = exam_results.student_id"
	leftJoinSql3 := "left join classes on classes.id = students.class_id "
	leftJoinSql2 := "left join exams on exams.id = exam_results.exam_id"

	var gradeAverageResultHistory []response.GradeAverageResultHistory
	db := global.GVA_DB.Model(&model.ExamResult{}).Select(selectSql).Joins(leftJoinSql1).Joins(leftJoinSql2).Joins(leftJoinSql3).Where("students.grade_id = ? and exams.grade_id = students.grade_id", info.GradeID).Group("exam_id, class_id").Order("exams.id")

	if err = db.Find(&gradeAverageResultHistory).Error; err != nil {
		err = errors.New("参加这次考试的学生数")
		return
	}

	return err, gradeAverageResultHistory
}

func GetGradeCourseAverageResultHistory1(info request.GradeResultAnalyse) (err error, list interface{}) {

	// 参加这次考试的学生数
	selectSql := " exams.name as  exam_name, classes.name as class_name, FORMAT(sum(`result`) / COUNT(distinct students.id), 2) as result"
	leftJoinSql1 := "left join students on students.id = exam_results.student_id"
	leftJoinSql3 := "left join classes on classes.id = students.class_id "
	leftJoinSql2 := "left join exams on exams.id = exam_results.exam_id"
	whereSql1 := "students.grade_id = ? and exams.grade_id = students.grade_id and exam_results.course_id = ?"
	var gradeAverageResultHistory []response.GradeAverageResultHistory
	db := global.GVA_DB.Model(&model.ExamResult{}).Select(selectSql).Joins(leftJoinSql1).Joins(leftJoinSql2).Joins(leftJoinSql3).Where(whereSql1, info.GradeID, info.CourseID).Group("exam_id, class_id, course_id").Order("exams.id")

	if err = db.Find(&gradeAverageResultHistory).Error; err != nil {
		err = errors.New("参加这次考试的学生数")
		return
	}

	return err, gradeAverageResultHistory
}
