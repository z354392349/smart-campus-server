package service

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @Author: 张佳伟
// @Function:CreateAllotExamRoom
// @Description:创建考场分配列表
// @Router:/allotExamRoom/createAllotExamRoom
// @Date: 2022/7/20 15:11

func CreateAllotExamRoom(Allots []model.AllotExamRoom) (err error) {
	return global.GVA_DB.Create(&Allots).Error
}

// @Author: 张佳伟
// @Function: 获取考场分配列表
// @Description: GetAllotExamRoomList
// @Router:/allotExamRoom/getAllotExamRoomList
// @Date: 2022/7/20 15:09

func GetAllotExamRoomList(info request.SearchAllotExamRoomParams) (err error, list interface{}, total int64) {
	fmt.Println(info)
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&model.AllotExamRoom{})
	var allotExamRoomList []model.AllotExamRoom

	//if info.Name != "" {
	//	db = db.Where("Name = ?", info.Name)
	//}
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Preload("ExamItem").Find(&allotExamRoomList).Error
	return err, allotExamRoomList, total
}

// @Author: 张佳伟
// @Function: UpAllotExamRoom
// @Description: 批量更新考场分配信息
// @Router: /allotExamRoom/upAllotExamRoom
// @Date: 2022/7/20 15:11

func UpAllotExamRoom(Allots []model.AllotExamRoom) (err error) {
	for _, allot := range Allots {
		err = global.GVA_DB.Where("id = ?", allot.ID).Find(&model.AllotExamRoom{}).Updates(&allot).Error
	}
	return err
}
