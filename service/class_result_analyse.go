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
// @Function:GetClassTotalResult
// @Description: 获取班级下每个学生总成绩
// @Router:/classResultAnalyse/getClassTotalResult
// @Date:2022/08/14 18:05:44

func GetClassTotalResult(info request.ClassResultAnalyse) (err error, list interface{}) {

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
	var classTotalResult []response.ClassTotalResult
	db := global.GVA_DB.Model(&model.ExamResult{}).Select(selectSql).Joins(leftJoinSql1).Where(whereSql1, info.ClassID, lastExam.ID).Group("students.id").Order("total desc")

	if info.CourseID != 0 {
		db = db.Where("exam_results.course_id  = ?", info.CourseID)
	}

	db.Find(&classTotalResult)

	if err = db.Find(&classTotalResult).Error; err != nil {
		err = errors.New("获取年级学生成绩失败")
		return
	}

	return err, classTotalResult
}

// @Author: 张佳伟
// @Function:GetClassPassPercent
// @Description: 获取指定班级的通过率
// @Router:/classResultAnalyse/getClassPassPercent
// @Date:2022/08/15 09:16:41

func GetClassPassPercent(info request.ClassResultAnalyse) (err error, percent interface{}) {

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
	var pasClassNum float64
	if err = global.GVA_DB.Debug().Model(&model.ExamResult{}).Select("count(*) as total").Joins(leftJoinSql1).Where("exam_results.exam_id  = ?  and students.class_id = ? and result >= ?", lastExam.ID, info.ClassID, 60).Find(&pasClassNum).Error; err != nil {
		err = errors.New("参加这次考试通过的学生数")
		return
	}

	percent = utils.NumToFixed(pasClassNum, attendStudentNum, "4") * 100

	return err, percent
}

// @Author: 张佳伟
// @Function:GetClassToTalResultHistory
// @Description:获取班级下每一个学生，历史考试总成绩成绩
// @Router:/classResultAnalyse/getClassToTalResultListHistory
// @Date:2022/08/15 09:52:31

func GetClassToTalResultHistory(info request.ClassResultAnalyse) (err error, list interface{}) {

	selectSql := "exams.name as exam_name, students.name as student_name, students.id as student_id ,sum(`result`) as result"
	leftJoinSql1 := "left join students on students.id = exam_results.student_id"
	leftJoinSql3 := "left join classes on classes.id = students.class_id "
	leftJoinSql2 := "left join exams on exams.id = exam_results.exam_id"
	whereSql1 := "exams.grade_id = students.grade_id and students.class_id = ?"
	var classTotalResultHistory []response.ClassTotalResultHistory
	db := global.GVA_DB.Model(&model.ExamResult{}).Select(selectSql).Joins(leftJoinSql1).Joins(leftJoinSql2).Joins(leftJoinSql3).Where(whereSql1, info.ClassID).Group("exam_id, class_id, students.id ").Order("exams.name,	student_id ")

	if err = db.Find(&classTotalResultHistory).Error; err != nil {
		err = errors.New("获取班级学生历史考试总成绩失败")
		return
	}

	return err, classTotalResultHistory
}

// @Author: 张佳伟
// @Function:GetClassCourseResultHistory1
// @Description:获取班级下每一个学生，历史考试总成绩成绩
// @Router:/classResultAnalyse/getClassCourseResultHistory1
// @Date:2022/08/15 09:52:31

func GetClassCourseResultHistory(info request.ClassResultAnalyse) (err error, list interface{}) {

	selectSql := "exams.name as exam_name, students.name as student_name, students.id as student_id ,sum(`result`) as result"
	leftJoinSql1 := "left join students on students.id = exam_results.student_id"
	leftJoinSql3 := "left join classes on classes.id = students.class_id "
	leftJoinSql2 := "left join exams on exams.id = exam_results.exam_id"
	whereSql1 := "exams.grade_id = students.grade_id and students.class_id = ? and exam_results.course_id = ?"
	var gradeAverageResultHistory []response.ClassTotalResultHistory
	db := global.GVA_DB.Model(&model.ExamResult{}).Select(selectSql).Joins(leftJoinSql1).Joins(leftJoinSql2).Joins(leftJoinSql3).Where(whereSql1, info.ClassID, info.GradeID).Group("exam_id, class_id, students.id ").Order("exams.name,	student_id ")

	if err = db.Find(&gradeAverageResultHistory).Error; err != nil {
		err = errors.New("获取班级学生历史考试单科成绩失败")
		return
	}

	return err, gradeAverageResultHistory
}
