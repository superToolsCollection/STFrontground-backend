package logic

import (
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
	// todo: add your logic here and delete this line

	return &types.QrCodeResp{}, nil
}
