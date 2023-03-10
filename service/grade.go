package service

import (
	"errors"
	"fmt"
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

func GetCreateList(info request.SearchGradeParams) (err error, list interface{}, total int64) {
	fmt.Println(info)
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.Grade{})
	var gradeList []model.Grade

	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Debug().Limit(limit).Offset(offset).Find(&gradeList).Error
	return err, gradeList, total
}

// @Author: 张佳伟
// @Function: UpCreate
// @Description: 更新年级
// @Router: /grade/upGrade
// @Date:2022/07/09 10:50:09
func UpCreate(grade model.Grade) (err error) {
	err = global.GVA_DB.Where("id = ?", grade.ID).First(&model.Grade{}).Updates(&grade).Error
	return err
}

// @Author: 张佳伟
// @Function:
// @Description:
// @Router:
// @Date:2022/07/09 10:32:22
func DeleteGrade(grade model.Grade) (err error) {
	err = global.GVA_DB.Debug().Where("id = ?", grade.ID).Delete(&grade).Error
	return err
}
