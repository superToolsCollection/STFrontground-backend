package logic

import (
	"context"

	"STFrontground-backend/api/internal/svc"
	"STFrontground-backend/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type MorseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMorseLogic(ctx context.Context, svcCtx *svc.ServiceContext) MorseLogic {
	return MorseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MorseLogic) Morse(req types.MorseReq) (*types.MorseResp, error) {
	// todo: add your logic here and delete this line

	return &types.MorseResp{}, nil
}
