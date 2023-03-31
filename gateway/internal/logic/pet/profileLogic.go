package pet

import (
	"context"

	"boat/gateway/internal/svc"
	"boat/gateway/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProfileLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProfileLogic {
	return &ProfileLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ProfileLogic) Profile(req *types.PetReq) (resp *types.PetResp, err error) {
	// todo: add your logic here and delete this line

	return
}
