package service

import (
	"errors"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"

	"gorm.io/gorm"
)

func CreateClass1(class model.Class) (err error) {
	if !errors.Is(global.GVA_DB.Where("name = ? ", class.Name).First(&model.Class{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在相同班级")
	}
	return global.GVA_DB.Create(&class).Error
}

func GetClassList1(info request.SearchClassParams) (err error, list interface{}, total int64) {
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

func UpClass1(class model.Class) (err error) {
	err = global.GVA_DB.Where("id = ?", class.ID).First(&model.Class{}).Updates(&class).Error
	return err
}

func DeleteClass1(class model.Class) (err error) {
	err = global.GVA_DB.Debug().Where("id = ?", class.ID).Delete(&class).Error
	return err
}
