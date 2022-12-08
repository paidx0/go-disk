package logic

import (
	"code/models"
	"context"
	"errors"

	"code/internal/svc"
	"code/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserfilenameupdateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserfilenameupdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserfilenameupdateLogic {
	return &UserfilenameupdateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserfilenameupdateLogic) Userfilenameupdate(req *types.UserFileNameUpdateRequest, UserIdentity string) (resp *types.UserFileNameUpdateResponse, err error) {
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
	// 检查新名字是否已被用户使用
	cnt, err := models.Engine.Where("name = ? AND user_identity = ?", req.NewName, UserIdentity).Count(userRepository)
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		return nil, errors.New("文件名已被使用，请重试！！！")
	}

	userRepository.Name = req.NewName
	_, err = models.Engine.Where("identity = ? AND user_identity = ? ", req.Identity, UserIdentity).Update(userRepository)
	if err != nil {
		return nil, err
	}
	resp = &types.UserFileNameUpdateResponse{Msg: "文件名修改成功！！！"}
	return
}
