package Controllers

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// @Summary 上传文件
// @Description 上传客户端发送的文件至本地服务器
// @Tags File
// @Accept json
// @Produce  json
// @Param dirName formData string true "文件本地路径"
// @Router /app/file/upload_file [post]
func FileUpload(c *gin.Context) {
	file, header, err := c.Request.FormFile("uploadFile")
	if err != nil {
		c.JSON(http.StatusOK, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	dirName := c.PostForm("dirName")
	if err != nil {
		c.JSON(http.StatusOK, Res{
			Code: -1,
			Msg:  "Error: " + err.Error(),
			Data: nil,
		})
		return
	}

	filename := header.Filename

	out, err := os.Create("/root/2021Summer/" + dirName + "/" + filename)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, Res{
		Code: 1,
		Msg:  "Upload Success",
		Data: filename,
	})
}

// @Summary 获取图片
// @Description 下载服务端文件夹中的图片
// @Tags File
// @Accept json
// @Produce  json
// @Param name query string true "文件路径"
// @Router /app/file/get_image [get]
func ImageDownload(c *gin.Context) {
	path := "/root/2021Summer/images/"
	fileName := path + c.Query("name")
	c.File(fileName)
}

// @Summary 获取音频
// @Description 下载服务端文件夹中的音频
// @Tags File
// @Accept json
// @Produce  json
// @Param name query string true "文件路径"
// @Router /app/file/get_audio [get]
func AudioDownload(c *gin.Context) {
	path := "/root/2021Summer/audios/"
	fileName := path + c.Query("name")
	c.File(fileName)
}

// @Summary 获取pdf
// @Description 下载服务端文件夹中的pdf
// @Tags File
// @Accept json
// @Produce  json
// @Param name query string true "文件路径"
// @Router /app/file/get_pdf [get]
func PdfDownload(c *gin.Context) {
	path := "/root/2021Summer/pdfs/"
	fileName := path + c.Query("name")
	c.File(fileName)
}
