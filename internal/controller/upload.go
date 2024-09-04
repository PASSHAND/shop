package controller

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"goframe-shop/api/backend"
	"goframe-shop/internal/consts"
	"goframe-shop/utility/upload"
	"golang.org/x/net/context"
)

type cUpload struct {
}

var Upload = cUpload{}

func (c *cUpload) UploadIngToCloud(ctx context.Context, req *backend.UploadImgToCloudReq) (res *backend.UploadImgToCloudRes, err error) {
	if req.File == nil {
		return nil, gerror.NewCode(gcode.CodeMissingParameter, consts.CodeMissingParameterMsg)
	}
	url, err := upload.UploadImgToCloud(ctx, req.File)
	if err != nil {
		return nil, err
	}
	return &backend.UploadImgToCloudRes{
		Url: url,
	}, nil
}
