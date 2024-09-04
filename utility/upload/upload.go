package upload

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"os"
)

func UploadImgToCloud(ctx context.Context, file *ghttp.UploadFile) (url string, err error) {
	dirPath := "/tmp/"
	name, err := file.Save(dirPath, true)
	if err != nil {
		return "", err
	}
	localFile := dirPath + name
	bucket := g.Cfg().MustGet(ctx, "qiniu.bucket").String()
	accessKey := g.Cfg().MustGet(ctx, "qiniu.accessKey").String()
	secretKey := g.Cfg().MustGet(ctx, "qiniu.secretKey").String()
	//对接七牛云sdk
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey) //鉴权,生成token
	upToken := putPolicy.UploadToken(mac)    //上传token
	cfg := storage.Config{}
	cfg.Zone = &storage.ZoneHuanan
	cfg.UseCdnDomains = false //根据需求官网配置

	formUpload := storage.NewFormUploader(&cfg)
	//上传结果结构体
	ret := storage.PutRet{}
	//可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	key := name
	//七牛云表单上传
	err = formUpload.PutFile(ctx, &ret, upToken, key, localFile, &putExtra)
	if err != nil {
		return "", err
	}
	//删除本地临时文件
	err = os.RemoveAll(localFile)
	if err != nil {
		return "", err
	}
	url = g.Cfg().MustGet(ctx, "qiniu.url").String() + ret.Key
	return
}
