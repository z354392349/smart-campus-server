package service

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"

	"gorm.io/gorm"
)

//@author: 张佳伟
//@function: CreateGrade
//@description: 新增年级
//@param: api model.SysApi
//@return: err error

func CreateGrade(grade model.Grade) (err error) {
	if !errors.Is(global.GVA_DB.Where("name = ? ", grade.Name).First(&model.Grade{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同年级")
	}
	return global.GVA_DB.Create(&grade).Error
}

//@author: 张佳伟
//@function: GetCreateList
//@description: 查询年级列表
//@param: info request.PageInfo
//@return: err error list interface{}  total int64
func GetCreateList(info request.PageInfo) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.Grade{})
	var gradeList []model.Grade
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Limit(limit).Offset(offset).Find(&gradeList).Error
	return err, gradeList, total
}
