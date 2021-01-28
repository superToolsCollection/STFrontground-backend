package logic

import (
	"STFrontground-backend/rpc/tools/tools"
	"context"

	"STFrontground-backend/api/internal/svc"
	"STFrontground-backend/api/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type QrcodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewQrcodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) QrcodeLogic {
	return QrcodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QrcodeLogic) Qrcode(req types.QrCodeReq) (*types.QrCodeResp, error) {
	resp, err := l.svcCtx.Tools.QrCode(l.ctx, &tools.QrCodeReq{
		Str: req.Str,
	})
	if err != nil {
		return nil, err
	}
	return &types.QrCodeResp{
		QrCodeStr: resp.QrCodeStr,
	}, nil
}
