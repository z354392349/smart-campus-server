package service

import (
	"errors"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"

	"gorm.io/gorm"
)

// @Author: 张佳伟
// @Function:CreateExam
// @Description:创建考试
// @Router:/exam/createExam
// @Date:2022/07/16 20:06:51

func CreateExam(exam model.Exam) (err error) {

	global.GVA_DB.Transaction(func(db *gorm.DB) error {

		// 返回 nil 提交事务

		if !errors.Is(global.GVA_DB.Where("name = ? ", exam.Name).First(&model.Exam{}).Error, gorm.ErrRecordNotFound) {
			return errors.New("存在相同考试")
		}

		// TODO: 插入学生成绩
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
					ExamID:     exam.ID,
					ExamItemID: k.ID,
					StudentID:  v.ID,
				}
				examResultList = append(examResultList, examResult)
			}
		}

		// TODO:examResultList 插入数据库
		// examResultList
		// return global.GVA_DB.Create(&exam).Error

		return nil

	})
	return
}

// @Author: 张佳伟
// @Function:GetExamList
// @Description:获取考试列表
// @Router:/exam/getExamList
// @Date:2022/07/16 20:10:37

func GetExamList(info request.SearchExamParams) (err error, list interface{}, total int64) {
	fmt.Println(info)
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&model.Exam{})
	db1 := global.GVA_DB.Model(&model.ExamItem{})
	var examList []model.Exam
	// var examItem []model.ExamItem

	if info.Name != "" {
		db = db.Where("Name = ?", info.Name)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("Grade").Find(&examList).Error

	for i, _ := range examList {

		leftJoinSql := "left join courses on exam_items.course_id = courses.id"
		selectSql := "exam_items.*, courses.name as course_name "
		err = db1.Debug().Select(selectSql).Where("exam_id = ?", examList[i].ID).Joins(leftJoinSql).Find(&examList[i].ExamItem).Error
	}

	return err, examList, total
}

// @Author: 张佳伟
// @Function:UpExam
// @Description: 更新考试
// @Router:/exam/upExam
// @Date:2022/07/16 20:17:14

func UpExam(exam model.Exam) (err error) {
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

// @Author: 张佳伟
// @Function:DeleteExam
// @Description:删除考试
// @Router:/exam/deleteExam
// @Date:2022/07/16 20:17:47

func DeleteExam(exam model.Exam) (err error) {

	if err = global.GVA_DB.Debug().Where("id = ?", exam.ID).Delete(&exam).Error; err != nil {
		return err
	}

	if err = global.GVA_DB.Where("exam_id = ? ", exam.ID).Delete(&model.ExamItem{}).Error; err != nil {
		return err
	}

	if err = global.GVA_DB.Where("exam_id = ? ", exam.ID).Delete(&model.AllotExamRoom{}).Error; err != nil {
		return err
	}

	// TODO: 删除考试成绩表
	return err

}
