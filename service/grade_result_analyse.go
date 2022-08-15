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
// @Function:GetGradeAverageResult
// @Description: 获取平均学习成绩
// @Router:/gradeResultAnalyse/getGradeAverageResult
// @Date:2022/08/08 21:02:59

func GetGradeAverageResult(info request.GradeResultAnalyse) (err error, list interface{}) {

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
	leftJoinSql1 := "left join students on students.id = exam_results.student_id"
	leftJoinSql2 := "left join classes on classes.id = students.class_id"
	var averageResult []response.GradeAverageResult
	db := global.GVA_DB.Model(&model.ExamResult{}).Select(selectSql).Joins(leftJoinSql1).Joins(leftJoinSql2).Where("exam_results.exam_id  = ?", lastExam.ID).Group("students.class_id").Order("students.class_id")

	if info.GradeID != 0 {
		db = db.Where("students.grade_id = ?", info.GradeID)
	}
	if info.CourseID != 0 {
		db = db.Where("exam_results.course_id  = ?", info.CourseID)
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

func GetGradePassPercent(info request.GradeResultAnalyse) (err error, percent interface{}) {

	// 获取最后一次考试
	var lastExam model.Exam
	if err = global.GVA_DB.Model(&model.Exam{}).Where("grade_id = ?", info.GradeID).Last(&lastExam).Error; err != nil {
		err = errors.New("获取最后一条考试信息")
		return
	}

	// 参加这次考试的学生数
	var attendStudentNum float64
	if err = global.GVA_DB.Debug().Model(&model.ExamResult{}).Select("count(*) as total").Where("exam_results.exam_id  = ? ", lastExam.ID).Find(&attendStudentNum).Error; err != nil {
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

// @Author: 张佳伟
// @Function:GetGradeAverageResultHistory
// @Description: 获取年级平均成绩
// @Router:/gradeResultAnalyse/GetGradeAverageResultHistory
// @Date:2022/08/13 17:13:08

func GetGradeAverageResultHistory(info request.GradeResultAnalyse) (err error, list interface{}) {

	// 参加这次考试的学生数
	selectSql := " exams.name as  exam_name, classes.name as class_name, FORMAT(sum(`result`) / COUNT(distinct students.id), 2) as result"
	leftJoinSql1 := "left join students on students.id = exam_results.student_id"
	leftJoinSql3 := "left join classes on classes.id = students.class_id "
	leftJoinSql2 := "left join exams on exams.id = exam_results.exam_id"

	var gradeAverageResultHistory []response.GradeAverageResultHistory
	db := global.GVA_DB.Model(&model.ExamResult{}).Select(selectSql).Joins(leftJoinSql1).Joins(leftJoinSql2).Joins(leftJoinSql3).Where("students.grade_id = ? and exams.grade_id = students.grade_id", info.GradeID).Group("exam_id, class_id").Order("exams.id")

	if err = db.Find(&gradeAverageResultHistory).Error; err != nil {
		err = errors.New("获取年级平均成绩失败")
		return
	}

	return err, gradeAverageResultHistory
}

// @Author: 张佳伟
// @Function:GetGradeCourseAverageResultHistory
// @Description: 获取年级科目平均成绩
// @Router:/gradeResultAnalyse/GetGradeCourseAverageResultHistory
// @Date:2022/08/13 17:13:08

func GetGradeCourseAverageResultHistory(info request.GradeResultAnalyse) (err error, list interface{}) {

	// 参加这次考试的学生数
	selectSql := " exams.name as  exam_name, classes.name as class_name, FORMAT(sum(`result`) / COUNT(distinct students.id), 2) as result"
	leftJoinSql1 := "left join students on students.id = exam_results.student_id"
	leftJoinSql3 := "left join classes on classes.id = students.class_id "
	leftJoinSql2 := "left join exams on exams.id = exam_results.exam_id"
	whereSql1 := "students.grade_id = ? and exams.grade_id = students.grade_id and exam_results.course_id = ?"
	var gradeAverageResultHistory []response.GradeAverageResultHistory
	db := global.GVA_DB.Model(&model.ExamResult{}).Select(selectSql).Joins(leftJoinSql1).Joins(leftJoinSql2).Joins(leftJoinSql3).Where(whereSql1, info.GradeID, info.CourseID).Group("exam_id, class_id, course_id").Order("exams.id")

	if err = db.Find(&gradeAverageResultHistory).Error; err != nil {
		err = errors.New("获取年级科目平均成绩失败")
		return
	}

	return err, gradeAverageResultHistory
}
