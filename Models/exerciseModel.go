package Models

import (
	orm "github.com/bigby/project/Database"
)

type Exercise struct {
	Eid          int64  `gorm:"primary_key;column:Eid;AUTO_INCREMENT"`
	Etitle       string `gorm:"column:Etitle"`
	Edescription string `gorm:"column:Edescription"`
	EnameTxt     string `gorm:"column:EnameTxt"`
	EnamePath    string `gorm:"column:EnamePath"`
	EanswerTxt   string `gorm:"column:EanswerTxt"`
	EanswerPath  string `gorm:"column:EanswerPath"`
	Egrade       string `gorm:"column:Egrade"`
	Esubject     string `gorm:"column:Esubject"`
	Etype        string `gorm:"column:Etype"`
	Edifficulty  string `gorm:"type:enum('简单','中等','困难','竞赛');column:Edifficulty;default:简单"`
	Eunit        string `gorm:"column:Eunit"`
}

type ExerciseGradeField struct {
	Egrade string `gorm:"column:Egrade"`
}

type ExerciseSubjectField struct {
	Esubject string `gorm:"column:Esubject"`
}

type ExerciseTypeField struct {
	Esubject string `gorm:"column:Esubject"`
	Etype    string `gorm:"column:Etype"`
}

type ExerciseDifficultyField struct {
	Edifficulty string `gorm:"type:enum('简单','中等','困难','竞赛');column:Edifficulty;default:简单"`
}

type ExerciseUnitField struct {
	Esubject string `gorm:"column:Esubject"`
	Eunit    string `gorm:"column:Eunit"`
}

//增
func (exercise Exercise) Insert() (Eid int64, err error) {
	//添加数据
	result := orm.Eloquent.Table("EXERCISE").Create(&exercise)
	Eid = exercise.Eid
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

//删
func (exercise *Exercise) Delete(Eid int64) (Result Exercise, err error) {

	if err = orm.Eloquent.Table("EXERCISE").Select([]string{"Eid"}).First(&exercise, Eid).Error; err != nil {
		return
	}

	if err = orm.Eloquent.Table("EXERCISE").Delete(&exercise).Error; err != nil {
		return
	}
	Result = *exercise
	return
}

//改
func (exercise *Exercise) Update(Eid int64) (updateExercise Exercise, err error) {

	if err = orm.Eloquent.Table("EXERCISE").Select([]string{"Eid"}).First(&updateExercise, Eid).Error; err != nil {
		return
	}

	//参数1:是要修改的数据
	//参数2:是修改的数据
	if err = orm.Eloquent.Table("EXERCISE").Model(&updateExercise).Updates(&exercise).Error; err != nil {
		return
	}
	return
}

//查
func (exercise *Exercise) Query() (exercises []Exercise, err error) {

	if err = orm.Eloquent.Table("EXERCISE").Find(&exercises).Error; err != nil {
		return
	}
	return
}

//按Eid查
func (exercise *Exercise) QueryByEId(Eid int64) (result Exercise, err error) {

	if err = orm.Eloquent.Table("EXERCISE").Where("Eid = ?", Eid).First(&result).Error; err != nil {
		return
	}
	return
}

// 查属性的所有值，分以下五个函数
// 查年级
func (exercise *Exercise) QueryGradeField() (result []ExerciseGradeField, err error) {
	if err = orm.Eloquent.Table("EXERCISE").Select("Egrade").Group("Egrade").Scan(&result).Error; err != nil {
		return
	}
	return
}

// 查科目
func (exercise *Exercise) QuerySubjectField() (result []ExerciseSubjectField, err error) {
	if err = orm.Eloquent.Table("EXERCISE").Select("Esubject").Group("Esubject").Scan(&result).Error; err != nil {
		return
	}
	return
}

// 查类型
func (exercise *Exercise) QueryTypeField() (result []ExerciseTypeField, err error) {
	if err = orm.Eloquent.Table("EXERCISE").Select("Esubject, Etype").Group("Esubject, Etype").Scan(&result).Error; err != nil {
		return
	}
	return
}

// 查难度
func (exercise *Exercise) QueryDifficultyField() (result []ExerciseDifficultyField, err error) {
	if err = orm.Eloquent.Table("EXERCISE").Select("Edifficulty").Group("Edifficulty").Scan(&result).Error; err != nil {
		return
	}
	return
}

// 查单元
func (exercise *Exercise) QueryUnitField() (result []ExerciseUnitField, err error) {
	if err = orm.Eloquent.Table("EXERCISE").Select("Esubject, Eunit").Group("Esubject, Eunit").Scan(&result).Error; err != nil {
		return
	}
	return
}

// 筛选
func (exercise *Exercise) QueryBySort() (result []Exercise, err error) {
	sql := orm.Eloquent.Table("EXERCISE")
	if exercise.Esubject != "不限" {
		sql = sql.Where("Esubject = ?", exercise.Esubject)
	}
	if exercise.Etype != "不限" {
		sql = sql.Where("Etype = ?", exercise.Etype)
	}
	if exercise.Egrade != "不限" {
		sql = sql.Where("Egrade = ?", exercise.Egrade)
	}
	if exercise.Edifficulty != "不限" {
		sql = sql.Where("Edifficulty = ?", exercise.Edifficulty)
	}
	if exercise.Eunit != "不限" {
		sql = sql.Where("Eunit = ?", exercise.Eunit)
	}
	sql.Find(&result)
	return
}

//搜索匹配（字符串，搜索域）
func (exercise *Exercise) Search(searchString string, searchField string) (result []Exercise, err error) {

	if err = orm.Eloquent.Table("EXERCISE").Where(searchField+" LIKE ?", "%"+searchString+"%").Find(&result).Error; err != nil {
		return
	}
	return
}
