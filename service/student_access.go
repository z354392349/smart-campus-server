package service

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
)

func CreateStudentAccess(studentAccess model.StudentAccess) (err error) {
	return global.GVA_DB.Debug().Create(&studentAccess).Error
}
func CreateStudentAccessList(studentAccess []model.StudentAccess) (err error) {
	return global.GVA_DB.Debug().Create(&studentAccess).Error
}

func GetStudentAccessList(info request.SearchStudentAccess) (err error, list interface{}, total int64) {
	fmt.Println(info)
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.StudentAccess{})
	var studentAccess []response.StudentAccess

	leftJoinSql1 := "left join students on students.id = student_accesses.student_id" // 学生姓名
	leftJoinSql2 := "left join grades on grades.id = students.grade_id"               // 年级名称
	leftJoinSql3 := "left join classes on classes.id = students.class_id"             // 班级名称
	selectSql := "student_accesses.*, students.name as student_name, grades.name as grade_name, classes.name as class_name"
	db = db.Debug().Select(selectSql).Joins(leftJoinSql1).Joins(leftJoinSql2).Joins(leftJoinSql3)

	if info.StudentName != "" {
		db = db.Where("students.name LIKE ?", "%"+info.StudentName+"%")
	}

	if info.GradeID != 0 {
		db = db.Where("grades.id = ?", info.GradeID)
	}
	if info.ClassID != 0 {
		db = db.Where("classes.id = ?", info.ClassID)
	}

	if info.StartTime != 0 {
		db = db.Where("time >= ?", info.StartTime)
	}
	if info.EndTime != 0 {
		db = db.Where("time <= ?", info.EndTime)
	}

	if err = db.Debug().Limit(limit).Offset(offset).Find(&studentAccess).Error; err != nil {
		return
	}

	err = db.Count(&total).Limit(-1).Offset(-1).Error
	return err, studentAccess, total
}
