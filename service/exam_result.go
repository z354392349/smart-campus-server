package service

import (
	"errors"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"

	"gorm.io/gorm"
)

func CreateExam1(exam model.Exam) (err error) {

	global.GVA_DB.Transaction(func(db *gorm.DB) error {

		if !errors.Is(global.GVA_DB.Where("name = ? ", exam.Name).First(&model.Exam{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同考试")
		}

		// 1.查询有年级有多少学生，
		// 2.每个学生分别对应发布的科目，
		// 3.整理出数据，插入 ExamResult
		var studentList []model.Student
		var examItemList = exam.ExamItem
		var examResultList []model.ExamResult

		if err = global.GVA_DB.Debug().Model(&model.Student{}).Where("grade_id = ?", exam.GradeID).Find(&studentList).Error; err != nil {
			return err
		}

		for _, v := range studentList {
			for _, k := range examItemList {
				examResult := model.ExamResult{
					ExamID:    exam.ID,
					StudentID: v.ID,
					CourseID:  k.CourseID,
				}
				examResultList = append(examResultList, examResult)
			}
		}

		if err = global.GVA_DB.Create(&examResultList).Error; err != nil {
			return err
		}

		return nil

	})
	return
}

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

	if info.GradeID != 0 {
		db = db.Where("grades.id  = ?", info.GradeID)
	}

	if info.ClassID != 0 {
		db = db.Where("classes.id  = ?", info.ClassID)
	}

	if err = db.Limit(limit).Offset(offset).Find(&examResultList).Error; err != nil {
		return
	}

	if err = db.Count(&total).Error; err != nil {
		return
	}
	return err, examResultList, total
}

// @Author: 张佳伟
// @Function:UpExam
// @Description: 更新考试
// @Router:/exam/upExam
// @Date:2022/07/16 20:17:14

func UpExam1(exam model.Exam) (err error) {
	examItems := exam.ExamItem
	examItemsUp := []model.ExamItem{}
	examItemsCreate := []model.ExamItem{}
	courseIDs := []uint{}
	for _, v := range examItems {
		if v.ID != 0 {
			examItemsUp = append(examItemsUp, v)
			courseIDs = append(courseIDs, v.CourseID)
		} else {
			v.ExamID = exam.ID
			examItemsCreate = append(examItemsCreate, v)
		}
	}

	// 更新考试数据
	if err = global.GVA_DB.Where("id = ?", exam.ID).Debug().First(&model.Exam{}).Updates(&exam).Error; err != nil {
		return
	}

	//删除考试项旧数据

	if err = global.GVA_DB.Model(&model.ExamItem{}).Debug().Where("exam_id = ? AND course_id not in ? ", exam.ID, courseIDs).Delete(model.ExamItem{}).Error; err != nil {
		return
	}

	//更新考试项已有数据
	if len(examItemsUp) != 0 {
		for _, v := range examItemsUp {
			if err = global.GVA_DB.Model(&model.ExamItem{}).Debug().Where("id = ?", v.ID).Updates(&v).Error; err != nil {
				return err
			}
		}
	}

	// 创建考试项新数据
	if len(examItemsCreate) != 0 {
		err = global.GVA_DB.Model(&model.ExamItem{}).Debug().Create(&examItemsCreate).Error
	}

	return err
}

func DeleteExam1(exam model.Exam) (err error) {

	if err = global.GVA_DB.Debug().Where("id = ?", exam.ID).Delete(&exam).Error; err != nil {
		return err
	}

	if err = global.GVA_DB.Where("exam_id = ? ", exam.ID).Delete(&model.ExamItem{}).Error; err != nil {
		return err
	}

	if err = global.GVA_DB.Where("exam_id = ? ", exam.ID).Delete(&model.AllotExamRoom{}).Error; err != nil {
		return err
	}

	return err

}
