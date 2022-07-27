package service

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

func GetAllotExamRoomList(info request.SearchAllotExamRoomParams) (err error, list interface{}, total int64) {
	fmt.Println(info)
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	db := global.GVA_DB.Model(&model.AllotExamRoom{})
	var allotExamRoomList []model.AllotExamRoom

	leftJoinSql1 := "left join students on students.id = allot_exam_rooms.student_id"       // 学生姓名
	leftJoinSql2 := "left join exams on exams.id = allot_exam_rooms.exam_id"                // 考试名称
	leftJoinSql3 := "left join grades on grades.id = students.grade_id"                     // 年级名称
	leftJoinSql4 := "left join classes on classes.id = students.class_id"                   // 班级名称
	leftJoinSql5 := "left join courses on courses.id = allot_exam_rooms.course_id"          // 科目名称
	leftJoinSql6 := "left join exam_rooms on exam_rooms.id = allot_exam_rooms.exam_room_id" // 考场名称

	selectSql := "allot_exam_rooms.id, students.name as student_name, exams.name as  exam_name, grades.name as grade_name,  classes.name as class_name, courses.name as course_name, exam_rooms.name as exam_room_name"
	db = db.Limit(limit).Debug().Offset(offset).Select(selectSql).Joins(leftJoinSql1).Joins(leftJoinSql2).Joins(leftJoinSql3).Joins(leftJoinSql4).Joins(leftJoinSql5).Joins(leftJoinSql6)

	if err = db.Count(&total).Error; err != nil {
		return
	}

	if info.Name != "" {
		db = db.Where("students.name LIKE ?", "%"+info.Name+"%")
	}

	if info.GradeID != 0 {
		db = db.Where("grade_id = ?", info.GradeID)
	}
	if info.ClassID != 0 {
		db = db.Where("grade_id = ?", info.ClassID)
	}

	err = db.Find(&allotExamRoomList).Error

	return err, allotExamRoomList, total
}
