package joinwaitlist

import (
	"context"
	"net/mail"

	"github.com/fantopia-dev/website/internal/svc"
	"github.com/fantopia-dev/website/internal/types"
	"github.com/fantopia-dev/website/model"
	"github.com/fantopia-dev/website/xerr"
	"github.com/pkg/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinWaitListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewJoinWaitListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinWaitListLogic {
	return &JoinWaitListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JoinWaitListLogic) JoinWaitList(req *types.JoinWaitListReq) (string, error) {

	// verify email
	logx.Infof("email is %v", req.Email)
	_, err := mail.ParseAddress(req.Email)
	if err != nil {
		return "", errors.Wrapf(xerr.NewErrCode(xerr.INVALID_EMAIL_ERROR), "invalid email")
	}

	// check exits
	if _, err = l.svcCtx.TbWaitlistModel.FindOneByEmail(l.ctx, req.Email); err == nil {
		return "success", nil
	}

	// if not exits
	if _, err = l.svcCtx.TbWaitlistModel.Insert(l.ctx, &model.TbWaitlist{
		Email: req.Email,
	}); err != nil {
		logx.Errorf("insert database error: %v", err.Error())
		return "", errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "database error")
	}

	return "success", nil
}
