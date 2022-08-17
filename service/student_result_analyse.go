package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
)

// @Author: 张佳伟
// @Function:GetStudentCourseResult
// @Description:获取学生每一个科目的成绩
// @Router:/studentResultAnalyse/getStudentCourseResult
// @Date:2022/08/16 17:44:32

func GetStudentCourseResult(info request.StudentResultAnalyse) (err error, list interface{}) {

	// 获取最后一次考试
	var lastExam model.Exam
	if err = global.GVA_DB.Model(&model.Exam{}).Where("grade_id = ?", info.GradeID).Last(&lastExam).Error; err != nil {
		err = errors.New("获取最后一条考试信息")
		return
	}

	// 获取学生每一科目成绩
	selectSql := "courses.name as course_name , exam_results.result "
	leftJoinSql1 := "left join students on students.id = exam_results.student_id"
	leftJoinSql2 := "left join courses on courses.id = exam_results.course_id "
	whereSql1 := "  exam_results.student_id = ? and exam_results.exam_id = ? "

	var studentTotalResult []response.StudentCourseResult
	db := global.GVA_DB.Model(&model.ExamResult{}).Select(selectSql).Joins(leftJoinSql1).Joins(leftJoinSql2).Where(whereSql1, info.StudentID, lastExam.ID)

	db.Find(&studentTotalResult)

	if err = db.Find(&studentTotalResult).Error; err != nil {
		err = errors.New("获取学生每一科成绩失败")
		return
	}

	return err, studentTotalResult
}

// @Author: 张佳伟
// @Function: GetStudentTotalResultHistory
// @Description: 获取学生历史考试每一次的总成绩
// @Router:/studentResultAnalyse/getStudentTotalResultHistory
// @Date:2022/08/17 16:22:07

func GetStudentTotalResultHistory(info request.StudentResultAnalyse) (err error, list interface{}) {

	selectSql := "exams.name as exam_name, sum(result) as total"
	leftJoinSql1 := "left join exams on exams.id = exam_results.exam_id"
	var studentExamTotalResultHistory []response.StudentExamTotalResultHistory
	db := global.GVA_DB.Model(&model.ExamResult{}).Select(selectSql).Joins(leftJoinSql1).Where("student_id = ? ", info.StudentID).Group("exams.id").Order("exams.id")

	if err = db.Find(&studentExamTotalResultHistory).Error; err != nil {
		err = errors.New("获取学生每一次考试总成绩失败")
		return
	}

	return err, studentExamTotalResultHistory
}

//  select
// 	exams.name as exam_name,
// 	sum(result) as total
// from
// 	exam_results
// left join exams on
// 	exams.id = exam_results.exam_id
// where
// 	student_id = 4
// 	and course_id = 1
