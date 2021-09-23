package Services

import "github.com/bigby/project/Models"

type Campus struct {
	Cid   int64  `gorm:"primary_key;column:Cid;AUTO_INCREMENT"`
	Cname string `gorm:"column:Cname"`
}

func (campus *Campus) QueryAll() (result []Models.Campus, err error) {
	var campusModel Models.Campus

	result, err = campusModel.QueryAll()

	return
}

func (campus *Campus) Insert() (Cid int64, err error) {
	var campusModel Models.Campus

	campusModel.Cid = campus.Cid
	campusModel.Cname = campus.Cname

	Cid, err = campusModel.Insert()

	return

}

func (campus *Campus) Delete(Cid int64) (result Models.Campus, err error) {
	var campusModel Models.Campus

	campusModel.Cid = campus.Cid

	result, err = campusModel.Delete(Cid)

	return
}

func (campus *Campus) Update(Cid int64) (result Models.Campus, err error) {
	var campusModel Models.Campus

	campusModel.Cid = campus.Cid
	campusModel.Cname = campus.Cname

	result, err = campusModel.Update(Cid)

	return
}

// 从数组更新
func (campus *Campus) UpdateFromArray(campusArray []Campus) (result []Models.Campus, err error) {
	var campusModel Models.Campus

	for i := 0; i < len(campusArray); i++ {
		// 首先查看是否需要修改
		resCampus, err := campusModel.QueryByCid(campusArray[i].Cid)
		if err != nil {
			return nil, err
		}

		var resultUpdateCampus Models.Campus
		campusModel.Cid = campusArray[i].Cid
		campusModel.Cname = campusArray[i].Cname
		if resCampus.Cname != campusArray[i].Cname {
			resultUpdateCampus, err = campusModel.Update(campusArray[i].Cid)
			if err != nil {
				return nil, err
			}
		}
		result = append(result, resultUpdateCampus)
	}

	return
}
