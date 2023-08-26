package service

import (
	"blog-service/global"
	"blog-service/pkg/upload"
	"errors"
	"mime/multipart"
	"os"
)

type FileInfo struct {
	Name      string
	AccessUrl string
}

func (svc *Service) UploadFile(fileType upload.FileType, file multipart.File, fileHeader *multipart.FileHeader) (*FileInfo, error) {
	//取fileName
	fileName := upload.GetFileName(fileHeader.Filename)
	//取得上傳檔案存放位置
	uploadSavePath := upload.GetSavePath()
	//最終檔案路徑 = uploadSavePath/filename
	dst := uploadSavePath + "/" + fileName
	//檢查是否是允許的檔案類型
	if !upload.CheckContainExt(fileType, fileName) {
		return nil, errors.New("file suffix is not supported")
	}
	//檢查存放路徑是否存在，如果不存在則創建路徑
	if upload.CheckSavePath(uploadSavePath) {
		err := upload.CreateSavePath(uploadSavePath, os.ModePerm)
		if err != nil {
			return nil, errors.New("failed to create save directory.")
		}
	}
	//檢查是否過檔案大小限制
	if upload.CheckMaxSize(fileType, file) {
		return nil, errors.New("exceeded maximum file limit.")
	}
	//檢查權限是否足夠
	if upload.CheckPermission(uploadSavePath) {
		return nil, errors.New("insufficient file permissions.")
	}
	//使用SaveFile函數Copy檔案
	if err := upload.SaveFile(fileHeader, dst); err != nil {
		return nil, err
	}

	accessUrl := global.AppSetting.UploadServerUrl + "/" + fileName
	return &FileInfo{Name: fileName, AccessUrl: accessUrl}, nil
}
