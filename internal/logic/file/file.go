package file

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"goframe-shop/internal/consts"
	"goframe-shop/internal/dao"
	"goframe-shop/internal/model"
	"goframe-shop/internal/model/entity"
	"goframe-shop/internal/service"
	"time"
)

type sFile struct {
}

func init() {
	service.RegisterFile(New())
}

func New() *sFile {
	return &sFile{}
}

// 要做：1.上传位置 2.校验 3.定义时间 4.入库 5.返回数据
func (s *sFile) Upload(ctx context.Context, in model.FileUploadInput) (out *model.FileUploadOutput, err error) {
	uploadPath := g.Cfg().MustGet(ctx, "upload.path").String()
	if uploadPath == "" {
		return nil, gerror.New("读取配置文件失败 上传路径不存在")
	}
	if in.Name != "" {
		in.File.Filename = in.Name
	}
	//安全校验：每人一分钟只能上传10次
	count, err := dao.FileInfo.Ctx(ctx).Where(dao.FileInfo.Columns().UserId, gconv.Int(ctx.Value(consts.CtxAdminId))).WhereGTE(dao.FileInfo.Columns().CreatedAt, gtime.Now().Add(-time.Minute)).Count()
	if err != nil {
		return nil, err
	}
	//避免在代码中写死常量
	if count >= consts.FileMaxUploadCountMinute {
		return nil, gerror.New("上传频繁，一分钟只能上传10次")
	}
	//定义时间
	dateDirName := gtime.Now().Format("Ymd")

	//拼接路径并返回文件名
	fileName, err := in.File.Save(gfile.Join(uploadPath, dateDirName), in.RandomName)
	if err != nil {
		return nil, err
	}
	//入库
	data := entity.FileInfo{
		Name: fileName,
		Src:  gfile.Join(uploadPath, dateDirName, fileName),
		Url:  "/upload/" + dateDirName + "/" + fileName, //与上面的join效果一样
	}
	id, err := dao.FileInfo.Ctx(ctx).Data(data).OmitEmpty().InsertAndGetId()
	if err != nil {
		return nil, err
	}
	return &model.FileUploadOutput{
		Id:   uint(id),
		Name: data.Name,
		Src:  data.Src,
		Url:  data.Url}, nil
}
