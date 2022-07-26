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

	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}

	if info.GradeID != 0 {
		db = db.Where("grade_id = ?", info.GradeID)
	}
	if info.ClassID != 0 {
		db = db.Where("grade_id = ?", info.ClassID)
	}
	if err = db.Count(&total).Error; err != nil {
		return
	}

	// TODO:sql 没写完呢
	// leftJoinSql := "left join courses on exam_items.course_id = courses.id"
	leftJoinSql1 := "left join students on students.id = allot_exam_rooms.student_id"       // 学生姓名
	leftJoinSql2 := "left join exams on exams.id = allot_exam_rooms.exam_id"                // 考试名称
	leftJoinSql3 := "left join exam_items on exam_items.id = allot_exam_rooms.exam_item_id" // 考试项目

	// selectSql := "exam_items.*, courses.name as course_name "
	// .Select(selectSql)
	err = db.Limit(limit).Debug().Offset(offset).Joins(leftJoinSql1).Joins(leftJoinSql2).Joins(leftJoinSql3).Find(&allotExamRoomList).Error
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
