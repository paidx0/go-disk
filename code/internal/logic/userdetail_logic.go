package logic

import (
	"code/models"
	"context"
	"github.com/pkg/errors"

	"code/internal/svc"
	"code/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserdetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserdetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserdetailLogic {
	return &UserdetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserdetailLogic) Userdetail(req *types.DetailRequest) (resp *types.DetailResponse, err error) {
	resp = new(types.DetailResponse)
	user := new(models.UserBasic)
	has, err := models.Engine.Where("identity= ?", req.Identity).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("用户不存在！！！")
	}
	resp.Name = user.Name
	resp.Email = user.Email
	return
}
