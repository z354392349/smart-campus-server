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

	if info.TeacherName != "" {
		db = db.Where("teachers.name LIKE ?", "%"+info.TeacherName+"%")
	}
	if info.CarNum != "" {
		db = db.Where("car_num LIKE ?", "%"+info.CarNum+"%")
	}
	if info.StartTime != 0 {
		db = db.Where("start_time = ?", info.StartTime)
	}
	if info.EndTime != 0 {
		db = db.Where("end_time = ?", info.EndTime)
	}

	leftJoinSql1 := "left join teachers on teachers.id = car_accesses.teacher_id" // 教师姓名
	selectSql := "car_accesses.*, teachers.name as teacher_name"
	if err = db.Debug().Limit(limit).Offset(offset).Select(selectSql).Joins(leftJoinSql1).Find(&carAccessList).Error; err != nil {
		return
	}

	err = db.Count(&total).Limit(-1).Offset(-1).Error
	return err, carAccessList, total
}

func UpClass1(class model.Class) (err error) {
	err = global.GVA_DB.Where("id = ?", class.ID).First(&model.Class{}).Updates(&class).Error
	return err
}

func DeleteClass1(class model.Class) (err error) {
	err = global.GVA_DB.Debug().Where("id = ?", class.ID).Delete(&class).Error
	return err
}

func SetClassMonitor1(info request.SetClassMonitor) (err error) {
	err = global.GVA_DB.Model(&model.Class{}).Where("id = ?", info.ClassID).Updates(&model.Class{MonitorID: info.StudentID}).Error
	return err
}

func SetClassTeacher1(info request.SetClassTeacher) (err error) {
	err = global.GVA_DB.Debug().Model(&model.Class{}).Where("id = ?", info.ClassID).Updates(&model.Class{TeacherID: info.TeacherID}).Error
	return err
}
