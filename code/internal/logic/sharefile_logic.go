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

type SharefileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSharefileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SharefileLogic {
	return &SharefileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SharefileLogic) Sharefile(req *types.ShareFileRequest, UserIdentity string) (resp *types.ShareFileResponse, err error) {
	userRepository := new(models.UserRepository)
	// 检查是否有权限修改
	has, err := models.Engine.Where("identity = ?", req.Identity).Select("user_identity,parent_id,repository_identity").Get(userRepository)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("查找的文件不存在！！！")
	}
	if userRepository.UserIdentity != UserIdentity {
		return nil, errors.New("无权修改他人文件！！！")
	}

	share := &models.ShareBasic{
		Identity:           utils.UUID(),
		ParentId:           userRepository.ParentId,
		UserIdentity:       userRepository.UserIdentity,
		RepositoryIdentity: userRepository.RepositoryIdentity,
		ExpiredTime:        req.ExpiredTime,
	}
	_, err = models.Engine.Insert(share)
	if err != nil {
		return nil, errors.New("分享失败，请重试！！！")
	}
	resp = new(types.ShareFileResponse)
	resp.Identity = share.Identity
	resp.Msg = "文件分享成功！！！"
	return
}
