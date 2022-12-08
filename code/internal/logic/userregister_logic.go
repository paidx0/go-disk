package logic

import (
	"code/internal/svc"
	"code/internal/types"
	"code/models"
	"code/utils"
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserregisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserregisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserregisterLogic {
	return &UserregisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserregisterLogic) Userregister(req *types.UserRegisterRequest) (resp *types.UserRegisterResponse, err error) {
	code, err := models.Rdb.Get(l.ctx, req.Email).Result()
	user := new(models.UserBasic)
	if err != nil {
		return nil, errors.New("邮箱验证码不存在！！！")
	}
	if code != req.Code {
		err = errors.New("验证码错误！！！")
		return
	}

	cnt, err := models.Engine.Where("name = ?", req.Name).Count(user)
	if err != nil {
		return nil, err
	}
	if cnt > 0 {
		err = errors.New("用户名已存在！！！")
		return
	}
	user = &models.UserBasic{
		Identity: utils.UUID(), // 随机编号
		Name:     req.Name,
		Password: utils.Md5(req.Password),
		Email:    req.Email,
	}
	_, err = models.Engine.Insert(user)
	if err != nil {
		return nil, errors.New("注册失败，请重试！！！")
	}
	resp = new(types.UserRegisterResponse)
	resp.Msg = fmt.Sprintf("%s,注册成功！！！", req.Name)
	return
}
