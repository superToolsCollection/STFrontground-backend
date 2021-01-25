package logic

import (
	"STFrontground-backend/rpc/tools/tools"
	"context"

	"STFrontground-backend/api/internal/svc"
	"STFrontground-backend/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type Rgb2hexLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRgb2hexLogic(ctx context.Context, svcCtx *svc.ServiceContext) Rgb2hexLogic {
	return Rgb2hexLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Rgb2hexLogic) Rgb2hex(req types.Rgb2HexReq) (*types.Rgb2HexResp, error) {
	resp, err := l.svcCtx.Tools.Rgb2Hex(l.ctx, &tools.Rgb2HexReq{
		Str: req.Str,
	})
	if err != nil {
		return nil, err
	}
	return &types.Rgb2HexResp{
		Rgb2HexStr: resp.Rgb2HexStr,
	}, nil
}
