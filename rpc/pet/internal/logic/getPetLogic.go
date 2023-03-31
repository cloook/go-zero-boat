package logic

import (
	"context"

	"boat/rpc/pet/internal/svc"
	"boat/rpc/pet/pb/pet"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPetLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPetLogic {
	return &GetPetLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetPetLogic) GetPet(in *pet.IdRequest) (*pet.PetResponse, error) {
	// todo: add your logic here and delete this line

	return &pet.PetResponse{}, nil
}
