package user

import (
	"boat/rpc/user/pb/user"
	"context"

	"boat/gateway/internal/svc"
	"boat/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.IdReq) (resp *types.UserResp, err error) {
	user, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdRequest{
		Id: "0",
	})
	if err != nil {
		return nil, err
	}

	return &types.UserResp{
		Id:   1,
		Name: user.Name,
	}, nil
}
