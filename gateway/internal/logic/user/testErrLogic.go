package user

import (
	"boat/rpc/user/pb/user"
	"context"

	"boat/gateway/internal/svc"
	"boat/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type TestErrLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewTestErrLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestErrLogic {
	return &TestErrLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *TestErrLogic) TestErr() (resp *types.UserResp, err error) {
	_, err = l.svcCtx.UserRpc.TestErr(l.ctx, &user.TestRequest{})
	if err != nil {
		return nil, err
	}
	return
}
