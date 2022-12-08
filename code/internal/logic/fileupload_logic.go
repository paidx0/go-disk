package logic

import (
	"code/internal/svc"
	"code/internal/types"
	"code/models"
	"code/utils"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type FileuploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFileuploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileuploadLogic {
	return &FileuploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FileuploadLogic) Fileupload(req *types.FileUploadRequest, UserIdentity string) (resp *types.FileUploadResponse, err error) {
	// 插入到存储仓库池
	rp := &models.RepositoryPool{
		Identity: utils.UUID(),
		Name:     req.Name,
		Hash:     req.Hash,
		Ext:      req.Ext,
		Size:     req.Size,
		Path:     req.Path,
	}
	_, err = models.Engine.Insert(rp)
	if err != nil {
		return nil, err
	}

	// 绑定到用户仓库
	var ParentId int64
	models.Engine.Table("user_basic").
		Select("id").Where("identity = ?", UserIdentity).
		Get(&ParentId)

	up := &models.UserRepository{
		Identity:           utils.UUID(),
		UserIdentity:       UserIdentity,
		ParentId:           ParentId,
		RepositoryIdentity: rp.Identity,
		Ext:                req.Ext,
		Name:               req.Name,
	}
	_, err = models.Engine.Insert(up)
	if err != nil {
		return nil, err
	}

	resp = new(types.FileUploadResponse)
	resp.Identity = rp.Identity
	resp.Ext = rp.Ext
	resp.Name = rp.Name
	return
}
