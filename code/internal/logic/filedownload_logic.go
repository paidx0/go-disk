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

type FiledownloadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFiledownloadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FiledownloadLogic {
	return &FiledownloadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FiledownloadLogic) Filedownload(req *types.FileDownLoadRequest, UserIdentity string) (resp *types.FileDownLoadResponse, err error) {
	userRepository := new(models.UserRepository)
	// 检查是否有权限修改
	has, err := models.Engine.Where("identity = ?", req.Identity).Select("user_identity").Get(userRepository)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("查找的文件不存在！！！")
	}
	if userRepository.UserIdentity != UserIdentity {
		return nil, errors.New("无权修改他人文件！！！")
	}

	// 下载拉取
	detail := struct {
		Name string `json:"name"`
		Ext  string `json:"ext"`
		Path string `json:"path"`
	}{}

	models.Engine.Table("user_repository").
		Select("repository_pool.name,repository_pool.ext,repository_pool.path").
		Join("LEFT", "repository_pool", "repository_pool.identity = user_repository.repository_identity").
		Where("user_repository.identity = ?", req.Identity).
		Get(&detail)

	err = utils.FileDownload(detail.Path, detail.Name, detail.Ext)
	if err != nil {
		return nil, err
	}

	resp = new(types.FileDownLoadResponse)
	resp.Msg = "开始下载！！！"
	return
}
