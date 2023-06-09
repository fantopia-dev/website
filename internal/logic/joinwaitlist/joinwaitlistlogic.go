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
	"github.com/zeromicro/go-zero/core/stores/redis"
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

func (l *JoinWaitListLogic) JoinWaitList(req *types.JoinWaitListReq) (*types.JoinWaitListResp, error) {

	var resp types.JoinWaitListResp
	// verify email
	logx.Infof("email is %v", req.Email)
	_, err := mail.ParseAddress(req.Email)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.INVALID_EMAIL_ERROR), "invalid email")
	}

	// set a lock into redis, to fix concurrent issue
	lock := redis.NewRedisLock(l.svcCtx.Redis, req.Email)
	lock.SetExpire(5)
	ok, err := lock.AcquireCtx(l.ctx)
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "get redis lock error")
	}
	if !ok {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.TOO_MANY_REQUEST_ERROR), "get redis lock failed")
	}
	logx.Info("====get redis lock ok")
	// defer func() {
	// 	lock.ReleaseCtx(l.ctx)
	// 	logx.Info("======release redis lock")
	// }()

	// check exits
	if one, err := l.svcCtx.TbWaitlistModel.FindOneByEmail(l.ctx, req.Email); err == nil {
		resp.Duplicated = true
		resp.Id = int(one.Id)
		return &resp, nil
	}

	// if not exits
	sqlRet, err := l.svcCtx.TbWaitlistModel.Insert(l.ctx, &model.TbWaitlist{
		Email: req.Email,
	})
	if err != nil {
		sqlRet.LastInsertId()
		logx.Errorf("insert database error: %v", err.Error())
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.SERVER_COMMON_ERROR), "database error")
	}

	id, _ := sqlRet.LastInsertId()
	resp.Duplicated = false
	resp.Id = int(id)
	return &resp, nil
}
