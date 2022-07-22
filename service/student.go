package service

import (
	"fmt"
	"gin-vue-admin/global"
	"gin-vue-admin/model"
	"gin-vue-admin/model/request"
)

// @Author: 张佳伟
// @Function: CreateStudent
// @Description: 创建学生
// @Router:/student/createStudent
// @Date: 2022/7/15 10:40:12

func CreateStudent(student model.Student) (err error) {
	return global.GVA_DB.Create(&student).Error
}

// @Author: 张佳伟
// @Function: GetStudentList1
// @Description: 创建学生
// @Router:/student/getStudentList
// @Date: 2022/7/15 10:50:22

func GetStudentList(info request.SearchStudentParams) (err error, list interface{}, total int64) {
	fmt.Println(info)
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.Student{})
	var studentList []model.Student
	if info.Name != "" {
		db = db.Where("name LIKE ?", "%"+info.Name+"%")
	}
	if info.GradeID != 0 {
		db = db.Where("grade_id = ?", info.GradeID)
	}
	if info.ClassID != 0 {
		db = db.Where("class_id = ?", info.ClassID)
	}
	err = db.Count(&total).Error
	//err = db.Debug().Limit(limit).Offset(offset).Preload("Grade").Preload("Class").Find(&studentList).Error
	join := "left join grades on grades.id = students.gradeID"
	//join2 := "left join grades on grades.id = students.gradeID"
	err = db.Debug().Limit(limit).Select("grades.name  as gradeName").Offset(offset).Joins(join).Find(&studentList).Error
	return err, studentList, total
}

// @Author: 张佳伟
// @Function: UpStudent
// @Description: 更新学生
// @Router:/student/upStudent
// @Date: 2022/7/15 10:44

func UpStudent(student model.Student) (err error) {
	err = global.GVA_DB.Where("id = ?", student.ID).First(&model.Student{}).Updates(&student).Error
	return err
}

// @Author: 张佳伟
// @Function: DeleteStudent
// @Description:删除学生
// @Router:/student/deleteStudent
// @Date: 2022/7/15 10:46

func DeleteStudent(student model.Student) (err error) {
	err = global.GVA_DB.Debug().Where("id = ?", student.ID).Delete(&student).Error
	return err
}

// @Author: 张佳伟
// @Function: SetStudentsGradeAndClass
// @Description: 批量设置学生的年级和班级
// @Router: /student/setStudentsGradeAndClass
// @Date: 2022/7/20 10:14:36

func SetStudentsGradeAndClass(info request.SetStudentsGradeAndClass) (err error) {
	studentsID := info.StudentsID
	gradeID := info.GradeID
	classID := info.ClassID
	err = global.GVA_DB.Model(&model.Student{}).Where("id IN ?", studentsID).Updates(model.Student{GradeID: gradeID, ClassID: classID}).Error
	return err
}
