package logic

import (
	"STFrontground-backend/rpc/tools/tools"
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
	resp, err := l.svcCtx.Tools.Morse(l.ctx, &tools.MorseReq{
		Str: req.Str,
	})
	if err != nil {
		return nil, err
	}

	return &types.MorseResp{
		MorseStr: resp.MorseStr,
	}, nil
}
