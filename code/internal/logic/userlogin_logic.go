package logic

import (
	"code/define"
	"code/models"
	"code/utils"
	"context"
	"github.com/pkg/errors"

	"code/internal/svc"
	"code/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserloginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserloginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserloginLogic {
	return &UserloginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserloginLogic) Userlogin(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	resp = new(types.LoginResponse)
	user := new(models.UserBasic)
	// 登录用户
	has, err := models.Engine.Where("name = ? AND password = ?", req.Name, utils.Md5(req.Password)).Get(user)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, errors.New("账户或密码错误，请重试！！！")
	}
	// 生成 token
	token, err := utils.CreateToken(user.Id, user.Identity, user.Name, define.TokenExpire)
	if err != nil {
		return nil, err
	}
	resp.Token = token
	return
}
