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
// @Function:GetStudentCourseResultList
// @Description:获取学生每一个科目的成绩
// @Router:
// @Date:2022/08/16 17:44:32

func GetStudentCourseResultList(info request.StudentResultAnalyse) (err error, list interface{}) {

	// 获取最后一次考试
	var lastExam model.Exam
	if err = global.GVA_DB.Model(&model.Exam{}).Where("grade_id = ?", info.GradeID).Last(&lastExam).Error; err != nil {
		err = errors.New("获取最后一条考试信息")
		return
	}

	// select courses.name as course_name , exam_results.`result` from exam_results left join students on students.id = exam_results.student_id left join courses on courses.id = exam_results.course_id where exam_results.exam_id = 12 and exam_results.student_id = 324

	// 获取学生成绩汇总
	selectSql := "courses.name as course_name , exam_results.result "
	leftJoinSql1 := "left join students on students.id = exam_results.student_id"
	leftJoinSql2 := "left join courses on courses.id = exam_results.course_id "
	whereSql1 := "  exam_results.student_id = ? and exam_results.exam_id = ? "

	var studentTotalResult []response.StudentCourseResult
	db := global.GVA_DB.Model(&model.ExamResult{}).Select(selectSql).Joins(leftJoinSql1).Joins(leftJoinSql2).Where(whereSql1, info.ClassID, lastExam.ID)

	db.Find(&studentTotalResult)

	if err = db.Find(&studentTotalResult).Error; err != nil {
		err = errors.New("获取年级学生每一科成绩失败")
		return
	}

	return err, studentTotalResult
}

// @Author: 张佳伟
// @Function:GetClassPassPercent
// @Description: 获取指定班级的通过率
// @Router:/studentResultAnalyse/getClassPassPercent
// @Date:2022/08/15 09:16:41

func GetClassPassPercent1(info request.ClassResultAnalyse) (err error, percent interface{}) {

	// 获取最后一次考试
	var lastExam model.Exam
	if err = global.GVA_DB.Model(&model.Exam{}).Where("grade_id = ?", info.GradeID).Last(&lastExam).Error; err != nil {
		err = errors.New("获取最后一条考试信息")
		return
	}

	// 参加这次考试的学生数
	var attendStudentNum float64
	leftJoinSql1 := "left join students on students.id = exam_results.student_id"
	if err = global.GVA_DB.Debug().Model(&model.ExamResult{}).Select("count(*) as total").Joins(leftJoinSql1).Where("exam_results.exam_id  = ? and students.class_id = ? ", lastExam.ID, info.ClassID).Find(&attendStudentNum).Error; err != nil {
		err = errors.New("参加这次考试的学生数")
		return
	}

	// 通过这次考试的学生数
	var pasStudentNum float64
	if err = global.GVA_DB.Debug().Model(&model.ExamResult{}).Select("count(*) as total").Joins(leftJoinSql1).Where("exam_results.exam_id  = ?  and students.class_id = ? and result >= ?", lastExam.ID, info.ClassID, 60).Find(&pasStudentNum).Error; err != nil {
		err = errors.New("参加这次考试通过的学生数")
		return
	}

	percent = utils.NumToFixed(pasStudentNum, attendStudentNum, "4") * 100

	return err, percent
}

// @Author: 张佳伟
// @Function:GetStudentToTalResultHistory
// @Description:获取班级下每一个学生，历史考试总成绩成绩
// @Router:/studentResultAnalyse/getStudentToTalResultListHistory
// @Date:2022/08/15 09:52:31

func GetStudentToTalResultHistory1(info request.ClassResultAnalyse) (err error, list interface{}) {

	selectSql := "exams.name as exam_name, students.name as student_name, students.id as student_id ,sum(`result`) as result"
	leftJoinSql1 := "left join students on students.id = exam_results.student_id"
	leftJoinSql3 := "left join classes on classes.id = students.class_id "
	leftJoinSql2 := "left join exams on exams.id = exam_results.exam_id"
	whereSql1 := "exams.grade_id = students.grade_id and students.class_id = ?"
	var gradeAverageResultHistory []response.StudentTotalResultHistory
	db := global.GVA_DB.Model(&model.ExamResult{}).Select(selectSql).Joins(leftJoinSql1).Joins(leftJoinSql2).Joins(leftJoinSql3).Where(whereSql1, info.ClassID).Group("exam_id, class_id, students.id ").Order("exams.name,	student_id ")

	if err = db.Find(&gradeAverageResultHistory).Error; err != nil {
		err = errors.New("获取班级学生历史考试总成绩失败")
		return
	}

	return err, gradeAverageResultHistory
}

// @Author: 张佳伟
// @Function:GetStudentCourseResultHistory1
// @Description:获取班级下每一个学生，历史考试总成绩成绩
// @Router:/studentResultAnalyse/getStudentCourseResultHistory1
// @Date:2022/08/15 09:52:31

func GetStudentCourseResultHistory1(info request.ClassResultAnalyse) (err error, list interface{}) {

	selectSql := "exams.name as exam_name, students.name as student_name, students.id as student_id ,sum(`result`) as result"
	leftJoinSql1 := "left join students on students.id = exam_results.student_id"
	leftJoinSql3 := "left join classes on classes.id = students.class_id "
	leftJoinSql2 := "left join exams on exams.id = exam_results.exam_id"
	whereSql1 := "exams.grade_id = students.grade_id and students.class_id = ? and exam_results.course_id = ?"
	var gradeAverageResultHistory []response.StudentTotalResultHistory
	db := global.GVA_DB.Model(&model.ExamResult{}).Select(selectSql).Joins(leftJoinSql1).Joins(leftJoinSql2).Joins(leftJoinSql3).Where(whereSql1, info.ClassID, info.GradeID).Group("exam_id, class_id, students.id ").Order("exams.name,	student_id ")

	if err = db.Find(&gradeAverageResultHistory).Error; err != nil {
		err = errors.New("获取班级学生历史考试单科成绩失败")
		return
	}

	return err, gradeAverageResultHistory
}
