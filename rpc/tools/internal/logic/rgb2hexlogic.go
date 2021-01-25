package logic

import (
	"STFrontground-backend/rpc/tools/internal/util"
	"context"

	"STFrontground-backend/rpc/tools/internal/svc"
	"STFrontground-backend/rpc/tools/tools"

	"github.com/tal-tech/go-zero/core/logx"
)

type Rgb2HexLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRgb2HexLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Rgb2HexLogic {
	return &Rgb2HexLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *Rgb2HexLogic) Rgb2Hex(req *tools.Rgb2HexReq) (*tools.Rgb2HexResp, error) {
	str, err := util.RgbToHex(req.Str)
	if err != nil {
		return nil, err
	}
	return &tools.Rgb2HexResp{
		Rgb2HexStr: str,
	}, nil
}
