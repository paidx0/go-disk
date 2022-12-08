package logic

import (
	"code/models"
	"code/utils"
	"context"
	"errors"

	"code/internal/svc"
	"code/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SharefiledownloadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSharefiledownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SharefiledownloadLogic {
	return &SharefiledownloadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SharefiledownloadLogic) Sharefiledownload(req *types.ShareFileDownloadRequest) (resp *types.ShareFileDownloadResponse, err error) {
	// 点击次数+1
	_, err = models.Engine.Exec("UPDATE share_basic SET click_num = click_num + 1 WHERE identity = ?", req.Identity)
	if err != nil {
		return nil, errors.New("执行失败，请重试！！！")
	}
	// 下载拉取
	detail := struct {
		Name string `json:"name"`
		Ext  string `json:"ext"`
		Path string `json:"path"`
	}{}

	has, err := models.Engine.Table("share_basic").
		Select("repository_pool.name,repository_pool.ext,repository_pool.path").
		Join("LEFT", "repository_pool", "repository_pool.identity = share_basic.repository_identity").
		Where("share_basic.identity = ?", req.Identity).
		Get(&detail)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("未查到分享文件，请重试！！！")
	}

	err = utils.FileDownload(detail.Path, detail.Name, detail.Ext)
	if err != nil {
		return nil, err
	}

	resp = new(types.ShareFileDownloadResponse)
	resp.Msg = "开始下载！！！"
	return
}
