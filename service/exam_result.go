package service

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"

	"github.com/fatih/structs"
)

// @Author: 张佳伟
// @Function:GetExamResultList
// @Description:获取学生成绩列表
// @Router:/examResult/GetExamResultList
// @Date:2022/07/27 17:58:39

func GetExamResultList(info request.SearchExamResultParams) (err error, list interface{}, total int64) {
	fmt.Println(info)
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&model.ExamResult{})
	var examResultList []response.ExamResult

	leftJoinSql1 := "left join students on students.id = exam_results.student_id" // 学生姓名
	leftJoinSql2 := "left join exams on exams.id = exam_results.exam_id"          // 考试名称
	leftJoinSql3 := "left join grades on grades.id = students.grade_id"           // 年级名称
	leftJoinSql4 := "left join classes on classes.id = students.class_id"         // 班级名称
	leftJoinSql5 := "left join courses on courses.id = exam_results.course_id"    // 科目名称

	selectSql := "exam_results.*, students.name as student_name, exams.name as exam_name, grades.name as grade_name, classes.name as class_name, courses.name as course_name"

	db = db.Limit(limit).Offset(offset).Select(selectSql).Joins(leftJoinSql1).Joins(leftJoinSql2).Joins(leftJoinSql3).Joins(leftJoinSql4).Joins(leftJoinSql5)

	if info.Name != "" {
		db = db.Where("students.name LIKE ?", "%"+info.Name+"%")
	}
	if info.ExamID != 0 {
		db = db.Where("exams.id = ?", info.ExamID)
	}

	if info.GradeID != 0 {
		db = db.Where("grades.id  = ?", info.GradeID)
	}

	if info.ClassID != 0 {
		db = db.Where("classes.id  = ?", info.ClassID)
	}

	if err = db.Limit(limit).Offset(offset).Find(&examResultList).Error; err != nil {
		return
	}

	if err = db.Limit(-1).Offset(-1).Count(&total).Error; err != nil {
		return
	}
	return err, examResultList, total
}

// @Author: 张佳伟
// @Function:UpExamResult
// @Description:更新学生成绩
// @Router:/examResult/upExamResult/
// @Date:2022/07/28 17:37:05

func UpExamResult(info request.SetExamResultParams) (err error) {
	infoMap := structs.Map(info)
	// err = global.GVA_DB.Where("id = ?", info.ID).Updates(&model.ExamResult{Description: info.Description, Result: *info.Result}).Error
	err = global.GVA_DB.Model(&model.ExamResult{}).Where("id = ?", info.ID).Updates(infoMap).Error
	return err
}
