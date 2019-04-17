package service

import (
	"context"
	"log"
	"mime/multipart"

	"fmt"
	"os"

	"github.com/1024casts/1024casts/util"
	"github.com/gin-gonic/gin"
	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
	"github.com/spf13/viper"
)

type QiNiuService struct {
}

func NewQiNiuService() *QiNiuService {
	return &QiNiuService{}
}

type UploadResponse struct {
	Key  string `json:"key"`
	Hash string `json:"hash"`
}

func (srv *QiNiuService) UploadImage(c *gin.Context, file *multipart.FileHeader) (resp UploadResponse, err error) {

	accessKey := viper.GetString("qiniu.AccessKey")
	secretKey := viper.GetString("qiniu.SecretKey")
	bucket := viper.GetString("qiniu.Bucket")

	saveRootPath := viper.GetString("upload.dst")
	imagePrefix := "uploads/avatar/" + util.GetDate() + "/"
	imagePath := saveRootPath + imagePrefix
	if err = os.MkdirAll(imagePath, 0777); err != nil {
		log.Fatal("[qiniu] create dir err", err)
		return
	}

	key := imagePrefix + file.Filename

	// Upload the file to specific dst.
	dst := saveRootPath + key
	if err = c.SaveUploadedFile(file, dst); err != nil {
		log.Fatal("[qiniu] upload file err", err)
		return
	}

	localFile := dst
	putPolicy := storage.PutPolicy{
		Scope: bucket + ":" + key,
	}

	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuabei
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": key,
		},
	}
	//putExtra.NoCrc32Check = true
	if err = formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra); err != nil {
		fmt.Println(err)
		return
	}

	resp.Key = ret.Key
	resp.Hash = ret.Hash

	return resp, nil

}
