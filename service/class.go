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
// @Function: CreateClass
// @Description: 创建班级
// @Router: /class/createClass
// @Date: 2022/7/12 10:37:12

func CreateClass(class model.Class) (err error) {
	if !errors.Is(global.GVA_DB.Where("name = ? AND grade_id = ?", class.Name, class.GradeID).First(&model.Class{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同班级")
	}
	return global.GVA_DB.Create(&class).Error
}

// @Author: 张佳伟
// @Function: GetClassList
// @Description: 获取班级列表
// @Router: /class/getClassList
// @Date: 2022/7/12 11:50:34

func GetClassList(info request.SearchClassParams) (err error, list interface{}, total int64) {
	fmt.Println(info)
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.Class{})
	var classList []model.Class
	if info.Name != "" {
		db = db.Where("name = ?", info.Name)
	}
	if info.GradeID != 0 {
		db = db.Where("grade_id = ?", info.GradeID)
	}
	err = db.Count(&total).Error
	err = db.Debug().Limit(limit).Offset(offset).Preload("Teacher").Preload("Grade").Find(&classList).Error
	return err, classList, total
}

// @Author: 张佳伟
// @Function: UpClass
// @Description: 更新班级
// @Router: /class/upClass
// @Date: 2022/7/14 16:23

func UpClass(class model.Class) (err error) {
	err = global.GVA_DB.Where("id = ?", class.ID).First(&model.Class{}).Updates(&class).Error
	return err
}

// @Author: 张佳伟
// @Function: DeleteClass
// @Description: 删除班级
// @Router: /class/deleteClass
// @Date: 2022/7/14 16:52:51

func DeleteClass(class model.Class) (err error) {
	err = global.GVA_DB.Debug().Where("id = ?", class.ID).Delete(&class).Error
	return err
}
