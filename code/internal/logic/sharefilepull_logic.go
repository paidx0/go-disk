package logic

import (
	"code/internal/svc"
	"code/internal/types"
	"code/models"
	"code/utils"
	"context"
	"errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type SharefilepullLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSharefilepullLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SharefilepullLogic {
	return &SharefilepullLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SharefilepullLogic) Sharefilepull(req *types.ShareFilePullRequest, UserIdentity string) (resp *types.ShareFilePullResponse, err error) {
	userRepository := new(models.UserRepository)
	shareDetail := struct {
		Identity           string `json:"identity"`
		UserIdentity       string `json:"user_identity"`
		ParentId           int64  `json:"parentId"`
		RepositoryIdentity string `json:"repository_identity"`
		Ext                string `json:"ext"`
		Name               string `json:"name"`
	}{}

	has, err := models.Engine.Table("share_basic").
		Select("repository_pool.ext,repository_pool.name,share_basic.repository_identity").
		Join("LEFT", "repository_pool", "repository_pool.identity = share_basic.repository_identity").
		Where("share_basic.identity = ?", req.Identity).
		Get(&shareDetail)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("未查到分享文件，请重试！！！")
	}

	ID := new(models.UserBasic)
	models.Engine.ShowSQL(true)
	_, err = models.Engine.Select("id").Where("identity = ?", UserIdentity).Get(ID)
	shareDetail.ParentId = int64(ID.Id)
	// 插入用户仓库
	userRepository.Identity = utils.UUID()
	userRepository.UserIdentity = UserIdentity
	userRepository.ParentId = shareDetail.ParentId
	userRepository.RepositoryIdentity = shareDetail.RepositoryIdentity
	userRepository.Ext = shareDetail.Ext
	userRepository.Name = shareDetail.Name
	_, err = models.Engine.Insert(userRepository)
	if err != nil {
		return nil, errors.New("插入失败，请重试！！！")
	}

	resp = new(types.ShareFilePullResponse)
	resp.Msg = "拉取成功！！！"
	return
}
