package service

import (
	"errors"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/utils"
	"strings"
)

func GetExamRoomList1(info request.SearchExamRoomParams) (err error, list interface{}, total int64) {
	fmt.Println(info)
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.ExamRoom{})
	var examRoomList []model.ExamRoom

	if info.Name != "" {
		db = db.Where("Name = ?", info.Name)
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}
	err = db.Debug().Limit(limit).Offset(offset).Preload("Teacher").Find(&examRoomList).Error
	return err, examRoomList, total
}

// @Author: 张佳伟
// @Function:UpExamItemRoomAllot
// @Description:考试项，考场分配
// @Router:	/examItem/upExamItemRoomAllot
// @Date:2022/07/25 16:47:35

func UpExamItemRoomAllot(info request.SetExamRoomItemAllot) (err error) {

	// ExamRoomIDs := info.ExamRoomIDs
	// 做判断，区分是否可以 分配
	// 1.根据 ExamItemID, 查询开始和结束时间，
	// 2.根据查询出的 开始和结束时间，查询出所有 在这个范围的 考试
	// 3.判断他们的 不符合的考场id，
	// 4.返回考场名称
	var examItem model.ExamItem            // 作为比较的考试项目
	var examItemList []model.ExamItem      // 可能冲突的考试项
	var clashExamRoomList []model.ExamRoom // 确定冲突的考场

	global.GVA_DB.Model(&model.ExamItem{}).Debug().Where("id = ?", info.ExamItemID).Find(&examItem)
	global.GVA_DB.Model(&model.ExamItem{}).Debug().Where("end_time > ?  and exam_room_ids is not null ", examItem.StartTime).Find(&examItemList)

	var clashExamRoomIDs []int // 保存所有冲突的id
	for _, v := range examItemList {
		clashExamRoomID := utils.Intersection(utils.StringToInt(strings.Split(v.ExamRoomIDs, ",")), utils.StringToInt(strings.Split(info.ExamRoomIDs, ","))) // 获取冲突的 ID
		clashExamRoomIDs = append(clashExamRoomIDs, clashExamRoomID...)
	}

	global.GVA_DB.Model(&model.ExamRoom{}).Debug().Where("id in ? ", clashExamRoomIDs).Find(&clashExamRoomList) // 获取确定冲突的考场

	if len(clashExamRoomList) != 0 {
		var str = ""
		for _, v := range clashExamRoomList {
			str += v.Name + ","
		}
		str = str[0 : len(str)-1]
		return errors.New(str + "时间冲突")
	}

	// 更新 ExamItem 的 ExamRoomIDs 字段，说明占用的 考场号
	if err = global.GVA_DB.Model(&model.ExamItem{}).Debug().Where("id = ? ", info.ExamItemID).Updates(model.ExamItem{ExamRoomIDs: info.ExamRoomIDs}).Error; err != nil {
		return
	}

	// 创建 AllotExamRoom 数据

	// err = global.GVA_DB.Where("id = ?", examRoom.ID).First(&model.ExamRoom{}).Updates(&examRoom).Error
	return err
}
