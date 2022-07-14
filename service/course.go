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
// @Function: CreateCourse
// @Description: 创建课程列表
// @Router: /course/createCourse
// @Date:2022/07/14 22:00:51

func CreateCourse(course model.Course) (err error) {
	if !errors.Is(global.GVA_DB.Where("name = ? ", course.Name).First(&model.Course{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("获取课程列表")
	}
	return global.GVA_DB.Create(&course).Error
}

// @Author: 张佳伟
// @Function: CreateCourse
// @Description: 获取课程列表
// @Router: /course/createCourse
// @Date:2022/07/14 22:02:34

func GetCourseList(info request.SearchCourseParams) (err error, list interface{}, total int64) {
	fmt.Println(info)
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.Course{})
	var courseList []model.Course

	if info.Name != "" {
		db = db.Where("Name = ?", info.Name)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Debug().Limit(limit).Offset(offset).Find(&courseList).Error
	return err, courseList, total
}

// @Author: 张佳伟
// @Function: UpCourse
// @Description: 更新课程
// @Router: /course/upCourse
// @Date:2022/07/14 22:05:30

func UpCourse(course model.Course) (err error) {
	err = global.GVA_DB.Where("id = ?", course.ID).First(&model.Course{}).Updates(&course).Error
	return err
}

// @Author: 张佳伟
// @Function:DeleteCourse
// @Description:删除课程
// @Router:/course/deleteCourse
// @Date:2022/07/14 22:06:56

func DeleteCourse(course model.Course) (err error) {
	err = global.GVA_DB.Debug().Where("id = ?", course.ID).Delete(&course).Error
	return err
}
