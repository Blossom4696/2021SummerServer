package Services

import (
	"fmt"
	"image"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/bigby/project/Models"
	"github.com/signintech/gopdf"
)

type Pdf struct {
	curWidth  float64
	curHeight float64
}

// 打印错题
func (mypdf *Pdf) WrongProblemToPDF(wps []int64) (result Res, err error) {
	mypdf.curHeight = 0

	var wpModel Models.WrongProblem

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	pdf.Br(10)
	mypdf.curHeight += 20
	err = pdf.AddTTFFont("SIMKAI", "ttf/SIMKAI.TTF")
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("SIMKAI", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}

	for i, item := range wps {
		var wp Models.WrongProblem
		wp, err = wpModel.QueryByWid(item)
		if err != nil {
			return
		}

		err = mypdf.putContent(&pdf, 20, strconv.Itoa(i+1)+"."+wp.Wtitle)
		if err != nil {
			return
		}

		pdf.SetX(50)
		err = mypdf.putContent(&pdf, 20, "题目:"+wp.Wdescription)
		if err != nil {
			return
		}

		pdf.SetX(70)
		err = mypdf.putContent(&pdf, 20, wp.WproblemTxt)
		if err != nil {
			return
		}

		if !(strings.Contains(wp.WproblemPath, ".mp3")) {
			problemImg := strings.Split(wp.WproblemPath, ";")
			for _, pImg := range problemImg {
				filePath := "/root/2021Summer/images/" + pImg
				var code int

				if pImg == "" {
					continue
				}
				code, err = mypdf.putImage(&pdf, filePath)
				if err != nil {
					return
				}

				if code == -1 {
					continue
				}

			}
		} else {
			pdf.SetX(70)
			err = mypdf.putContent(&pdf, 20, "这是一道听力题")
			if err != nil {
				return
			}
		}

		pdf.SetX(50)
		err = mypdf.putContent(&pdf, 20, "正确答案:"+wp.WanswerTxt)
		if err != nil {
			return
		}

		answerImg := strings.Split(wp.WanswerPath, ";")
		for _, aImg := range answerImg {
			filePath := "/root/2021Summer/images/" + aImg
			var code int

			if aImg == "" {
				continue
			}
			code, err = mypdf.putImage(&pdf, filePath)
			if err != nil {
				return
			}

			if code == -1 {
				continue
			}

		}

		pdf.SetX(50)
		err = mypdf.putContent(&pdf, 20, "我的作答:"+wp.WmyAnswerTxt)
	}

	pdf.WritePdf("/root/2021Summer/pdfs/wrong_problem" + ".pdf")

	if err != nil {
		return Res{
			Code: -1,
			Msg:  "Error!",
			Data: err.Error(),
		}, err
	}

	resData := make(map[string]string)
	resData["filePath"] = "wrong_problem.pdf"

	return Res{
		Code: 1,
		Msg:  "Print Success!",
		Data: resData,
	}, err
}

