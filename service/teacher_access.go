package service

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
)

// @Author: 张佳伟
// @Function:CreateTeacherAccess
// @Description:创建教师通行记录
// @Router:/teacherAccess/createTeacherAccess
// @Date:2022/08/01 10:45:48

func CreateTeacherAccess(teacherAccess model.TeacherAccess) (err error) {
	return global.GVA_DB.Debug().Create(&teacherAccess).Error
}

// @Author: 张佳伟
// @Function:GetTeacherAccessList
// @Description:获取教师通行记录
// @Router:/teacherAccess/getTeacherAccessList
// @Date:2022/08/01 10:45:48

func GetTeacherAccessList(info request.SearchTeacherAccess) (err error, list interface{}, total int64) {
	fmt.Println(info)
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.TeacherAccess{})
	var teacherAccess []response.TeacherAccess

	leftJoinSql1 := "left join teachers on teachers.id = teacher_accesses.teacher_id" // 教师姓名
	selectSql := "teacher_accesses.*, teachers.name as teacher_name"
	db = db.Debug().Select(selectSql).Joins(leftJoinSql1)

	if info.TeacherName != "" {
		db = db.Where("teachers.name LIKE ?", "%"+info.TeacherName+"%")
	}

	if info.StartTime != 0 {
		db = db.Where("time >= ?", info.StartTime)
	}
	if info.EndTime != 0 {
		db = db.Where("time <= ?", info.EndTime)
	}

	if err = db.Debug().Limit(limit).Offset(offset).Find(&teacherAccess).Error; err != nil {
		return
	}

	err = db.Count(&total).Limit(-1).Offset(-1).Error
	return err, teacherAccess, total
}
