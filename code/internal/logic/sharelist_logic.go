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

type SharelistLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSharelistLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SharelistLogic {
	return &SharelistLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SharelistLogic) Sharelist(req *types.ShareFileListRequest) (resp *types.ShareFileListResponse, err error) {
	pageNum := req.PageNum
	PageSize := req.PageSize
	if PageSize == 0 {
		PageSize = define.PageSize
	}
	if pageNum == 0 {
		pageNum = 1
	}
	// 分享列表
	shareList := make([]*types.ShareFile, 0)
	err = models.Engine.Table("share_basic").
		Select("share_basic.identity, share_basic.expired_time,share_basic.click_num,repository_pool.ext,repository_pool.name, repository_pool.path, repository_pool.size").
		Join("LEFT", "repository_pool", "share_basic.repository_identity = repository_pool.identity").
		Where("share_basic.deleted_at = ? OR share_basic.deleted_at IS NULL", time.Time{}.Format(define.Datetime)).
		Limit(PageSize, (pageNum-1)*PageSize).Find(&shareList)
	if err != nil {
		return nil, err
	}
	// 分享文件数
	cnt, err := models.Engine.Count(new(models.ShareBasic))
	if err != nil {
		return nil, err
	}
	resp = new(types.ShareFileListResponse)
	resp.List = shareList
	resp.Count = cnt
	return
}
