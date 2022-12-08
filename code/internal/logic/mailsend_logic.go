package logic

import (
	"code/define"
	"code/internal/svc"
	"code/internal/types"
	"code/models"
	"code/utils"
	"context"
	"errors"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
)

type MailsendLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMailsendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MailsendLogic {
	return &MailsendLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MailsendLogic) Mailsend(req *types.MailSendRequest) (resp *types.MailSendResponse, err error) {
	resp = new(types.MailSendResponse)
	user := new(models.UserBasic)
	cnt, err := models.Engine.Where("email = ?", req.Email).Count(user)
	if err != nil {
		return
	}
	if cnt > 0 {
		err = errors.New("邮箱已被注册!!!")
		return
	}

	code := utils.MailCodeCreate()
	// 存入 Redis
	models.InitRedis().Set(l.ctx, req.Email, code, define.CodeExpire)
	// 发送验证码
	err = utils.MailCodeSend(req.Email, code)
	if err != nil {
		return nil, err
	}
	resp = new(types.MailSendResponse)
	resp.Msg = fmt.Sprint("验证码发送成功！！！")

	return
}
