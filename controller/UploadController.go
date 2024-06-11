package controller

import (
	"crypto/sha1"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"hmdp-go/common/codes"
	"hmdp-go/common/constants"
	"hmdp-go/common/logger"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type Upload struct {
	Log logger.ILogger
}

func (a *Upload) UploadImage(c *gin.Context) {
	// 接收文件
	file, err := c.FormFile("file")
	if err != nil {
		RespFail(c, http.StatusOK, codes.ERROR, err.Error())
		return
	}

	// 生成新文件名
	fileName, createErr := createNewFileName(file.Filename)
	if createErr != nil {
		RespFail(c, http.StatusOK, codes.ERROR, createErr.Error())
		return
	}

	// 保存路径
	filePath := path.Join(constants.ImageUploadDir, fileName)
	if err != nil {
		RespFail(c, http.StatusOK, codes.ERROR, err.Error())
		return
	}

	err = c.SaveUploadedFile(file, filePath)
	if err != nil {
		RespFail(c, http.StatusOK, codes.ERROR, err.Error())
		return
	}

	// 返回数据
	RespData(c, http.StatusOK, codes.SUCCESS, file.Filename)
}

func (a *Upload) DeleteBlogImg(c *gin.Context) {
	filename := c.Query("name")
	filePath := constants.ImageUploadDir + filename
	info, err := os.Stat(filePath)
	if err != nil {
		RespFail(c, http.StatusOK, codes.ERROR, err.Error())
		return
	}

	if info.IsDir() {
		RespFail(c, http.StatusOK, codes.ERROR, "错误的文件名称")
		return
	}

	err = os.Remove(filePath)
	if err != nil {
		RespFail(c, http.StatusOK, codes.ERROR, "错误的文件名称")
		return
	}

	RespSuccess(c)
}

// createNewFileName 生成新的文件名，并返回完整路径
func createNewFileName(originalFilename string) (string, error) {
	// 获取后缀
	suffix := filepath.Ext(originalFilename)
	if suffix == "" {
		return "", fmt.Errorf("no file extension found in original filename: %s", originalFilename)
	}

	// 生成UUID
	name := uuid.New().String()

	// 计算哈希值并获取目录结构
	hasher := sha1.New()
	io.WriteString(hasher, name)
	hash := hasher.Sum(nil)
	d1 := hash[0] & 0xF
	d2 := (hash[1] >> 4) & 0xF

	// 构造目录路径
	uploadDir := constants.ImageUploadDir
	dirPath := filepath.Join(uploadDir, fmt.Sprintf("/blogs/%d/%d", d1, d2))

	// 创建目录
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		return "", err
	}

	// 生成文件名
	newFileName := filepath.Join(dirPath, fmt.Sprintf("%s%s", strings.TrimPrefix(name, "-"), suffix))
	return newFileName, nil
}
