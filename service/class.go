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
	if !errors.Is(global.GVA_DB.Where("name = ? ", class.Name).First(&model.Class{}).Error, gorm.ErrRecordNotFound) {
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
		db = db.Where("Name = ?", info.Name)
	}
	err = db.Count(&total).Error
	err = db.Debug().Limit(limit).Offset(offset).Preload("Teacher").Preload("Grade").Find(&classList).Error
	return err, classList, total
}

func UpCreate1(grade model.Grade) (err error) {
	err = global.GVA_DB.Where("id = ?", grade.ID).First(&model.Grade{}).Updates(&grade).Error
	return err
}

func DeleteGrade1(grade model.Grade) (err error) {
	err = global.GVA_DB.Debug().Where("id = ?", grade.ID).Delete(&grade).Error
	return err
}
