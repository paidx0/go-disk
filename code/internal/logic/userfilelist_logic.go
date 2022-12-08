package logic

import (
	"code/define"
	"code/models"
	"context"
	"time"

	"code/internal/svc"
	"code/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserfilelistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserfilelistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserfilelistLogic {
	return &UserfilelistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserfilelistLogic) Userfilelist(req *types.UserFileListRequest, UserIdentity string) (resp *types.UserFileListResponse, err error) {
	userfilelist := make([]*types.UserFile, 0)
	pageNum := req.PageNum
	PageSize := req.PageSize
	if PageSize == 0 {
		PageSize = define.PageSize
	}
	if pageNum == 0 {
		pageNum = 1
	}
	userRepository := new(models.UserRepository)
	userRepository.UserIdentity = UserIdentity
	// 用户文件列表
	err = models.Engine.Table("user_repository").Where("user_identity = ? ", userRepository.UserIdentity).
		Select("user_repository.id, user_repository.identity, user_repository.repository_identity, user_repository.ext,user_repository.name, repository_pool.path, repository_pool.size").
		Join("LEFT", "repository_pool", "user_repository.repository_identity = repository_pool.identity").
		Where("user_repository.deleted_at = ? OR user_repository.deleted_at IS NULL", time.Time{}.Format(define.Datetime)).
		Limit(PageSize, (pageNum-1)*PageSize).Find(&userfilelist)
	if err != nil {
		return nil, err
	}
	// 用户文件总数
	cnt, err := models.Engine.Where("user_identity = ? ", userRepository.UserIdentity).Count(userRepository)
	if err != nil {
		return nil, err
	}
	resp = new(types.UserFileListResponse)
	resp.List = userfilelist
	resp.Count = cnt
	return
}
