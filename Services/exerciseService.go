package Services

import (
	"github.com/bigby/project/Models"
	"github.com/bigby/project/Utils"
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

type ExerciseField struct {
	Egrade      []string
	Esubject    []string
	Etype       map[string][]string
	Edifficulty []string
	Eunit       map[string][]string
}

func (exercise *Exercise) Insert() (Eid int64, err error) {
	var exerciseModel Models.Exercise

	// exerciseModel.Eid = exercise.Eid
	exerciseModel.Etitle = exercise.Etitle
	exerciseModel.Edescription = exercise.Edescription
	exerciseModel.EnameTxt = exercise.EnameTxt
	exerciseModel.EnamePath = exercise.EnamePath
	exerciseModel.EanswerTxt = exercise.EanswerTxt
	exerciseModel.EanswerPath = exercise.EanswerPath
	exerciseModel.Egrade = exercise.Egrade
	exerciseModel.Esubject = exercise.Esubject
	exerciseModel.Etype = exercise.Etype
	exerciseModel.Edifficulty = exercise.Edifficulty
	exerciseModel.Eunit = exercise.Eunit

	Eid, err = exerciseModel.Insert()

	return

}

func (exercise *Exercise) Update(Eid int64) (result Models.Exercise, err error) {
	var exerciseModel Models.Exercise

	exerciseModel.Eid = exercise.Eid
	exerciseModel.Etitle = exercise.Etitle
	exerciseModel.Edescription = exercise.Edescription
	exerciseModel.EnameTxt = exercise.EnameTxt
	exerciseModel.EnamePath = exercise.EnamePath
	exerciseModel.EanswerTxt = exercise.EanswerTxt
	exerciseModel.EanswerPath = exercise.EanswerPath
	exerciseModel.Egrade = exercise.Egrade
	exerciseModel.Esubject = exercise.Esubject
	exerciseModel.Etype = exercise.Etype
	exerciseModel.Edifficulty = exercise.Edifficulty
	exerciseModel.Eunit = exercise.Eunit

	result, err = exerciseModel.Update(exerciseModel.Eid)

	return
}

func (exercise *Exercise) Delete(Eid int64) (result Models.Exercise, err error) {
	var exerciseModel Models.Exercise

	exerciseModel.Eid = exercise.Eid

	result, err = exerciseModel.Delete(exerciseModel.Eid)

	return
}

func (exercise *Exercise) Query() (result []Models.Exercise, err error) {
	var exerciseModel Models.Exercise

	result, err = exerciseModel.Query()

	return
}

func (exercise *Exercise) QueryByEid() (result Models.Exercise, err error) {
	var exerciseModel Models.Exercise

	result, err = exerciseModel.QueryByEId(exercise.Eid)

	return
}

func (exercise *Exercise) QueryBySort() (result []Models.Exercise, err error) {
	var exerciseModel Models.Exercise

	exerciseModel.Esubject = exercise.Esubject
	exerciseModel.Egrade = exercise.Egrade
	exerciseModel.Edifficulty = exercise.Edifficulty
	exerciseModel.Etype = exercise.Etype
	exerciseModel.Eunit = exercise.Eunit

	result, err = exerciseModel.QueryBySort()

	return
}

func (exercise *Exercise) QueryAllField() (result ExerciseField, err error) {
	var exerciseModel Models.Exercise

	grades, _ := exerciseModel.QueryGradeField()

	subjects, _ := exerciseModel.QuerySubjectField()

	types, _ := exerciseModel.QueryTypeField()

	difficulty, _ := exerciseModel.QueryDifficultyField()

	units, err := exerciseModel.QueryUnitField()

	// 处理年级
	result.Egrade = append(result.Egrade, "不限")
	for i := 0; i < len(grades); i++ {
		if grades[i].Egrade != "" {
			result.Egrade = append(result.Egrade, grades[i].Egrade)
		}
	}

	// 处理科目
	result.Esubject = append(result.Esubject, "不限")
	for i := 0; i < len(subjects); i++ {
		if subjects[i].Esubject != "" {
			result.Esubject = append(result.Esubject, subjects[i].Esubject)
		}
	}

	// 处理类型
	result.Etype = make(map[string][]string)
	for i := 0; i < len(types); i++ {
		result.Etype["不限"] = append(result.Etype["不限"], types[i].Etype)
		if _, ok := result.Etype[types[i].Esubject]; !ok {
			result.Etype[types[i].Esubject] = append(result.Etype[types[i].Esubject], "不限")
		}
		if types[i].Etype != "" {
			result.Etype[types[i].Esubject] = append(result.Etype[types[i].Esubject], types[i].Etype)
		}
	}
	result.Etype["不限"] = DeleteDuplicateValue(result.Etype["不限"])
	result.Etype["不限"] = append(result.Etype["不限"][:0], "不限")

	// 处理单元
	result.Eunit = make(map[string][]string)
	for i := 0; i < len(units); i++ {

		if _, ok := result.Eunit[units[i].Esubject]; !ok {
			result.Eunit[units[i].Esubject] = append(result.Eunit[units[i].Esubject], "不限")
		}
		if units[i].Eunit != "" {
			result.Eunit["不限"] = append(result.Eunit["不限"], units[i].Eunit)
			result.Eunit[units[i].Esubject] = append(result.Eunit[units[i].Esubject], units[i].Eunit)
		}
	}
	result.Eunit["不限"] = DeleteDuplicateValue(result.Eunit["不限"])
	result.Eunit["不限"] = append(result.Eunit["不限"][:0], "不限")

	// 处理难度
	result.Edifficulty = append(result.Edifficulty, "不限")
	for i := 0; i < len(difficulty); i++ {
		if difficulty[i].Edifficulty != "" {
			result.Edifficulty = append(result.Edifficulty, difficulty[i].Edifficulty)
		}
	}

	return
}

// 搜索匹配
func (exercise *Exercise) Search(searchString string) (result []Exercise, err error) {
	var exerciseModel Models.Exercise

	// 匹配title
	tmpTitle, err := exerciseModel.Search(searchString, "Etitle")
	if err != nil {
		return
	}
	// 匹配subject
	tmpSubject, err := exerciseModel.Search(searchString, "Esubject")
	if err != nil {
		return
	}
	// 匹配type
	tmpType, err := exerciseModel.Search(searchString, "Etype")
	if err != nil {
		return
	}
	// 匹配difficulty
	tmpDifficulty, err := exerciseModel.Search(searchString, "Edifficulty")
	if err != nil {
		return
	}
	// 匹配unit
	tmpUnit, err := exerciseModel.Search(searchString, "Eunit")
	if err != nil {
		return
	}
	// 匹配grade
	tmpGrade, err := exerciseModel.Search(searchString, "Egrade")
	if err != nil {
		return
	}

	tmpExercise := append(tmpTitle, tmpSubject...)
	tmpExercise = append(tmpExercise, tmpType...)
	tmpExercise = append(tmpExercise, tmpDifficulty...)
	tmpExercise = append(tmpExercise, tmpUnit...)
	tmpExercise = append(tmpExercise, tmpGrade...)
	tmpExercise = DeleteDuplicateValueExercise(tmpExercise)

	result = make([]Exercise, len(tmpExercise))
	// 把teacher model转为teacher service
	for i := 0; i < len(tmpExercise); i++ {
		err = Utils.CopyFields(&result[i], tmpExercise[i])
		if err != nil {
			return
		}
	}

	return
}

func DeleteDuplicateValue(s []string) (ret []string) {
	tmpM := make(map[string]int) // key的类型要和切片中的数据类型一致
	for _, v := range s {
		tmpM[v] = 1
	}
	// 先清空s
	s = []string{}
	for i := range tmpM {
		s = append(s, i)
	}
	return s
}

// 去重切片
func DeleteDuplicateValueExercise(s []Models.Exercise) (ret []Models.Exercise) {
	tmpM := make(map[Models.Exercise]int) // key的类型要和切片中的数据类型一致
	for _, v := range s {
		tmpM[v] = 1
	}
	// 先清空s
	s = []Models.Exercise{}
	for i := range tmpM {
		s = append(s, i)
	}
	return s
}
