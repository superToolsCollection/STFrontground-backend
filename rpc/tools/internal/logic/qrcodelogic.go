package logic

import (
	"STFrontground-backend/rpc/tools/internal/util"
	"context"

	"STFrontground-backend/rpc/tools/internal/svc"
	"STFrontground-backend/rpc/tools/tools"

	"github.com/tal-tech/go-zero/core/logx"
)

type QrCodeLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQrCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QrCodeLogic {
	return &QrCodeLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *QrCodeLogic) QrCode(req *tools.QrCodeReq) (*tools.QrCodeResp, error) {
	qrCodeBytes, err := util.GenerateQRCodeByte(req.Str)
	if err != nil {
		return nil, err
	}
	encode, err := util.EncodeBase64(string(qrCodeBytes))
	if err != nil {
		return nil, err
	}
	return &tools.QrCodeResp{
		QrCodeStr: encode,
	}, nil
}
