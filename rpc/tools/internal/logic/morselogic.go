package logic

import (
	"STFrontground-backend/rpc/tools/internal/util"
	"context"

	"STFrontground-backend/rpc/tools/internal/svc"
	"STFrontground-backend/rpc/tools/tools"

	"github.com/tal-tech/go-zero/core/logx"
)

type MorseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMorseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MorseLogic {
	return &MorseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MorseLogic) Morse(req *tools.MorseReq) (*tools.MorseResp, error) {
	morseStr, err := util.GenerateMorse(req.Str)
	if err != nil{
		return nil, err
	}
	return &tools.MorseResp{
		MorseStr:morseStr,
	}, nil
}
