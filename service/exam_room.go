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
// @Function:CreateExamRoom
// @Description:创建考场
// @Router:/examRoom/createExamRoom
// @Date: 2022/7/15 15:11

func CreateExamRoom(examRoom model.ExamRoom) (err error) {
	if !errors.Is(global.GVA_DB.Where("name = ? ", examRoom.Name).First(&model.ExamRoom{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("获取考场列表")
	}
	return global.GVA_DB.Create(&examRoom).Error
}

// @Author: 张佳伟
// @Function: GetExamRoomList
// @Description: 获取考场列表
// @Router:/examRoom/getExamRoomList
// @Date: 2022/7/15 15:17

func GetExamRoomList(info request.SearchExamRoomParams) (err error, list interface{}, total int64) {
	fmt.Println(info)
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.ExamRoom{})
	var examRoomList []model.ExamRoom

	if info.Name != "" {
		db = db.Where("Name = ?", info.Name)
	}
	err = db.Count(&total).Error
	err = db.Debug().Limit(limit).Offset(offset).Find(&examRoomList).Error
	return err, examRoomList, total
}

// @Author: 张佳伟
// @Function: UpExamRoom
// @Description: 更新考场
// @Router: /examRoom/UpExamRoom
// @Date: 2022/7/15 15:23:12

func UpExamRoom(examRoom model.ExamRoom) (err error) {
	err = global.GVA_DB.Where("id = ?", examRoom.ID).First(&model.ExamRoom{}).Updates(&examRoom).Error
	return err
}

// @Author: 张佳伟
// @Function: DeleteExamRoom
// @Description:删除考场
// @Router: /examRoom/deleteExamRoom
// @Date: 2022/7/15 15:23:32

func DeleteExamRoom(examRoom model.ExamRoom) (err error) {
	err = global.GVA_DB.Debug().Where("id = ?", examRoom.ID).Delete(&examRoom).Error
	return err
}
