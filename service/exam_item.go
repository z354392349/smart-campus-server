package service

import (
	"errors"
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/utils"
	"strconv"
	"strings"
)

// @Author: 张佳伟
// @Function:AllotExamItemRoom
// @Description:考试项，考场分配
// @Router:	/examItem/allotExamItemRoom
// @Date:2022/07/25 16:47:35

func AllotExamItemRoom(info request.AllotExamRoomItem) (err error) {

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

	// 查看是否重复已经分配
	if examItem.ExamRoomIDs != "" {
		return errors.New("已经分配过考场,请勿重复分配")
	}

	// 查看时间是否冲突
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

	// 判断是否可以容纳学生
	var student []model.Student
	var examRoom []model.ExamRoom
	var examAmountTotal int

	if err = global.GVA_DB.Model(&model.Student{}).Debug().Where("grade_id = ?", info.GradeID).Order("rand()").Find(&student).Error; err != nil {
		return
	}
	if err = global.GVA_DB.Model(&model.ExamRoom{}).Debug().Where("id  in (?)", utils.StringToIntArr(info.ExamRoomIDs)).Find(&examRoom).Error; err != nil {
		return
	}

	for _, v := range examRoom {
		examAmountTotal += v.Amount
	}

	if len(student) > examAmountTotal {
		return errors.New("考场容量无法容纳" + strconv.Itoa(len(student)) + "人")
	}

	// 可以容纳学生，更新 ExamItem 的 ExamRoomIDs 字段，说明占用的 考场号
	if err = global.GVA_DB.Model(&model.ExamItem{}).Debug().Where("id = ? ", info.ExamItemID).Updates(model.ExamItem{ExamRoomIDs: info.ExamRoomIDs}).Error; err != nil {
		return
	}

	// 创建 AllotExamRoom 数据
	// 1.找到年级下所有的学生， GradeID
	// 2.查找出所有的考场，ExamRoomIDs,
	// 3.根据考场随机抽取 学生，组合成 数据
	// var allotExamRoom [] model.AllotExamRoom  // StudentID // ExamID // ExamItemID // ExamRoomID

	ExamRoomI := 0                              // 当前使用的考场数组下标
	ExamRoomNum := examRoom[0].Amount           // 当前已使用的考场有多少座位数
	var allotExamRoomList []model.AllotExamRoom // 分配考场信息

	for i, v := range student {
		if i == ExamRoomNum {
			ExamRoomI += 1
			ExamRoomNum += examRoom[ExamRoomI].Amount
		}
		allotExamRoom := model.AllotExamRoom{
			StudentID:  v.ID,
			ExamID:     info.ExamID,
			CourseID:   examItem.CourseID,
			ExamRoomID: examRoom[ExamRoomI].ID,
		}
		allotExamRoomList = append(allotExamRoomList, allotExamRoom)
	}

	ExamRoomI = 0 // 当前使用的考场数组下标
	fmt.Println(ExamRoomI)

	err = global.GVA_DB.Create(&allotExamRoomList).Error
	return err
}

// @Author: 张佳伟
// @Function:CancelAllotExamItemRoom
// @Description:撤销已分配的考场
// @Router: /examItem/cancelAllotExamItemRoom
// @Date:2022/07/26 16:03:21

func CancelAllotExamItemRoom(info request.CancelAllotExamRoomItem) (err error) {
	// 1.删除 exam_items，  根据 exam_id、course_id， 或者 ID 删除
	// 2.删除 allot_exam_rooms， 根据 exam_item_id

	if err = global.GVA_DB.Debug().Where("id = ?", info.ExamItemID).Delete(&model.ExamItem{}).Error; err != nil {
		return err
	}

	if err = global.GVA_DB.Debug().Where("exam_item_id = ?", info.ExamItemID).Delete(&model.AllotExamRoom{}).Error; err != nil {
		return err
	}

	return
}
