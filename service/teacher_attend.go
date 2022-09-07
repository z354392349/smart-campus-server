package service

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/request"
	"gin-vue-admin/model/response"
)

// @Author: 张佳伟
// @Function: GetTeacherAttendList
// @Description: 获取教师通勤列表
// @Router:/teacherAttend/getTeacherAttendList
// @Date:2022/08/02 14:38:36

func GetTeacherAttendList(info request.SearchTeacherAttend) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	var teacherAttend []response.TeacherAttend

	leftJoinSql1 := "left join teachers on teachers.id = tb.teacher_id "
	selectSql := "select if(direction=1, min(time), max(time)) as time, direction, teacher_id, teachers.name  as teacher_name from  (select teacher_id, time, direction  from teacher_accesses union all select teacher_id, time, direction from car_accesses) tb "
	whereSql := ""
	groupSql := "group by left(FROM_UNIXTIME(time), 10), direction, tb.teacher_id "
	orderSql := "ORDER by time desc "
	var paras []interface{}

	// if info.TeacherID != 0 {
	// 	db = db.Where("teacher_id = ?", info.TeacherID)
	// }

	if info.TeacherName != "" {
		whereSql += "teachers.name LIKE ? "
		paras = append(paras, "%"+info.TeacherName+"%")
	}

	if info.TeacherID != 0 {
		whereSql += "teachers.id = ? "
		paras = append(paras, info.TeacherID)
	}

	if info.StartTime != 0 {
		if whereSql != "" {
			whereSql += " and "
		}
		whereSql += "time >= ? "
		paras = append(paras, info.StartTime)
	}

	if info.EndTime != 0 {
		if whereSql != "" {
			whereSql += " and "
		}
		whereSql += "time <= ? "
		paras = append(paras, info.EndTime)
	}

	if whereSql != "" {
		whereSql = "where " + whereSql
	}

	global.GVA_DB.Debug().Raw("select count(*) as total from ("+selectSql+leftJoinSql1+whereSql+groupSql+orderSql+") tb2", paras...).Scan(&total)

	pageSql := "LIMIT ? OFFSET ? "
	paras = append(paras, limit)
	paras = append(paras, offset)

	global.GVA_DB.Debug().Raw(selectSql+leftJoinSql1+whereSql+groupSql+orderSql+pageSql, paras...).Scan(&teacherAttend)

	return err, teacherAttend, total
}