// 打印习题
func (mypdf *Pdf) ExerciseToPDF(Eids []int64, isPrintAnswer bool) (result Res, err error) {
	mypdf.curHeight = 0

	var exerciseModel Models.Exercise

	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	pdf.Br(10)
	mypdf.curHeight += 20
	err = pdf.AddTTFFont("SIMKAI", "ttf/SIMKAI.TTF")
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("SIMKAI", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}

	for i, item := range Eids {
		var exercise Models.Exercise
		exercise, err = exerciseModel.QueryByEid(item)
		if err != nil {
			return
		}

		err = mypdf.putContent(&pdf, 20, strconv.Itoa(i+1)+"."+exercise.Etitle)
		if err != nil {
			return
		}

		pdf.SetX(50)
		err = mypdf.putContent(&pdf, 20, "题目:"+exercise.Edescription)
		if err != nil {
			return
		}

		pdf.SetX(70)
		err = mypdf.putContent(&pdf, 20, exercise.EnameTxt)

		if !(strings.Contains(exercise.EnamePath, ".mp3")) {
			problemImg := strings.Split(exercise.EnamePath, ";")
			for _, pImg := range problemImg {
				filePath := "/root/2021Summer/images/" + pImg
				var code int

				if pImg == "" {
					continue
				}
				code, err = mypdf.putImage(&pdf, filePath)

				if code == -1 {
					continue
				}

			}
		} else {
			pdf.SetX(70)
			err = mypdf.putContent(&pdf, 20, "这是一道听力题")
		}

		if isPrintAnswer {
			pdf.SetX(50)
			err = mypdf.putContent(&pdf, 20, "正确答案:"+exercise.EanswerTxt)

			answerImg := strings.Split(exercise.EanswerPath, ";")
			for _, aImg := range answerImg {
				filePath := "/root/2021Summer/images/" + aImg
				var code int

				if aImg == "" {
					continue
				}
				code, err = mypdf.putImage(&pdf, filePath)

				if code == -1 {
					continue
				}

			}
		}

	}

	pdf.WritePdf("/root/2021Summer/pdfs/exercise" + ".pdf")

	if err != nil {
		return Res{
			Code: -1,
			Msg:  "Error!",
			Data: err.Error(),
		}, err
	}

	resData := make(map[string]string)
	resData["filePath"] = "exercise.pdf"

	return Res{
		Code: 1,
		Msg:  "Print Success!",
		Data: resData,
	}, err
}

// 放置文本，移动光标纵轴坐标，同时判断是否分页
func (mypdf *Pdf) putContent(pdf *gopdf.GoPdf, num float64, content string) (err error) {
	if mypdf.curHeight+num > gopdf.PageSizeA4.H {
		pdf.AddPage()
		mypdf.curHeight = 0
		err = pdf.Cell(nil, content)
		pdf.Br(10)
	} else {
		mypdf.curHeight += num
		err = pdf.Cell(nil, content)
		pdf.Br(num)
	}
	return
}

// 放置图片，移动光标纵轴坐标，同时判断是否分页
func (mypdf *Pdf) putImage(pdf *gopdf.GoPdf, filePath string) (code int, err error) {
	code = 1
	if isExist, _ := mypdf.isExistFile(filePath); !isExist {
		log.Println("file does exist.")
		code = -1
		return
	}

	width, height, err := mypdf.getImgConfig(filePath)
	if err != nil {
		log.Print(err.Error())
		return
	}
	realHeight := (gopdf.PageSizeA4.W - 2*50) / float64(width) * float64(height)

	file, err := os.Open(filePath)
	if err != nil {
		log.Print(err.Error())
		return
	}

	imgH2, err := gopdf.ImageHolderByReader(file)
	if err != nil {
		log.Print(err.Error())
		return
	}

	if mypdf.curHeight+realHeight > gopdf.PageSizeA4.H {
		pdf.AddPage()
		pdf.Br(10)
		mypdf.curHeight = 10
		err = pdf.ImageByHolder(imgH2, 50, mypdf.curHeight, nil)
		mypdf.curHeight += realHeight
		pdf.Br(realHeight)
	} else {
		err = pdf.ImageByHolder(imgH2, 50, mypdf.curHeight, nil)
		mypdf.curHeight += realHeight
		pdf.Br(realHeight)
	}
	return
}

func (mypdf *Pdf) getImgConfig(fileStr string) (width int, height int, err error) {
	file, err := os.Open(fileStr)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}

	c, _, err := image.DecodeConfig(file)
	if err != nil {
		fmt.Println("err = ", err)
		return
	}
	width = c.Width
	height = c.Height
	return
}

func (mypdf *Pdf) isExistFile(fileStr string) (isExist bool, err error) {
	_, err = os.Stat(fileStr)
	isExist = true
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatalln("file does not exist")
			isExist = false
			return
		}
	}
	return
}
