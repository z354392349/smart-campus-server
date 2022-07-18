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
	if !errors.Is(global.GVA_DB.Where("name = ? ", exam.Name).First(&model.Exam{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同考试")
	}
	return global.GVA_DB.Create(&exam).Error
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
	var examList []model.Exam

	if info.Name != "" {
		db = db.Where("Name = ?", info.Name)
	}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("ExamItem").Find(&examList).Error
	return err, examList, total
}

// @Author: 张佳伟
// @Function:UpExam
// @Description: 更新考试
// @Router:/exam/upExam
// @Date:2022/07/16 20:17:14

func UpExam(exam model.Exam) (err error) {
	examItems := exam.ExamItem
	err = global.GVA_DB.Where("id = ?", exam.ID).First(&model.Exam{}).Updates(&exam).Error
	for _, examItem := range examItems {
		err = global.GVA_DB.Model(&model.ExamItem{}).Where("exam_id = ? AND  course_id = ? ", exam.ID, examItem.CourseID).Updates(&examItem).Error
	}
	return err
}

// @Author: 张佳伟
// @Function:DeleteExam
// @Description:删除考试
// @Router:/exam/deleteExam
// @Date:2022/07/16 20:17:47

func DeleteExam(exam model.Exam) (err error) {
	err = global.GVA_DB.Debug().Where("id = ?", exam.ID).Delete(&exam).Error
	err = global.GVA_DB.Model(&model.ExamItem{}).Where("exam_id = ? AND  course_id = ? ", exam.ID).Error
	return err
}
