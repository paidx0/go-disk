package logic

import (
	"code/models"
	"context"
	"errors"

	"code/internal/svc"
	"code/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserfileowenupdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserfileowenupdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserfileowenupdateLogic {
	return &UserfileowenupdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserfileowenupdateLogic) Userfileowenupdate(req *types.UserFileOwenUpdateRequest, UserIdentity string) (resp *types.UserFileOwenUpdateResponse, err error) {
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

	var newOwen int64
	has, err = models.Engine.Table("user_basic").
		Select("id").Where("identity = ?", req.UserIdentity).
		Get(&newOwen)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("转让用户不存在，请重试！！！")
	}

	// 更新用户仓库对应文件的拥有者的ID
	_, err = models.Engine.Where("identity = ? AND user_identity = ?", req.Identity, UserIdentity).
		Update(models.UserRepository{ParentId: newOwen, UserIdentity: req.UserIdentity})
	if err != nil {
		return nil, errors.New("转让失败，请重试！！！")
	}
	resp = new(types.UserFileOwenUpdateResponse)
	resp.Msg = "文件转让成功！！！"
	return
}
