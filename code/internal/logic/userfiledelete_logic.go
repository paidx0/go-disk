package logic

import (
	"code/internal/svc"
	"code/internal/types"
	"code/models"
	"context"
	"errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserfiledeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserfiledeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserfiledeleteLogic {
	return &UserfiledeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserfiledeleteLogic) Userfiledelete(req *types.UserFileDeleteRequest, UserIdentity string) (resp *types.UserFileDeleteResponse, err error) {
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

	_, err = models.Engine.Where("user_identity = ? AND identity = ?", UserIdentity, req.Identity).Delete(userRepository)
	if err != nil {
		return nil, errors.New("文件删除失败，请重试！！！")
	}
	resp = new(types.UserFileDeleteResponse)
	resp.Msg = "文件删除成功！！！"
	return
}
