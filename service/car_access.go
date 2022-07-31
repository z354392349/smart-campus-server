package service

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
)

// @Author: 张佳伟
// @Function:CreateCarAccess
// @Description:创建通行记录
// @Router:/carAccess/createCarAccess
// @Date:2022/07/29 10:49:26

func CreateCarAccess(carAccess model.CarAccess) (err error) {
	return global.GVA_DB.Debug().Create(&carAccess).Error
}

// @Author: 张佳伟
// @Function:GetCarAccessList
// @Description:获取车辆通行列表
// @Router:/carAccess/getCarAccessList
// @Date:2022/07/29 17:35:43

func GetCarAccessList(info request.SearchCarAccess) (err error, list interface{}, total int64) {
	fmt.Println(info)
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.CarAccess{})
	var carAccessList []response.CarAccess

	leftJoinSql1 := "left join teachers on teachers.id = car_accesses.teacher_id" // 教师姓名
	selectSql := "car_accesses.*, teachers.name as teacher_name"
	db = db.Debug().Select(selectSql).Joins(leftJoinSql1)

	if info.TeacherName != "" {
		db = db.Where("teachers.name LIKE ?", "%"+info.TeacherName+"%")
	}
	if info.CarNum != "" {
		db = db.Where("car_num LIKE ?", "%"+info.CarNum+"%")
	}
	if info.StartTime != 0 {
		db = db.Where("time >= ?", info.StartTime)
	}
	if info.EndTime != 0 {
		db = db.Where("time <= ?", info.EndTime)
	}

	if err = db.Debug().Limit(limit).Offset(offset).Find(&carAccessList).Error; err != nil {
		return
	}

	err = db.Count(&total).Limit(-1).Offset(-1).Error
	return err, carAccessList, total
}
