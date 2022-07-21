package service

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @Author: 张佳伟
// @Function: CreateTeacher
// @Description: 创建教师
// @Router: /teacher/createTeacher
// @Date:2022/07/10 19:43:52

func CreateTeacher(teacher model.Teacher) (err error) {
	return global.GVA_DB.Debug().Create(&teacher).Error
}

// @Author: 张佳伟
// @Function: GetTeacherList
// @Description: 获取教师列表
// @Router: /teacher/GetTeacherList
// @Date: 2022/7/11 9:45:52

func GetTeacherList(info request.SearchTeacherParams) (err error, list interface{}, total int64) {
	fmt.Println(info)
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.Teacher{})
	var teacherList []model.Teacher

	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	err = db.Debug().Limit(limit).Offset(offset).Preload("Course").Find(&teacherList).Error
	return err, teacherList, total
}

// @Author: 张佳伟
// @Function: UpTeacher
// @Description: 更新教师
// @Router:  /teacher/upTeacher
// @Date: 2022/7/11 16:37

func UpTeacher(teacher model.Teacher) (err error) {
	err = global.GVA_DB.Debug().Where("id = ?", teacher.ID).First(&model.Teacher{}).Updates(&teacher).Error
	return err
}

// @Author: 张佳伟
// @Function:
// @Description:
// @Router:
// @Date:2022/07/09 10:32:22

func DeleteTeacher(teacher model.Teacher) (err error) {
	err = global.GVA_DB.Debug().Where("id = ?", teacher.ID).Delete(&teacher).Error
	return err
}
