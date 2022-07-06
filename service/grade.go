package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"

	"gorm.io/gorm"
)

//@author: 张佳伟
//@function: CreateApi
//@description: 新增基础api
//@param: api model.SysApi
//@return: err error

func CreateGrade(grade model.Grade) (err error) {
	if !errors.Is(global.GVA_DB.Where("name = ? ", grade.Name).First(&model.Grade{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同年级")
	}
	return global.GVA_DB.Create(&grade).Error
}
