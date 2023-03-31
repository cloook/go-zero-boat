package logic

import (
	"context"

	"boat/common/xerr"
	"boat/rpc/user/internal/svc"
	"boat/rpc/user/pb/user"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type TestErrLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewTestErrLogic(ctx context.Context, svcCtx *svc.ServiceContext) *TestErrLogic {
	return &TestErrLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *TestErrLogic) TestErr(in *user.TestRequest) (*user.UserResponse, error) {

	return nil, errors.Wrapf(xerr.NewErrCodeMsg(xerr.SERVER_COMMON_ERROR, "testErr Msg"), "some Msg log... ")
}
