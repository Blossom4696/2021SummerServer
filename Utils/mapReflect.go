package Utils

func GradeIntToString(gradeInt int64) (gradeString string) {
	gradeMap := make(map[int64]string)
	gradeMap[7] = "七年级"
	gradeMap[8] = "八年级"
	gradeMap[9] = "九年级"
	gradeMap[10] = "高一"
	gradeMap[11] = "高二"
	gradeMap[12] = "高三"
	return gradeMap[gradeInt]
}

func GradeStringToInt(gradeString string) (gradeInt int64) {
	gradeMap := make(map[string]int64)
	gradeMap["七年级"] = 7
	gradeMap["八年级"] = 8
	gradeMap["九年级"] = 9
	gradeMap["高一"] = 10
	gradeMap["高二"] = 11
	gradeMap["高三"] = 12
	return gradeMap[gradeString]
}
